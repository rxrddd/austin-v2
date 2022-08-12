package biz

import (
	"context"
	v1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/data/entity"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrRecordNotFound       = "管理员不存在"
	ErrAdministratorForbid  = "管理员已禁用"
	ErrAdministratorDeleted = "管理员已删除"
)

type Administrator struct {
	Id        int64
	Username  string
	Password  string
	Mobile    string
	Nickname  string
	Avatar    string
	Status    int64
	IsDeleted bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type AdministratorRepo interface {
	VerifyPassword(ctx context.Context, id int64, password string) (bool, error)
	CreateAdministrator(ctx context.Context, reqData *Administrator) (*Administrator, error)
	UpdateAdministrator(ctx context.Context, reqData *Administrator) (*Administrator, error)
	GetAdministrator(ctx context.Context, params map[string]interface{}) (*Administrator, error)
	ListAdministrator(ctx context.Context, pageNum, pageSize int64) ([]*Administrator, int64, error)
	DeleteAdministrator(ctx context.Context, id int64) error
}

type AdministratorUseCase struct {
	repo AdministratorRepo
	log  *log.Helper
}

func NewAdministratorUseCase(repo AdministratorRepo, logger log.Logger) *AdministratorUseCase {
	logs := log.NewHelper(log.With(logger, "module", "administrator/interface"))
	return &AdministratorUseCase{
		repo: repo,
		log:  logs,
	}
}

func (ac AdministratorUseCase) FindLoginAdministratorByUsername(ctx context.Context, in *v1.GetLoginAdministratorByUsernameRequest) (*v1.GetLoginAdministratorByUsernameReply, error) {
	params := make(map[string]interface{})
	params["username"] = in.Username
	administrator, err := ac.repo.GetAdministrator(ctx, params)
	if err != nil {
		return &v1.GetLoginAdministratorByUsernameReply{}, err
	}
	if administrator.Status == entity.AdministratorStatusForbid {
		return &v1.GetLoginAdministratorByUsernameReply{}, errors.New(400, "ADMINISTRATOR_FORBIDDEN", ErrAdministratorForbid)
	}
	if administrator.DeletedAt != "" {
		return &v1.GetLoginAdministratorByUsernameReply{}, errors.New(400, "ADMINISTRATOR_DELETED", ErrAdministratorDeleted)
	}

	return &v1.GetLoginAdministratorByUsernameReply{
		Id:       administrator.Id,
		Username: administrator.Username,
	}, nil
}

func (ac AdministratorUseCase) VerifyAdministratorPassword(ctx context.Context, in *v1.VerifyPasswordRequest) (bool, error) {
	result, err := ac.repo.VerifyPassword(ctx, in.Id, in.Password)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (uc *AdministratorUseCase) Create(ctx context.Context, data *Administrator) (*Administrator, error) {
	return uc.repo.CreateAdministrator(ctx, data)
}

func (uc *AdministratorUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteAdministrator(ctx, id)
}

func (uc *AdministratorUseCase) Update(ctx context.Context, data *Administrator) (*Administrator, error) {
	return uc.repo.UpdateAdministrator(ctx, data)
}

func (uc *AdministratorUseCase) Get(ctx context.Context, params map[string]interface{}) (*Administrator, error) {
	return uc.repo.GetAdministrator(ctx, params)
}

func (uc *AdministratorUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Administrator, int64, error) {
	return uc.repo.ListAdministrator(ctx, pageNum, pageSize)
}
