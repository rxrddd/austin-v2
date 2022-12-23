package entity

type RoleMenuBtn struct {
	Id     int64
	RoleId int64
	MenuId int64
	BtnId  int64
}

func (RoleMenuBtn) TableName() string {
	return "authorization_role_menu_btn"
}
