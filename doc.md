# sql
--
    import "."


## Usage

```go
var (
	SQLServer = SQLDialect{Name: "mssql", DefaultPort: 1433}
	MySQL     = SQLDialect{Name: "mysql", DefaultPort: 3306}
	Postgres  = SQLDialect{Name: "postgres", DefaultPort: 5432}
	SQLite    = SQLDialect{Name: "sqlite3"}
)
```
SQL Dialects

#### func  Paginate

```go
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB
```
Paginate returns a GORM scope function to apply pagination.

#### type Connection

```go
type Connection interface {
	GetConnection() (*gorm.DB, error)
}
```

Connection is an interface that defines the method to get a database connection.

#### func  NewSQLConnection

```go
func NewSQLConnection(opts ...*DBOption) (Connection, error)
```
NewSQLConnection creates a new SQL database connection based on the provided
options. It merges the provided options and initializes the dialector. If the
dialector is empty, it returns an error.

#### type DBConnection

```go
type DBConnection struct {
}
```

DBConnection holds the details for connecting to the database.

#### func (*DBConnection) GetConnection

```go
func (r *DBConnection) GetConnection() (*gorm.DB, error)
```
GetConnection establishes and returns a GORM DB connection. If the connection is
already established, it returns the existing one. It sets up the logger,
connection pool configurations, and handles errors during the connection
process.

#### type DBOption

```go
type DBOption struct {
}
```

DBOption holds the configuration for database connection

#### func  Config

```go
func Config() *DBOption
```
Config initializes a new DBOption with default values

#### func  MergeOptions

```go
func MergeOptions(opts ...*DBOption) (*DBOption, error)
```
MergeOptions merges multiple DBOption instances into one, applying default
values where necessary. It returns the merged DBOption and an error if any
mandatory option is missing or invalid.

#### func (*DBOption) ConnMaxIdleTime

```go
func (c *DBOption) ConnMaxIdleTime(connMaxIdleTime time.Duration) *DBOption
```
ConnMaxIdleTime sets the maximum amount of time a connection may be idle before
being closed. It returns the modified DBOption for chaining.

#### func (*DBOption) ConnMaxLifetime

```go
func (c *DBOption) ConnMaxLifetime(connMaxLifetime time.Duration) *DBOption
```
ConnMaxLifetime sets the maximum amount of time a connection may be reused. It
returns the modified DBOption for chaining.

#### func (*DBOption) DatabaseName

```go
func (c *DBOption) DatabaseName(databaseName string) *DBOption
```
DatabaseName sets the name of the database to connect to. It returns the
modified DBOption for chaining.

#### func (*DBOption) Host

```go
func (c *DBOption) Host(host string) *DBOption
```
Host sets the host address of the database server. It returns the modified
DBOption for chaining.

#### func (*DBOption) MaxIdleConns

```go
func (c *DBOption) MaxIdleConns(maxIdleConns int) *DBOption
```
MaxIdleConns sets the maximum number of idle connections in the connection pool.
It returns the modified DBOption for chaining.

#### func (*DBOption) MaxOpenConns

```go
func (c *DBOption) MaxOpenConns(maxOpenConns int) *DBOption
```
MaxOpenConns sets the maximum number of open connections to the database. It
returns the modified DBOption for chaining.

#### func (*DBOption) Password

```go
func (c *DBOption) Password(password string) *DBOption
```
Password sets the password for database authentication. It returns the modified
DBOption for chaining.

#### func (*DBOption) Port

```go
func (c *DBOption) Port(port int) *DBOption
```
Port sets the port on which the database server is listening. It returns the
modified DBOption for chaining.

#### func (*DBOption) SetSQLDialect

```go
func (c *DBOption) SetSQLDialect(sqlDialect SQLDialect) *DBOption
```
SetSQLDialect sets the SQL dialect to be used (e.g., SQLServer, MySQL,
Postgres). It returns the modified DBOption for chaining.

#### func (*DBOption) Timezone

```go
func (c *DBOption) Timezone(timezone string) *DBOption
```
Timezone sets the timezone to be used for the database connection. It returns
the modified DBOption for chaining.

#### func (*DBOption) User

```go
func (c *DBOption) User(user string) *DBOption
```
User sets the username for database authentication. It returns the modified
DBOption for chaining.

#### type SQLDialect

```go
type SQLDialect struct {
	Name        string
	DefaultPort int
}
```

SQLDialect represents a SQL dialect with its name and default port
