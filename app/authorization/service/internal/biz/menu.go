package biz

import (
	"context"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/validate"
	kerrors "github.com/go-kratos/kratos/v2/errors"
)

type MenuBtn struct {
	Id          int64
	MenuId      int64  `validate:"gte=0" label:"菜单id"`
	Name        string `validate:"required,max=255" label:"按钮名称"`
	Description string `validate:"required,max=255" label:"按钮描述"`
	CreatedAt   string
	UpdatedAt   string
}

type Menu struct {
	Id        int64
	Name      string `validate:"required,max=50" label:"路由名称"`
	Path      string `validate:"required,max=255" label:"路由path"`
	ParentId  int64  `validate:"gte=0" label:"父级ID"`
	Hidden    int64  `validate:"oneof=0 1" label:"是否隐藏"`
	Component string `validate:"required,max=255" label:"前端文件路径"`
	Sort      int64  `validate:"gte=1" label:"菜单排序"`
	Title     string `validate:"required,max=255" label:"名称"`
	Icon      string `validate:"required,max=255" label:"icon图标"`
	CreatedAt string
	UpdatedAt string
	MenuBtns  []MenuBtn
	Children  []Menu
}

func (uc *AuthorizationUsecase) CreateMenu(ctx context.Context, data *Menu) (*Menu, error) {
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Menu{}, kerrors.BadRequest(errResponse.ReasonParamsError, err.Error())
	}
	for _, v := range data.MenuBtns {
		err := validate.ValidateStructCN(v)
		if err != nil {
			return &Menu{}, kerrors.BadRequest(errResponse.ReasonParamsError, err.Error())
		}
	}

	result, err := uc.repo.CreateMenu(ctx, data)
	if err != nil {
		return &Menu{}, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) DeleteMenu(ctx context.Context, id int64) error {
	if id == 0 {
		return errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	return uc.repo.DeleteMenu(ctx, id)
}

func (uc *AuthorizationUsecase) UpdateMenu(ctx context.Context, data *Menu) (*Menu, error) {
	if data.Id == 0 {
		return &Menu{}, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Menu{}, kerrors.BadRequest(errResponse.ReasonParamsError, err.Error())
	}
	for _, v := range data.MenuBtns {
		err := validate.ValidateStructCN(v)
		if err != nil {
			return &Menu{}, kerrors.BadRequest(errResponse.ReasonParamsError, err.Error())
		}
	}
	result, err := uc.repo.UpdateMenu(ctx, data)
	if err != nil {
		return &Menu{}, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) GetMenuTree(ctx context.Context) ([]*Menu, error) {
	result, err := uc.repo.GetMenuTree(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) GetMenuAll(ctx context.Context) ([]*Menu, error) {
	result, err := uc.repo.GetMenuAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
