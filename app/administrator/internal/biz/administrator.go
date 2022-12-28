package biz

import (
	v1 "austin-v2/api/administrator/v1"
	"austin-v2/pkg/errResponse"
	"context"
	"net/http"

	"austin-v2/pkg/utils/typeConvert"
	"austin-v2/pkg/validate"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Administrator struct {
	Id            int64
	Username      string `validate:"required,max=50" label:"用户名"`
	Password      string `validate:"required,max=50" label:"密码"`
	Salt          string
	Mobile        string `validate:"required,max=20" label:"手机号码"`
	Nickname      string `validate:"required,max=50" label:"昵称"`
	Avatar        string `validate:"required,max=150" label:"头像地址"`
	Status        int64  `validate:"required,oneof=1 2" label:"状态"`
	Role          string
	LastLoginTime string
	LastLoginIp   string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
}

// AdministratorRepo 模块接口
type AdministratorRepo interface {
	CreateAdministrator(ctx context.Context, reqData *Administrator) (*Administrator, error)
	UpdateAdministrator(ctx context.Context, reqData *Administrator) (*Administrator, error)
	GetAdministrator(ctx context.Context, params map[string]interface{}) (*Administrator, error)
	ListAdministrator(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*Administrator, int64, error)
	DeleteAdministrator(ctx context.Context, id int64) error
	RecoverAdministrator(ctx context.Context, id int64) error
	VerifyAdministratorPassword(ctx context.Context, id int64, password string) (bool, error)
	UpdateAdministratorLoginInfo(ctx context.Context, id int64, loginTime string, loginIp string) error
	AdministratorStatusChange(ctx context.Context, id int64, status int64) error
}

type AdministratorUseCase struct {
	repo AdministratorRepo
	log  *log.Helper
}

func NewAdministratorUseCase(repo AdministratorRepo, logger log.Logger) *AdministratorUseCase {
	return &AdministratorUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer"))}
}

func (uc *AdministratorUseCase) Create(ctx context.Context, data *Administrator) (*Administrator, error) {
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Administrator{}, errors.New(http.StatusBadRequest, errResponse.ReasonParamsError, err.Error())
	}
	return uc.repo.CreateAdministrator(ctx, data)
}

func (uc *AdministratorUseCase) Delete(ctx context.Context, id int64) error {
	if id == 0 {
		return errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	return uc.repo.DeleteAdministrator(ctx, id)
}

func (uc *AdministratorUseCase) Recover(ctx context.Context, id int64) error {
	if id == 0 {
		return errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	return uc.repo.RecoverAdministrator(ctx, id)
}

func (uc *AdministratorUseCase) AdministratorStatusChange(ctx context.Context, id int64, status int64) error {
	if id == 0 {
		return errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}
	return uc.repo.AdministratorStatusChange(ctx, id, status)
}

func (uc *AdministratorUseCase) Update(ctx context.Context, data *Administrator) (*Administrator, error) {
	if data.Id == 0 {
		return &Administrator{}, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingId)
	}

	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Administrator{}, errors.New(http.StatusBadRequest, errResponse.ReasonParamsError, err.Error())
	}
	return uc.repo.UpdateAdministrator(ctx, data)
}

func (uc *AdministratorUseCase) Get(ctx context.Context, params map[string]interface{}) (*Administrator, error) {
	params = typeConvert.ClearMapZeroValue(params)
	return uc.repo.GetAdministrator(ctx, params)
}

func (uc *AdministratorUseCase) List(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*Administrator, int64, error) {
	params = typeConvert.ClearMapZeroValue(params)
	return uc.repo.ListAdministrator(ctx, page, pageSize, params)
}

func (uc *AdministratorUseCase) VerifyAdministratorPassword(ctx context.Context, in *v1.VerifyAdministratorPasswordRequest) (bool, error) {

	result, err := uc.repo.VerifyAdministratorPassword(ctx, in.Id, in.Password)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (uc *AdministratorUseCase) UpdateAdministratorLoginInfo(ctx context.Context, in *v1.AdministratorLoginSuccessRequest) error {
	return uc.repo.UpdateAdministratorLoginInfo(ctx, in.Id, in.LastLoginTime, in.LastLoginIp)

}
