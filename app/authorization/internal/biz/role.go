package biz

import (
	"context"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/validate"
	kerrors "github.com/go-kratos/kratos/v2/errors"
)

type Role struct {
	Id        int64
	Name      string  `validate:"required,max=50" label:"角色名称"`
	ParentId  int64   `validate:"gte=0" label:"父级ID"`
	ParentIds []int64 `validate:"required" label:"父级ID数组"`
	CreatedAt string
	UpdatedAt string
	Children  []Role
}

func (uc *AuthorizationUsecase) CreateRole(ctx context.Context, data *Role) (*Role, error) {
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Role{}, kerrors.BadRequest(errResponse.ReasonParamsError, err.Error())
	}
	result, err := uc.repo.CreateRole(ctx, data)
	if err != nil {
		return &Role{}, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) DeleteRole(ctx context.Context, id int64) error {
	if id == 0 {
		return errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	return uc.repo.DeleteRole(ctx, id)
}

func (uc *AuthorizationUsecase) UpdateRole(ctx context.Context, data *Role) (*Role, error) {
	if data.Id == 0 {
		return &Role{}, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Role{}, kerrors.BadRequest(errResponse.ReasonParamsError, err.Error())
	}
	result, err := uc.repo.UpdateRole(ctx, data)
	if err != nil {
		return &Role{}, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) GetRoleList(ctx context.Context) ([]*Role, error) {
	result, err := uc.repo.GetRoleList(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
