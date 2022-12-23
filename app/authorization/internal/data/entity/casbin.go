package entity

type CasbinRule struct {
	ID     int64  `json:"id"`
	PType  string `json:"ptype" gorm:"column:p_type" `
	Role   string `json:"role_name" gorm:"column:v0" ` // 代表角色 管理员昵称不用
	Path   string `json:"path" gorm:"column:v1" `
	Method string `json:"method" gorm:"column:v2" `
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
