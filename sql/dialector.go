// dialect.go
package sql

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// SQLDialect represents a SQL dialect with its name and default port.
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

// URL formats for various SQL dialects.
const (
	urlSQLServerFormat = "sqlserver://%s:%s@%s:%d?database=%s"
	urlMysqlFormat     = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	urlPostgresFormat  = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s"
)

// getGormDialector returns the appropriate GORM Dialector based on the SQL dialect and other options.
func (c *DBOption) getGormDialector() gorm.Dialector {
	switch c.sqlDialect {
	case SQLServer:
		escapedPassword := url.QueryEscape(c.password)
		dsn := fmt.Sprintf(urlSQLServerFormat, c.user, escapedPassword, c.host, c.port, c.databaseName)
		return sqlserver.Open(dsn)
	case MySQL:
		dsn := fmt.Sprintf(urlMysqlFormat, c.user, c.password, c.host, c.port, c.databaseName)
		return mysql.Open(dsn)
	case Postgres:
		dsn := fmt.Sprintf(urlPostgresFormat, c.host, c.user, c.password, c.databaseName, c.port, c.timezone)
		return postgres.Open(dsn)
	case SQLite:
		return sqlite.Open(fmt.Sprintf("%s.db", c.databaseName))
	default:
		return nil
	}
}
