package sql

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewSQLConnection(t *testing.T) {
	t.Run("Test connection OK", func(t *testing.T) {
		conn, err := NewSQLConnection(
			WithSQLDialect(SQLite),
			WithDatabaseName(testDBName),
		)
		assert.NoError(t, err)
		assert.NotNil(t, conn)
	})

	t.Run("MissingMandatoryOptions", func(t *testing.T) {
		conn, err := NewSQLConnection()
		assert.Error(t, err)
		assert.Nil(t, conn)
	})
}

func TestGetConnection(t *testing.T) {
	t.Run("Test connection OK", func(t *testing.T) {
		conn, err := NewSQLConnection(
			WithSQLDialect(SQLite),
			WithDatabaseName(testDBName),
		)
		assert.NoError(t, err)
		assert.NotNil(t, conn)

		dbConn, err := conn.GetConnection()
		assert.NoError(t, err)
		assert.NotNil(t, dbConn)

		dbConn2, err := conn.GetConnection()
		assert.NoError(t, err)
		assert.NotNil(t, dbConn2)
		assert.Equal(t, dbConn, dbConn2)
	})
}

func TestRetryConnection(t *testing.T) {
	t.Run("Test Retry Connection Success", func(t *testing.T) {
		attempts := 3
		delay := 10 * time.Millisecond
		callCount := 0
		connectFunc := func() (*gorm.DB, error) {
			callCount++
			return &gorm.DB{}, nil
		}

		db, err := retryConnection(attempts, delay, connectFunc)
		assert.NoError(t, err)
		assert.NotNil(t, db)
		assert.Equal(t, 1, callCount)
	})

	t.Run("Test Retry Connection Failure", func(t *testing.T) {
		attempts := 3
		delay := 10 * time.Millisecond
		callCount := 0
		connectFunc := func() (*gorm.DB, error) {
			callCount++
			return nil, errors.New("connection failed")
		}

		db, err := retryConnection(attempts, delay, connectFunc)
		assert.Error(t, err)
		assert.Nil(t, db)
		assert.Equal(t, attempts, callCount)
	})

	t.Run("Test Retry Connection Succeeds After Retry", func(t *testing.T) {
		attempts := 3
		delay := 10 * time.Millisecond
		callCount := 0
		connectFunc := func() (*gorm.DB, error) {
			callCount++
			if callCount == 2 {
				return &gorm.DB{}, nil
			}
			return nil, errors.New("connection failed")
		}

		db, err := retryConnection(attempts, delay, connectFunc)
		assert.NoError(t, err)
		assert.NotNil(t, db)
		assert.Equal(t, 2, callCount)
	})
}
