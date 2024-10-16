package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGormDialector_SQLite(t *testing.T) {
	options := &DBOption{
		sqlDialect:   SQLite,
		databaseName: testDBName,
	}

	dialector := options.getGormDialector()
	assert.NotNil(t, dialector)
}

func TestGetGormDialector_Postgres(t *testing.T) {
	options := &DBOption{
		sqlDialect:   Postgres,
		databaseName: testDBName,
		host:         "localhost",
		port:         5432,
		user:         "user",
		password:     "password",
		timezone:     "UTC",
	}

	dialector := options.getGormDialector()
	assert.NotNil(t, dialector)
}

func TestGetGormDialector_InvalidDialect(t *testing.T) {
	options := &DBOption{
		sqlDialect:   SQLDialect{Name: "unknown"},
		databaseName: testDBName,
	}

	dialector := options.getGormDialector()
	assert.Nil(t, dialector)
}
