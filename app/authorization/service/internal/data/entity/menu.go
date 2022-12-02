package entity

import (
	"time"
)

type Menu struct {
	Id        int64
	ParentId  int64
	ParentIds string
	Name      string
	Path      string
	Hidden    int64
	Component string
	Sort      int64
	Title     string
	Icon      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	MenuBtns  []*MenuBtn `gorm:"foreignKey:menu_id;"`
	Children  []Menu     `gorm:"-"`
}

func (Menu) TableName() string {
	return "authorization_menus"
}
