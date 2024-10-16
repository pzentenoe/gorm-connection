package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDBName = "test"

func TestValidateMandatoryOptions_AllValid(t *testing.T) {
	options := &DBOption{
		sqlDialect:   Postgres,
		databaseName: testDBName,
		host:         "localhost",
		port:         5432,
		user:         "user",
		password:     "password",
	}

	err := validateMandatoryOptions(options)
	assert.NoError(t, err)
}

func TestValidateMandatoryOptions_MissingDialect(t *testing.T) {
	options := &DBOption{
		databaseName: testDBName,
		host:         "localhost",
		port:         5432,
		user:         "user",
		password:     "password",
	}

	err := validateMandatoryOptions(options)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid sqlDialect")
}

func TestValidateMandatoryOptions_MissingDatabaseName(t *testing.T) {
	options := &DBOption{
		sqlDialect: Postgres,
		host:       "localhost",
		port:       5432,
		user:       "user",
		password:   "password",
	}

	err := validateMandatoryOptions(options)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid database name")
}

func TestValidateMandatoryOptions_SQLite(t *testing.T) {
	options := &DBOption{
		sqlDialect:   SQLite,
		databaseName: testDBName,
	}

	err := validateMandatoryOptions(options)
	assert.NoError(t, err)
}
