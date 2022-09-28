package entity

import (
	"time"
)

type Role struct {
	Id        int64
	Name      string
	ParentId  int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Children  []Role `gorm:"-"`
}

func (Role) TableName() string {
	return "authorization_roles"
}
