package data

import (
	administratorV1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
	authorizationV1 "github.com/ZQCard/kratos-base-project/api/authorization/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/conf"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRedisCmd,
	NewRegistrar,
	NewDiscovery,
	NewAdministratorRepo,
	NewAuthorizationRepo,
	NewAdministratorServiceClient,
	NewAuthorizationServiceClient,
)

var auth *conf.Auth

var RedisCli redis.Cmdable

func GetAuthApiKey() string {
	return auth.ApiKey
}

// Data .
type Data struct {
	log                 *log.Helper
	redisCli            redis.Cmdable
	administratorClient administratorV1.AdministratorClient
	authorizationClient authorizationV1.AuthorizationClient
}

// NewData .
func NewData(
	logger log.Logger,
	redisCli redis.Cmdable,
	administratorClient administratorV1.AdministratorClient,
	authorizationClient authorizationV1.AuthorizationClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
		log:                 l,
		redisCli:            redisCli,
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

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	logs := log.NewHelper(log.With(logger, "module", "administrator-service/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	err := client.Ping().Err()
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}
	RedisCli = client
	return client
}
