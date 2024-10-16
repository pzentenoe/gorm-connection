// factory.go
package sql

// NewSQLiteConnection Factory function for SQLite
func NewSQLiteConnection(databaseName string, options ...DBOptionFunc) (Connection, error) {
	return NewConnection(SQLite,
		append([]DBOptionFunc{
			WithDatabaseName(databaseName),
		}, options...)...)
}

// NewSQLServerConnection Factory function for SQL Server
func NewSQLServerConnection(host, databaseName, user, password string, options ...DBOptionFunc) (Connection, error) {
	return NewConnection(SQLServer,
		append([]DBOptionFunc{
			WithHost(host),
			WithDatabaseName(databaseName),
			WithUser(user),
			WithPassword(password),
		}, options...)...)
}

// NewMySQLConnection Factory function for MySQL
func NewMySQLConnection(host, databaseName, user, password string, options ...DBOptionFunc) (Connection, error) {
	return NewConnection(MySQL,
		append([]DBOptionFunc{
			WithHost(host),
			WithDatabaseName(databaseName),
			WithUser(user),
			WithPassword(password),
		}, options...)...)
}

// NewPostgresConnection Factory function for PostgreSQL
func NewPostgresConnection(host, databaseName, user, password string, options ...DBOptionFunc) (Connection, error) {
	return NewConnection(Postgres,
		append([]DBOptionFunc{
			WithHost(host),
			WithDatabaseName(databaseName),
			WithUser(user),
			WithPassword(password),
		}, options...)...)
}

// NewConnection is a generalized factory for creating a database connection.
func NewConnection(dialect SQLDialect, options ...DBOptionFunc) (Connection, error) {
	opts := append([]DBOptionFunc{
		WithSQLDialect(dialect),
		WithPort(dialect.DefaultPort),
	}, options...)
	return NewSQLConnection(opts...)
}
