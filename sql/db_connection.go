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

// Connection is an interface that defines the method to get a database connection.
type Connection interface {
	GetConnection() (*gorm.DB, error)
}

// dbConnection holds the details for connecting to the database.
type dbConnection struct {
	dialector  gorm.Dialector
	options    *DBOption
	connection *gorm.DB
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

	return &dbConnection{
		options:   options,
		dialector: dialector,
	}, nil
}

// GetConnection establishes and returns a GORM DB connection.
// If the connection is already established, it returns the existing one.
func (r *dbConnection) GetConnection() (*gorm.DB, error) {
	if r.connection == nil {
		newLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  r.options.logLevel,
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
			},
		)
		connection, err := gorm.Open(r.dialector, &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			return nil, fmt.Errorf("error trying to connect to DB: %w", err)
		}

		sqlDB, errConnect := connection.DB()
		if errConnect != nil {
			return nil, fmt.Errorf("error getting DB instance: %w", errConnect)
		}

		sqlDB.SetMaxIdleConns(r.options.maxIdleConns)
		sqlDB.SetMaxOpenConns(r.options.maxOpenConns)
		sqlDB.SetConnMaxLifetime(r.options.connMaxLifetime)
		sqlDB.SetConnMaxIdleTime(r.options.connMaxIdleTime)

		r.connection = connection
	}
	return r.connection, nil
}
