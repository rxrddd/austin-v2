package service

import (
	v1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdministratorService)

type AdministratorService struct {
	v1.UnimplementedAdministratorServer
	administratorCase *biz.AdministratorUseCase
	log *log.Helper
}

func NewAdministratorService(
	administratorCase *biz.AdministratorUseCase,
	logger log.Logger) *AdministratorService {
	return &AdministratorService{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		administratorCase:  administratorCase,
	}
}

