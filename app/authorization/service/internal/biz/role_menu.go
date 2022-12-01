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
	if len(menuIds) == 0 {
		return kerrors.BadRequest(errResponse.ReasonParamsError, "菜单id不得为空")
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
