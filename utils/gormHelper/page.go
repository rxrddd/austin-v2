package gormHelper

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Int interface {
	int | int8 | int32 | int64 | uint | uint8 | uint32 | uint64
}

func Page[T Int](page T, pageSize T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		if pageSize < 1 {
			pageSize = 20
		}
		return db.
			Limit(int(pageSize)).
			Offset(int((page - 1) * pageSize))
	}
}

func QueryPage[T Int](page T, pageSize T) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if page < 1 {
			page = 1
		}
		if pageSize < 1 {
			pageSize = 20
		}

		if pageSize > 100 {
			pageSize = 100
		}

		return dao.
			Limit(int(pageSize)).
			Offset(int((page - 1) * pageSize))
	}
}
