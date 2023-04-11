// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLaSystemAuthRole = "la_system_auth_role"

// LaSystemAuthRole mapped from table <la_system_auth_role>
type LaSystemAuthRole struct {
	ID         int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键
	Name       string     `gorm:"column:name;not null" json:"name"`                  // 角色名称
	Remark     string     `gorm:"column:remark;not null" json:"remark"`              // 备注信息
	Sort       int32      `gorm:"column:sort;not null" json:"sort"`                  // 角色排序
	IsDisable  int32      `gorm:"column:is_disable;not null" json:"is_disable"`      // 是否禁用: 0=否, 1=是
	CreateTime int64      `gorm:"column:create_time;not null" json:"create_time"`    // 创建时间
	UpdateTime int64      `gorm:"column:update_time;not null" json:"update_time"`    // 更新时间
	MenuIds    SplitSlice `gorm:"column:menu_ids" json:"menu_ids"`
}

// TableName LaSystemAuthRole's table name
func (*LaSystemAuthRole) TableName() string {
	return TableNameLaSystemAuthRole
}
