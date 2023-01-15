package service

import (
	"austin-v2/api/project/admin/v1"
	"austin-v2/app/project/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminInterface)

type AdminInterface struct {
	v1.UnimplementedAdminServer
	administratorRepo    *data.AdministratorRepo
	authorizationRepo    *data.AuthorizationRepo
	filesRepo            *data.FilesRepo
	msgPusherManagerRepo *data.MsgPusherManagerRepo

	log *log.Helper
}

func NewAdminInterface(
	administratorRepo *data.AdministratorRepo,
	authorizationRepo *data.AuthorizationRepo,
	filesRepo *data.FilesRepo,
	msgPusherManagerRepo *data.MsgPusherManagerRepo,
	logger log.Logger,
) *AdminInterface {
	return &AdminInterface{
		log:                  log.NewHelper(log.With(logger, "module", "service/interface")),
		administratorRepo:    administratorRepo,
		authorizationRepo:    authorizationRepo,
		filesRepo:            filesRepo,
		msgPusherManagerRepo: msgPusherManagerRepo,
	}
}
