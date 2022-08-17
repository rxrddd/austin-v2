package entity

import (
	"gorm.io/gorm"
	"time"
)

const (
	// AdministratorDeleted 已经删除
	AdministratorDeleted = "1"
	// AdministratorUnDeleted 未删除
	AdministratorUnDeleted = "2"
	// AdministratorStatusOK 状态正常
	AdministratorStatusOK = 1
	// AdministratorStatusForbid 状态禁用
	AdministratorStatusForbid = 2
)

type AdministratorEntity struct {
	Id        int64
	Username  string
	Password  string
	Salt      string
	Mobile    string
	Nickname  string
	Avatar    string
	Status    int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt string
}

func (AdministratorEntity) TableName() string {
	return "sys_administrator"
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
