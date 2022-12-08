package biz

import (
	"context"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	kerrors "github.com/go-kratos/kratos/v2/errors"
)

func (uc *AuthorizationUsecase) SaveRoleMenu(ctx context.Context, roleId int64, menuIds []int64) error {
	if roleId == 0 {
		return kerrors.BadRequest(errResponse.ReasonParamsError, "角色id不得为空")
	}
	return uc.repo.SaveRoleMenu(ctx, roleId, menuIds)
}

func (uc *AuthorizationUsecase) GetRoleMenu(ctx context.Context, role string) ([]*Menu, error) {
	result, err := uc.repo.GetRoleMenu(ctx, role)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) GetRoleMenuTree(ctx context.Context, role string) ([]*Menu, error) {
	result, err := uc.repo.GetRoleMenuTree(ctx, role)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) GetRoleMenuBtn(ctx context.Context, roleId int64, roleName string, menuId int64) ([]*MenuBtn, error) {
	if roleId == 0 && roleName == "" {
		return nil, kerrors.BadRequest(errResponse.ReasonParamsError, "角色信息不得为空")
	}
	return uc.repo.GetRoleMenuBtn(ctx, roleId, roleName, menuId)
}

func (uc *AuthorizationUsecase) SetRoleMenuBtn(ctx context.Context, roleId int64, menuId int64, btnIds []int64) error {
	if roleId == 0 {
		return kerrors.BadRequest(errResponse.ReasonParamsError, "角色id不得为空")
	}
	if menuId == 0 {
		return kerrors.BadRequest(errResponse.ReasonParamsError, "菜单id不得为空")
	}
	return uc.repo.SetRoleMenuBtn(ctx, roleId, menuId, btnIds)
}
