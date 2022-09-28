package data

import (
	administratorV1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
	authorizationV1 "github.com/ZQCard/kratos-base-project/api/authorization/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/conf"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRegistrar,
	NewDiscovery,
	NewAdministratorRepo,
	NewAuthorizationRepo,
	NewAdministratorServiceClient,
	NewAuthorizationServiceClient,
)

// Data .
type Data struct {
	log                 *log.Helper
	administratorClient administratorV1.AdministratorClient
	authorizationClient authorizationV1.AuthorizationClient
}

// NewData .
func NewData(
	logger log.Logger,
	administratorClient administratorV1.AdministratorClient,
	authorizationClient authorizationV1.AuthorizationClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
		log:                 l,
		administratorClient: administratorClient,
		authorizationClient: authorizationClient,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
