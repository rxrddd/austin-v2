package service

import (
	"github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminInterface)

type AdminInterface struct {
	v1.UnimplementedAdminServer
	administratorUseCase *biz.AdministratorUseCase
	authUseCase          *biz.AuthUseCase
	GoodsUseCase         *biz.GoodsUseCase
	log                  *log.Helper
}

func NewAdminInterface(
	administratorUseCase *biz.AdministratorUseCase,
	authUseCase *biz.AuthUseCase,
	goodsCase *biz.GoodsUseCase,
	logger log.Logger) *AdminInterface {
	return &AdminInterface{
		log:                  log.NewHelper(log.With(logger, "module", "service/interface")),
		administratorUseCase: administratorUseCase,
		authUseCase:          authUseCase,
		GoodsUseCase:         goodsCase,
	}
}
