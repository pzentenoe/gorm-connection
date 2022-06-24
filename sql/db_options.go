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

type DBOption struct {
	sqlDialect   SQLDialect
	databaseName *string
	host         *string
	port         *int
	user         *string
	password     *string
}

func Config() *DBOption {
	return new(DBOption)
}

func (c *DBOption) SetSQLDialect(sqlDialect SQLDialect) *DBOption {
	c.sqlDialect = sqlDialect
	return c
}

func (c *DBOption) DatabaseName(databaseName string) *DBOption {
	c.databaseName = &databaseName
	return c
}

func (c *DBOption) Host(host string) *DBOption {
	c.host = &host
	return c
}

func (c *DBOption) Port(port int) *DBOption {
	c.port = &port
	return c
}

func (c *DBOption) User(user string) *DBOption {
	c.user = &user
	return c
}

func (c *DBOption) Password(password string) *DBOption {
	c.password = &password
	return c
}

func MergeOptions(opts ...*DBOption) *DBOption {
	option := new(DBOption)

	for _, opt := range opts {
		if opt.sqlDialect != "" {
			option.sqlDialect = opt.sqlDialect
		} else {
			panic("Invalid sqlDialect")
		}
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		} else {
			panic("Invalid database name")
		}
		if opt.host != nil {
			option.host = opt.host
		} else {
			panic("invalid host")
		}
		if opt.port != nil {
			option.port = opt.port
		} else {
			panic("invalid port")
		}
		if opt.user != nil {
			option.user = opt.user
		} else {
			panic("invalid user")
		}
		if opt.password != nil {
			option.password = opt.password
		} else {
			panic("empty password")
		}
	}
	return option
}

type SQLDialect string

// Url formats
const (
	// "sqlserver://user:pass@host:port?database=dbName"
	urlSQLServerFormat = "sqlserver://%v:%v@%v:%v?database=%v"
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	urlMysqlFormat = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	urlPostgresFormat = "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=America/Santiago"
	// "/tmp/gorm.db"
	urlSQLiteFormat = "%v.db"
)

// SQL Dialects
const (
	SQLServer SQLDialect = "mssql"
	MySQL     SQLDialect = "mysql"
	Postgres  SQLDialect = "postgres"
	SQLite    SQLDialect = "sqlite3"
)

func (c *DBOption) getGormDialector() gorm.Dialector {
	switch c.sqlDialect {
	case SQLServer:
		url := fmt.Sprintf(urlSQLServerFormat, *c.user, urlesc.QueryEscape(*c.password), *c.host, *c.port, *c.databaseName)
		return sqlserver.Open(url)
	case MySQL:
		url := fmt.Sprintf(urlMysqlFormat, *c.user, *c.password, *c.host, *c.port, *c.databaseName)
		return mysql.Open(url)
	case Postgres:
		url := fmt.Sprintf(urlPostgresFormat, *c.host, *c.user, *c.password, *c.databaseName, *c.port)
		return postgres.Open(url)
	case SQLite:
		return sqlite.Open(fmt.Sprintf(urlSQLiteFormat, *c.databaseName))
	}
	return nil
}
