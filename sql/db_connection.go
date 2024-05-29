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

type Connection interface {
	GetConnection() (*gorm.DB, error)
}

type DBConnection struct {
	dialector  gorm.Dialector
	options    *DBOption
	connection *gorm.DB
}

func NewSQLConnection(opts ...*DBOption) (Connection, error) {
	databaseOptions, err := MergeOptions(opts...)
	if err != nil {
		return nil, err
	}
	dialector := databaseOptions.getGormDialector()
	if dialector == nil {
		return nil, errors.New("error creating connection, empty dialector")
	}
	return &DBConnection{
		options:   databaseOptions,
		dialector: dialector,
	}, nil
}

func (r *DBConnection) GetConnection() (*gorm.DB, error) {
	if r.connection == nil {
		newLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Second,
				LogLevel:      gormLogger.Warn,
				Colorful:      false,
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
