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
// It merges the provided options and initializes the dialector. If the dialector is empty, it returns an error.
func NewSQLConnection(opts ...*DBOption) (Connection, error) {
	databaseOptions, err := MergeOptions(opts...)
	if err != nil {
		return nil, err
	}
	dialector := databaseOptions.getGormDialector()
	if dialector == nil {
		return nil, errors.New("error creating connection, empty dialector")
	}
	return &dbConnection{
		options:   databaseOptions,
		dialector: dialector,
	}, nil
}

// GetConnection establishes and returns a GORM DB connection. If the connection is already established, it returns the existing one.
// It sets up the logger, connection pool configurations, and handles errors during the connection process.
func (r *dbConnection) GetConnection() (*gorm.DB, error) {
	if r.connection == nil {
		newLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  *r.options.logLevel,
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

		sqlDB.SetMaxIdleConns(*r.options.maxIdleConns)
		sqlDB.SetMaxOpenConns(*r.options.maxOpenConns)
		sqlDB.SetConnMaxLifetime(*r.options.connMaxLifetime)
		sqlDB.SetConnMaxIdleTime(*r.options.connMaxIdleTime)

		r.connection = connection
	}
	return r.connection, nil
}
