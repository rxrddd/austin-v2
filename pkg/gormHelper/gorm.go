package gormHelper

import (
	"gorm.io/gorm"
)

func RecordUnDeleted() func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at = 0")
	}
}

func RecordDeleted() func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at != 0")
	}
}

