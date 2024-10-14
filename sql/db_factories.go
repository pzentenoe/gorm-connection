package sql

// Factory function for SQLite
func NewSQLiteConnection(databaseName string, options ...DBOptionFunc) (Connection, error) {
	opts := append([]DBOptionFunc{
		WithSQLDialect(SQLite),
		WithDatabaseName(databaseName),
	}, options...)
	return NewSQLConnection(opts...)
}

// Factory function for SQL Server
func NewSQLServerConnection(host, databaseName, user, password string, options ...DBOptionFunc) (Connection, error) {
	opts := append([]DBOptionFunc{
		WithSQLDialect(SQLServer),
		WithHost(host),
		WithDatabaseName(databaseName),
		WithUser(user),
		WithPassword(password),
		WithPort(SQLServer.DefaultPort),
	}, options...)
	return NewSQLConnection(opts...)
}

// Factory function for MySQL
func NewMySQLConnection(host, databaseName, user, password string, options ...DBOptionFunc) (Connection, error) {
	opts := append([]DBOptionFunc{
		WithSQLDialect(MySQL),
		WithHost(host),
		WithDatabaseName(databaseName),
		WithUser(user),
		WithPassword(password),
		WithPort(MySQL.DefaultPort),
	}, options...)
	return NewSQLConnection(opts...)
}

// Factory function for PostgreSQL
func NewPostgresConnection(host, databaseName, user, password string, options ...DBOptionFunc) (Connection, error) {
	opts := append([]DBOptionFunc{
		WithSQLDialect(Postgres),
		WithHost(host),
		WithDatabaseName(databaseName),
		WithUser(user),
		WithPassword(password),
		WithPort(Postgres.DefaultPort),
	}, options...)
	return NewSQLConnection(opts...)
}
