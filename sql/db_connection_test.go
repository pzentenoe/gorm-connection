package sql

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewSQLConnection(t *testing.T) {
	t.Run("Should be able to create new connection", func(t *testing.T) {
		options := Config().
			SetSQLDialect(SQLServer).
			Host("host").
			Port(3323).
			User("user").
			Timezone("America/Los_Angeles").
			MaxIdleConns(10).
			MaxOpenConns(50).
			ConnMaxLifetime(time.Minute * 20).
			ConnMaxIdleTime(time.Minute * 10).
			Password("password").
			DatabaseName("test")
		//SetLogLevel(gormLogger.Info)

		conn, err := NewSQLConnection(options)
		assert.NoError(t, err)
		assert.NotNil(t, conn)
	})

	t.Run("Should be able to create new connection with default port", func(t *testing.T) {
		options := Config().
			SetSQLDialect(SQLServer).
			Host("host").
			User("user").
			Password("password").
			DatabaseName("test")

		conn, err := NewSQLConnection(options)
		assert.NoError(t, err)
		assert.NotNil(t, conn)
	})
}

func TestDBConnection_GetConnection(t *testing.T) {
	t.Run("Should be able to get connection", func(t *testing.T) {
		options := Config().
			SetSQLDialect(SQLite).
			DatabaseName("test")

		conn, err := NewSQLConnection(options)
		assert.NoError(t, err)
		assert.NotNil(t, conn)

		db, err := conn.GetConnection()
		assert.NoError(t, err)
		assert.NotNil(t, db)

		sqlDB, err := db.DB()
		assert.NoError(t, err)

		assert.Equal(t, sqlDB.Stats().Idle, 1)
		assert.Equal(t, sqlDB.Stats().MaxOpenConnections, defaultMaxOpenCons)
	})

}
