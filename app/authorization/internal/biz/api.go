package biz

import (
	"context"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/validate"
	kerrors "github.com/go-kratos/kratos/v2/errors"
)

type Api struct {
	Id        int64
	Group     string `validate:"required,max=100" label:"api分组名称"`
	Name      string `validate:"required,max=100" label:"api名称"`
	Method    string `validate:"required,max=100" label:"请求类型"`
	Path      string `validate:"required,max=100" label:"请求路径"`
	CreatedAt string
	UpdatedAt string
}

func (uc *AuthorizationUsecase) CreateApi(ctx context.Context, data *Api) (*Api, error) {
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Api{}, kerrors.BadRequest(errResponse.ReasonParamsError, err.Error())
	}
	result, err := uc.repo.CreateApi(ctx, data)
	if err != nil {
		return &Api{}, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) DeleteApi(ctx context.Context, id int64) error {
	if id == 0 {
		return errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	return uc.repo.DeleteApi(ctx, id)
}

func (uc *AuthorizationUsecase) UpdateApi(ctx context.Context, data *Api) (*Api, error) {
	if data.Id == 0 {
		return &Api{}, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Api{}, kerrors.BadRequest(errResponse.ReasonParamsError, err.Error())
	}
	result, err := uc.repo.UpdateApi(ctx, data)
	if err != nil {
		return &Api{}, err
	}
	return result, nil
}

func (uc *AuthorizationUsecase) GetApiList(ctx context.Context, page int64, pageSize int64, params map[string]interface{}) ([]*Api, int64, error) {
	result, count, err := uc.repo.GetApiList(ctx, page, pageSize, params)
	if err != nil {
		return nil, 0, err
	}
	return result, count, nil
}

func (uc *AuthorizationUsecase) GetApiAll(ctx context.Context) ([]*Api, error) {
	result, err := uc.repo.GetApiAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
