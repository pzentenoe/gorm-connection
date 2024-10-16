// connection.go
package sql

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Connection defines the method to get a database connection.
type Connection interface {
	GetConnection() (*gorm.DB, error)
}

// dbConnection holds the details for connecting to the database.
type dbConnection struct {
	dialector   gorm.Dialector
	options     *DBOption
	connection  *gorm.DB
	logger      gormLogger.Interface
	connectFunc func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error)
}

// NewSQLConnection creates a new SQL database connection based on the provided options.
func NewSQLConnection(optionFuncs ...DBOptionFunc) (Connection, error) {
	options := defaultConfig()
	for _, optFunc := range optionFuncs {
		optFunc(options)
	}

	if err := validateMandatoryOptions(options); err != nil {
		return nil, err
	}

	if options.port == 0 && options.sqlDialect.DefaultPort != 0 {
		options.port = options.sqlDialect.DefaultPort
	}

	dialector := options.getGormDialector()
	if dialector == nil {
		return nil, errors.New("error creating connection, empty dialector")
	}

	logger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  options.logLevel,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
		},
	)

	return &dbConnection{
		options:     options,
		dialector:   dialector,
		logger:      logger,
		connectFunc: gorm.Open,
	}, nil
}

// GetConnection establishes and returns a GORM DB connection.
// If the connection is already established, it returns the existing one.
func (r *dbConnection) GetConnection() (*gorm.DB, error) {
	if r.connection == nil {
		connection, err := retryConnection(3, 2*time.Second, func() (*gorm.DB, error) {
			return gorm.Open(r.dialector, &gorm.Config{
				Logger: r.logger,
			})
		})
		if err != nil {
			return nil, err
		}

		sqlDB, err := connection.DB()
		if err != nil {
			return nil, fmt.Errorf("error getting DB instance: %w", err)
		}

		sqlDB.SetMaxIdleConns(r.options.maxIdleConns)
		sqlDB.SetMaxOpenConns(r.options.maxOpenConns)
		sqlDB.SetConnMaxLifetime(r.options.connMaxLifetime)
		sqlDB.SetConnMaxIdleTime(r.options.connMaxIdleTime)

		r.connection = connection
	}
	return r.connection, nil
}

func retryConnection(attempts int, delay time.Duration, connectFunc func() (*gorm.DB, error)) (*gorm.DB, error) {
	var err error
	for i := 0; i < attempts; i++ {
		var db *gorm.DB
		if db, err = connectFunc(); err == nil {
			return db, nil
		}
		log.Printf("Attempt %d/%d failed: %v. Retrying in %s...", i+1, attempts, err, delay)
		time.Sleep(delay)
	}
	return nil, fmt.Errorf("could not establish a connection after %d attempts: %w", attempts, err)
}
