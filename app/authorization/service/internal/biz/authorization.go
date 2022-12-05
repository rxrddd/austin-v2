package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Authorization struct {
	UserIdentity string
	Sub          string
	Obj          string
	Act          string
}

type AuthorizationRepo interface {
	CheckAuthorization(context.Context, *Authorization) (isSuccess bool, err error)
	GetRoleList(ctx context.Context) ([]*Role, error)
	CreateRole(ctx context.Context, role *Role) (*Role, error)
	UpdateRole(ctx context.Context, role *Role) (*Role, error)
	DeleteRole(ctx context.Context, id int64) error
	AddRolesForUser(ctx context.Context, username string, roles []string) (bool, error)
	GetRolesForUser(ctx context.Context, username string) ([]string, error)
	GetUsersForRole(ctx context.Context, user string) ([]string, error)
	DeleteRoleForUser(ctx context.Context, username string, role string) (bool, error)
	DeleteRolesForUser(ctx context.Context, username string) (bool, error)
	GetPolicies(ctx context.Context, role string) ([]*PolicyRules, error)
	UpdatePolicies(ctx context.Context, username string, rules []PolicyRules) (bool, error)
	GetApiList(ctx context.Context, page int64, pageSize int64, group, name, method, path string) ([]*Api, int64, error)
	GetApiAll(ctx context.Context) ([]*Api, error)
	CreateApi(ctx context.Context, api *Api) (*Api, error)
	UpdateApi(ctx context.Context, api *Api) (*Api, error)
	DeleteApi(ctx context.Context, id int64) error
	GetMenuTree(ctx context.Context) ([]*Menu, error)
	GetMenuAll(ctx context.Context) ([]*Menu, error)
	CreateMenu(ctx context.Context, menu *Menu) (*Menu, error)
	UpdateMenu(ctx context.Context, menu *Menu) (*Menu, error)
	DeleteMenu(ctx context.Context, id int64) error
	SaveRoleMenu(ctx context.Context, roleId int64, menuIds []int64) error
	GetRoleMenuBtn(ctx context.Context, roleId int64, menuId int64) ([]int64, error)
	SetRoleMenuBtn(ctx context.Context, roleId int64, menuId int64, btnIds []int64) error
	GetRoleMenu(ctx context.Context, role string) ([]*Menu, error)
	GetRoleMenuTree(ctx context.Context, role string) ([]*Menu, error)
}

type AuthorizationUsecase struct {
	repo AuthorizationRepo
	log  *log.Helper
}

func NewAuthorizationUsecase(repo AuthorizationRepo, logger log.Logger) *AuthorizationUsecase {
	return &AuthorizationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AuthorizationUsecase) CheckAuthorization(ctx context.Context, authorization *Authorization) (bool, error) {
	result, err := uc.repo.CheckAuthorization(ctx, authorization)
	if err != nil {
		return false, err
	}
	return result, nil
}
