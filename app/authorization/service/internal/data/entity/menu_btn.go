package entity

import (
	"time"
)

type MenuBtn struct {
	Id          int64
	MenuId      int64
	Name        string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (MenuBtn) TableName() string {
	return "authorization_menu_btns"
}
