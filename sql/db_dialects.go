package sql

import (
	"fmt"
	"github.com/PuerkitoBio/urlesc"
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
	urlSQLServerFormat = "sqlserver://%v:%v@%v:%v?database=%v"
	urlMysqlFormat     = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	urlPostgresFormat  = "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v"
)

// getGormDialector returns the appropriate GORM Dialector based on the SQL dialect and other options.
func (c *DBOption) getGormDialector() gorm.Dialector {
	switch c.sqlDialect {
	case SQLServer:
		url := fmt.Sprintf(urlSQLServerFormat, c.user, urlesc.QueryEscape(c.password), c.host, c.port, c.databaseName)
		return sqlserver.Open(url)
	case MySQL:
		url := fmt.Sprintf(urlMysqlFormat, c.user, c.password, c.host, c.port, c.databaseName)
		return mysql.Open(url)
	case Postgres:
		url := fmt.Sprintf(urlPostgresFormat, c.host, c.user, c.password, c.databaseName, c.port, c.timezone)
		return postgres.Open(url)
	case SQLite:
		return sqlite.Open(fmt.Sprintf("%s.db", c.databaseName))
	default:
		return nil
	}
}
