// options.go
package sql

import (
	"fmt"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var (
	defaultTimezone        = "UTC"
	defaultMaxIdleConns    = 10
	defaultMaxOpenConns    = 100
	defaultConnMaxLifetime = 30 * time.Minute
	defaultConnMaxIdleTime = 30 * time.Minute
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
		maxIdleConns:    defaultMaxIdleConns,
		maxOpenConns:    defaultMaxOpenConns,
		connMaxLifetime: defaultConnMaxLifetime,
		connMaxIdleTime: defaultConnMaxIdleTime,
		logLevel:        defaultLogLevel,
	}
}

// DBOptionFunc defines a function type for setting options.
type DBOptionFunc func(*DBOption)

// WithSQLDialect sets the SQL dialect option.
func WithSQLDialect(dialect SQLDialect) DBOptionFunc {
	return func(o *DBOption) {
		o.sqlDialect = dialect
	}
}

// WithDatabaseName sets the database name option.
func WithDatabaseName(name string) DBOptionFunc {
	return func(o *DBOption) {
		o.databaseName = name
	}
}

// WithHost sets the database host option.
func WithHost(host string) DBOptionFunc {
	return func(o *DBOption) {
		o.host = host
	}
}

// WithPort sets the database port option.
func WithPort(port int) DBOptionFunc {
	return func(o *DBOption) {
		o.port = port
	}
}

// WithUser sets the database user option.
func WithUser(user string) DBOptionFunc {
	return func(o *DBOption) {
		o.user = user
	}
}

// WithPassword sets the database password option.
func WithPassword(password string) DBOptionFunc {
	return func(o *DBOption) {
		o.password = password
	}
}

// WithTimezone sets the timezone option.
func WithTimezone(timezone string) DBOptionFunc {
	return func(o *DBOption) {
		o.timezone = timezone
	}
}

// WithMaxIdleConns sets the maximum number of idle connections.
func WithMaxIdleConns(maxIdleConns int) DBOptionFunc {
	return func(o *DBOption) {
		o.maxIdleConns = maxIdleConns
	}
}

// WithMaxOpenConns sets the maximum number of open connections.
func WithMaxOpenConns(maxOpenConns int) DBOptionFunc {
	return func(o *DBOption) {
		o.maxOpenConns = maxOpenConns
	}
}

// WithConnMaxLifetime sets the maximum connection lifetime.
func WithConnMaxLifetime(connMaxLifetime time.Duration) DBOptionFunc {
	return func(o *DBOption) {
		o.connMaxLifetime = connMaxLifetime
	}
}

// WithConnMaxIdleTime sets the maximum idle time for a connection.
func WithConnMaxIdleTime(connMaxIdleTime time.Duration) DBOptionFunc {
	return func(o *DBOption) {
		o.connMaxIdleTime = connMaxIdleTime
	}
}

// WithLogLevel sets the log level for GORM.
func WithLogLevel(logLevel gormLogger.LogLevel) DBOptionFunc {
	return func(o *DBOption) {
		o.logLevel = logLevel
	}
}

// validateMandatoryOptions checks if all mandatory options are set and valid.
func validateMandatoryOptions(option *DBOption) error {
	var validationErrors []string

	if option.sqlDialect.Name == "" {
		validationErrors = append(validationErrors, "invalid sqlDialect")
	}
	if option.databaseName == "" {
		validationErrors = append(validationErrors, "invalid database name")
	}
	if option.sqlDialect.Name != SQLite.Name {
		if option.host == "" {
			validationErrors = append(validationErrors, "invalid host")
		}
		if option.port == 0 {
			validationErrors = append(validationErrors, "invalid port")
		}
		if option.user == "" {
			validationErrors = append(validationErrors, "invalid user")
		}
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("validation errors: %v", validationErrors)
	}
	return nil
}
