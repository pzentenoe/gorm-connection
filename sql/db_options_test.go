package sql

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestMergeOptions(t *testing.T) {
	t.Run("when merge options works ok", func(t *testing.T) {
		opts1 := Config().
			SetSQLDialect(Postgres).
			DatabaseName("testdb").
			Host("localhost").
			Port(5432).
			User("testuser").
			Password("testpassword")

		opts2 := Config().
			SetSQLDialect(Postgres)

		mergedOptions, err := MergeOptions(opts1, opts2)
		assert.NoError(t, err)

		assert.Equal(t, Postgres, mergedOptions.sqlDialect)
		assert.Equal(t, "testdb", *mergedOptions.databaseName)
		assert.Equal(t, "localhost", *mergedOptions.host)
		assert.Equal(t, 5432, *mergedOptions.port)
		assert.Equal(t, "testuser", *mergedOptions.user)
		assert.Equal(t, "testpassword", *mergedOptions.password)
	})
}

func TestDBOption_getGormDialector(t *testing.T) {
	t.Run("", func(t *testing.T) {
		opts := Config().
			SetSQLDialect(SQLite).
			DatabaseName("test")

		dialector := opts.getGormDialector()
		assert.NotNil(t, dialector)

		_, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		assert.NoError(t, err)
	})
}
