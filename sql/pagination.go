package sql

import "gorm.io/gorm"

const (
	maxPageSize = 100
	maxSize     = 10
)

// Paginate returns a GORM scope function to apply pagination
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 || pageSize > maxPageSize {
			pageSize = maxSize
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
