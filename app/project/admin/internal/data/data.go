package data

import (
	administratorV1 "austin-v2/api/administrator/v1"
	authorizationV1 "austin-v2/api/authorization/v1"
	filesServiceV1 "austin-v2/api/files/v1"
	msgpushermanagerV1 "austin-v2/api/msgpusher-manager/v1"
	"austin-v2/app/project/admin/internal/conf"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	etcdclient "go.etcd.io/etcd/client/v3"
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
	NewFilesRepo,
	NewMsgPusherManagerRepo,
	NewAdministratorServiceClient,
	NewAuthorizationServiceClient,
	NewFilesServiceClient,
	NewMsgPusherManagerClient,
)

var auth *conf.Auth

var RedisCli redis.Cmdable

func GetAuthApiKey() string {
	return auth.ApiKey
}

// Data .
type Data struct {
	log                    *log.Helper
	redisCli               redis.Cmdable
	administratorClient    administratorV1.AdministratorClient
	authorizationClient    authorizationV1.AuthorizationClient
	filesClient            filesServiceV1.FilesClient
	msgPusherManagerClient msgpushermanagerV1.MsgPusherManagerClient
}

// NewData .
func NewData(
	logger log.Logger,
	redisCli redis.Cmdable,
	administratorClient administratorV1.AdministratorClient,
	authorizationClient authorizationV1.AuthorizationClient,
	filesClient filesServiceV1.FilesClient,
	msgPusherManagerClient msgpushermanagerV1.MsgPusherManagerClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
		log:                    l,
		redisCli:               redisCli,
		administratorClient:    administratorClient,
		authorizationClient:    authorizationClient,
		filesClient:            filesClient,
		msgPusherManagerClient: msgPusherManagerClient,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	point := conf.Etcd.Address
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{point},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	point := conf.Etcd.Address
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{point},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
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
	err := client.Ping(context.Background()).Err()
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}
	RedisCli = client
	return client
}
