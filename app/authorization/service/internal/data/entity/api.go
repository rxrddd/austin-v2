package entity

import (
	"time"
)

type Api struct {
	Id        int64
	Group     string
	Name      string
	Method    string
	Path      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (Api) TableName() string {
	return "authorization_api"
}
