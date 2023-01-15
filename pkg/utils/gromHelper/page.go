package gromHelper

import "gorm.io/gorm"

func Page(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		return db.
			Limit(int(pageSize)).
			Offset(int((page - 1) * pageSize))
	}
}
