package service

import (
	"github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminInterface)

type AdminInterface struct {
	v1.UnimplementedAdminServer
	administratorRepo *data.AdministratorRepo
	authorizationRepo *data.AuthorizationRepo
	log               *log.Helper
}

func NewAdminInterface(
	administratorRepo *data.AdministratorRepo,
	authorizationRepo *data.AuthorizationRepo,
	logger log.Logger) *AdminInterface {
	return &AdminInterface{
		log:               log.NewHelper(log.With(logger, "module", "service/interface")),
		administratorRepo: administratorRepo,
		authorizationRepo: authorizationRepo,
	}
}
