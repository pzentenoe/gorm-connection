package sql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type TestTable struct {
	ID   uint
	Name string
}

func TestPaginate(t *testing.T) {
	t.Run("when paginate works ok", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		assert.NoError(t, err)

		// Creación de la tabla antes de insertar datos
		err = db.AutoMigrate(&TestTable{})
		assert.NoError(t, err)

		// Inserción de datos en la tabla de prueba
		for i := 0; i < 100; i++ {
			db.Create(&TestTable{Name: fmt.Sprintf("Name %d", i)})
		}

		var results []TestTable

		page := 1
		pageSize := 10

		db.Scopes(Paginate(page, pageSize)).Find(&results)
		assert.Len(t, results, pageSize)
	})
}
