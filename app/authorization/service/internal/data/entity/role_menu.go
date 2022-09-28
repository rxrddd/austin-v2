package entity

type RoleMenu struct {
	Id     int64
	RoleId int64
	MenuId int64
}

func (RoleMenu) TableName() string {
	return "authorization_role_menu"
}
