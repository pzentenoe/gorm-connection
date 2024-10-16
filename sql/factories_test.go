package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSQLiteConnection(t *testing.T) {
	conn, err := NewSQLiteConnection(testDBName)
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

func TestNewSQLServerConnection(t *testing.T) {
	conn, err := NewSQLServerConnection("localhost", testDBName, "user", "password")
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

func TestNewMySQLConnection(t *testing.T) {
	conn, err := NewMySQLConnection("localhost", testDBName, "user", "password")
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

func TestNewPostgresConnection(t *testing.T) {
	conn, err := NewPostgresConnection("localhost", testDBName, "user", "password")
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}
