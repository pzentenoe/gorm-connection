package sql

import (
	"fmt"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var (
	defaultTimezone        = "UTC"
	defaultMaxIdleCons     = 10
	defaultMaxOpenCons     = 100
	defaultConnMaxLifetime = time.Minute * 30
	defaultConnMaxIdleTime = time.Minute * 30
	defaultLogLevel        = gormLogger.Info
)

// DBOption holds the configuration for database connection.
type DBOption struct {
	sqlDialect      SQLDialect
	databaseName    string
	host            string
	port            int
	user            string
	password        string
	timezone        string
	maxIdleConns    int
	maxOpenConns    int
	logLevel        gormLogger.LogLevel
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration
}

// defaultConfig initializes a new DBOption with default values.
func defaultConfig() *DBOption {
	return &DBOption{
		timezone:        defaultTimezone,
		maxIdleConns:    defaultMaxIdleCons,
		maxOpenConns:    defaultMaxOpenCons,
		connMaxLifetime: defaultConnMaxLifetime,
		connMaxIdleTime: defaultConnMaxIdleTime,
		logLevel:        defaultLogLevel,
	}
}

// DBOptionFunc defines a function type for setting options.
type DBOptionFunc func(*DBOption)

// Option setters.
func WithSQLDialect(dialect SQLDialect) DBOptionFunc {
	return func(o *DBOption) {
		o.sqlDialect = dialect
	}
}

func WithDatabaseName(name string) DBOptionFunc {
	return func(o *DBOption) {
		o.databaseName = name
	}
}

func WithHost(host string) DBOptionFunc {
	return func(o *DBOption) {
		o.host = host
	}
}

func WithPort(port int) DBOptionFunc {
	return func(o *DBOption) {
		o.port = port
	}
}

func WithUser(user string) DBOptionFunc {
	return func(o *DBOption) {
		o.user = user
	}
}

func WithPassword(password string) DBOptionFunc {
	return func(o *DBOption) {
		o.password = password
	}
}

func WithTimezone(timezone string) DBOptionFunc {
	return func(o *DBOption) {
		o.timezone = timezone
	}
}

func WithMaxIdleConns(maxIdleConns int) DBOptionFunc {
	return func(o *DBOption) {
		o.maxIdleConns = maxIdleConns
	}
}

func WithMaxOpenConns(maxOpenConns int) DBOptionFunc {
	return func(o *DBOption) {
		o.maxOpenConns = maxOpenConns
	}
}

func WithConnMaxLifetime(connMaxLifetime time.Duration) DBOptionFunc {
	return func(o *DBOption) {
		o.connMaxLifetime = connMaxLifetime
	}
}

func WithConnMaxIdleTime(connMaxIdleTime time.Duration) DBOptionFunc {
	return func(o *DBOption) {
		o.connMaxIdleTime = connMaxIdleTime
	}
}

func WithLogLevel(logLevel gormLogger.LogLevel) DBOptionFunc {
	return func(o *DBOption) {
		o.logLevel = logLevel
	}
}

// validateMandatoryOptions checks if all mandatory options are set and valid.
func validateMandatoryOptions(option *DBOption) error {
	if option.sqlDialect.Name == "" {
		return fmt.Errorf("invalid sqlDialect")
	}
	if option.databaseName == "" {
		return fmt.Errorf("invalid database name")
	}

	if option.sqlDialect != SQLite {
		if option.host == "" {
			return fmt.Errorf("invalid host")
		}
		if option.port == 0 {
			return fmt.Errorf("invalid port")
		}
		if option.user == "" {
			return fmt.Errorf("invalid user")
		}
	}
	return nil
}
