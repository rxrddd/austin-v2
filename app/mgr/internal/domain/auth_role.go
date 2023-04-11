package domain

type RoleListReq struct {
	PageNo   int64
	PageSize int64
	Keywords string
}

type AdminListReq struct {
	PageNo   int64
	PageSize int64
	Username string
	Nickname string
	Role     string
}

type UpdateInfoReq struct {
	Avatar          string
	Username        string
	Nickname        string
	Password        string //新密码
	PasswordConfirm string //确认密码
}
