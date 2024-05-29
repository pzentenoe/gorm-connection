package sql

import (
	"fmt"
	"github.com/PuerkitoBio/urlesc"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

var (
	defaultTimezone        = "UTC"
	defaultMaxIdleCons     = 10
	defaultMaxOpenCons     = 100
	defaultConnMaxLifetime = time.Minute * 30
	defaultConnMaxIdleTime = time.Minute * 30
)

// DBOption holds the configuration for database connection
type DBOption struct {
	sqlDialect      SQLDialect
	databaseName    *string
	host            *string
	port            *int
	user            *string
	password        *string
	timezone        *string
	maxIdleConns    *int
	maxOpenConns    *int
	connMaxLifetime *time.Duration
	connMaxIdleTime *time.Duration
}

// Config initializes a new DBOption with default values
func Config() *DBOption {
	return new(DBOption)
}

// SetSQLDialect sets the SQL dialect to be used (e.g., SQLServer, MySQL, Postgres).
// It returns the modified DBOption for chaining.
func (c *DBOption) SetSQLDialect(sqlDialect SQLDialect) *DBOption {
	c.sqlDialect = sqlDialect
	return c
}

// DatabaseName sets the name of the database to connect to.
// It returns the modified DBOption for chaining.
func (c *DBOption) DatabaseName(databaseName string) *DBOption {
	c.databaseName = &databaseName
	return c
}

// Host sets the host address of the database server.
// It returns the modified DBOption for chaining.
func (c *DBOption) Host(host string) *DBOption {
	c.host = &host
	return c
}

// Port sets the port on which the database server is listening.
// It returns the modified DBOption for chaining.
func (c *DBOption) Port(port int) *DBOption {
	c.port = &port
	return c
}

// User sets the username for database authentication.
// It returns the modified DBOption for chaining.
func (c *DBOption) User(user string) *DBOption {
	c.user = &user
	return c
}

// Password sets the password for database authentication.
// It returns the modified DBOption for chaining.
func (c *DBOption) Password(password string) *DBOption {
	c.password = &password
	return c
}

// Timezone sets the timezone to be used for the database connection.
// It returns the modified DBOption for chaining.
func (c *DBOption) Timezone(timezone string) *DBOption {
	c.timezone = &timezone
	return c
}

// MaxIdleConns sets the maximum number of idle connections in the connection pool.
// It returns the modified DBOption for chaining.
func (c *DBOption) MaxIdleConns(maxIdleConns int) *DBOption {
	c.maxIdleConns = &maxIdleConns
	return c
}

// MaxOpenConns sets the maximum number of open connections to the database.
// It returns the modified DBOption for chaining.
func (c *DBOption) MaxOpenConns(maxOpenConns int) *DBOption {
	c.maxOpenConns = &maxOpenConns
	return c
}

// ConnMaxLifetime sets the maximum amount of time a connection may be reused.
// It returns the modified DBOption for chaining.
func (c *DBOption) ConnMaxLifetime(connMaxLifetime time.Duration) *DBOption {
	c.connMaxLifetime = &connMaxLifetime
	return c
}

// ConnMaxIdleTime sets the maximum amount of time a connection may be idle before being closed.
// It returns the modified DBOption for chaining.
func (c *DBOption) ConnMaxIdleTime(connMaxIdleTime time.Duration) *DBOption {
	c.connMaxIdleTime = &connMaxIdleTime
	return c
}

// MergeOptions merges multiple DBOption instances into one, applying default values where necessary.
// It returns the merged DBOption and an error if any mandatory option is missing or invalid.
func MergeOptions(opts ...*DBOption) (*DBOption, error) {
	option := new(DBOption)

	for _, opt := range opts {
		if opt.sqlDialect.Name != "" {
			option.sqlDialect = opt.sqlDialect
		}
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		}
		if opt.host != nil {
			option.host = opt.host
		}
		if opt.port != nil {
			option.port = opt.port
		} else {
			port := option.sqlDialect.DefaultPort
			option.port = &port
		}
		if opt.user != nil {
			option.user = opt.user
		}
		if opt.password != nil {
			option.password = opt.password
		}
		if opt.timezone != nil {
			option.timezone = opt.timezone
		} else {
			option.timezone = &defaultTimezone
		}
		if opt.maxIdleConns != nil {
			option.maxIdleConns = opt.maxIdleConns
		} else {
			option.maxIdleConns = &defaultMaxIdleCons
		}
		if opt.maxOpenConns != nil {
			option.maxOpenConns = opt.maxOpenConns
		} else {
			option.maxOpenConns = &defaultMaxOpenCons
		}
		if opt.connMaxLifetime != nil {
			option.connMaxLifetime = opt.connMaxLifetime
		} else {
			option.connMaxLifetime = &defaultConnMaxLifetime
		}
		if opt.connMaxIdleTime != nil {
			option.connMaxIdleTime = opt.connMaxIdleTime
		} else {
			option.connMaxIdleTime = &defaultConnMaxIdleTime
		}
	}

	err := validateMandatoryOptions(option)
	if err != nil {
		return nil, err
	}

	return option, nil
}

// validateMandatoryOptions checks if all mandatory options are set and valid.
// It returns an error if any mandatory option is missing or invalid.
func validateMandatoryOptions(option *DBOption) error {
	if option.sqlDialect.Name == "" {
		return fmt.Errorf("Invalid sqlDialect")
	}
	if option.databaseName == nil {
		return fmt.Errorf("Invalid database name")
	}

	if option.sqlDialect != SQLite {
		if option.host == nil {
			return fmt.Errorf("Invalid host")
		}
		if option.port == nil {
			return fmt.Errorf("Invalid port")
		}
		if option.user == nil {
			return fmt.Errorf("Invalid user")
		}
		if option.password == nil {
			return fmt.Errorf("Empty password")
		}
	}
	return nil
}

// SQLDialect represents a SQL dialect with its name and default port
type SQLDialect struct {
	Name        string
	DefaultPort int
}

// SQL Dialects
var (
	SQLServer = SQLDialect{Name: "mssql", DefaultPort: 1433}
	MySQL     = SQLDialect{Name: "mysql", DefaultPort: 3306}
	Postgres  = SQLDialect{Name: "postgres", DefaultPort: 5432}
	SQLite    = SQLDialect{Name: "sqlite3"}
)

// URL formats for various SQL dialects
const (
	urlSQLServerFormat = "sqlserver://%v:%v@%v:%v?database=%v"
	urlMysqlFormat     = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	urlPostgresFormat  = "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v"
	urlSQLiteFormat    = "%v.db"
)

// getGormDialector returns the appropriate GORM Dialector based on the SQL dialect and other options
func (c *DBOption) getGormDialector() gorm.Dialector {
	switch c.sqlDialect {
	case SQLServer:
		url := fmt.Sprintf(urlSQLServerFormat, *c.user, urlesc.QueryEscape(*c.password), *c.host, *c.port, *c.databaseName)
		return sqlserver.Open(url)
	case MySQL:
		url := fmt.Sprintf(urlMysqlFormat, *c.user, *c.password, *c.host, *c.port, *c.databaseName)
		return mysql.Open(url)
	case Postgres:
		url := fmt.Sprintf(urlPostgresFormat, *c.host, *c.user, *c.password, *c.databaseName, *c.port, *c.timezone)
		return postgres.Open(url)
	case SQLite:
		return sqlite.Open(fmt.Sprintf(urlSQLiteFormat, *c.databaseName))
	default:
		return nil
	}
}
