package entity

import (
	"time"
)

const (
	// AdministratorStatusOK 状态正常
	AdministratorStatusOK = 1
	// AdministratorStatusForbid 状态禁用
	AdministratorStatusForbid = 2
)

type AdministratorEntity struct {
	Id            int64
	Username      string
	Password      string
	Salt          string
	Mobile        string
	Nickname      string
	Avatar        string
	Status        int64
	Role          string
	LastLoginTime string
	LastLoginIp   string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     string
}

func (AdministratorEntity) TableName() string {
	return "sys_administrator"
}
