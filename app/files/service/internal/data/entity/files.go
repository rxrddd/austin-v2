package entity

import (
	"gorm.io/gorm"
	"time"
)

const (
	// FilesDeleted 已经删除
	FilesDeleted = "1"
	// FilesUnDeleted 未删除
	FilesUnDeleted = "2"
)

type FilesEntity struct {
	Id        int64
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt string
}

func (FilesEntity) TableName() string {
	return "files"
}

// Paginate 分页
func Paginate(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

// UnDelete 非删除数据
func UnDelete() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at = ''")
	}
}

func HasDelete() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at != ''")
	}
}
