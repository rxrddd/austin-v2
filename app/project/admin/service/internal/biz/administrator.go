package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Administrator struct {
	Id        int64
	Username  string
	Mobile    string
	Nickname  string
	Avatar    string
	Status    int64
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type AdministratorRepo interface {
	FindLoginAdministratorByUsername(ctx context.Context, username string) (*Administrator, error)
	GetAdministrator(ctx context.Context, id int64) (*Administrator, error)
	VerifyPassword(ctx context.Context, id int64, password string) error
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


func (receiver *AdministratorUseCase) GetAdministrator(ctx context.Context, id int64) (*Administrator, error) {
	return receiver.repo.GetAdministrator(ctx, id)
}
