package sql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

type MockDialector struct {
	sqlite.Dialector
}

func (m MockDialector) Initialize(*gorm.DB) error {
	return fmt.Errorf("mock error initializing")
}

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
	opts := Config().
		MaxIdleConns(1).
		MaxOpenConns(1).
		ConnMaxLifetime(time.Minute * 20).
		ConnMaxIdleTime(time.Minute * 20)

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

	t.Run("connection is nil", func(t *testing.T) {

		r := &DBConnection{
			dialector: MockDialector{sqlite.Dialector{}},
			options:   opts,
		}

		_, err := r.GetConnection()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mock error initializing")
	})

}
