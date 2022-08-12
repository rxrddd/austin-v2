package data

import (
	"context"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/conf"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	administratorClientV1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
	goodsClientV1 "github.com/ZQCard/kratos-base-project/api/goods/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRegistrar,
	NewDiscovery,
	NewAdministratorServiceClient,
	NewAdministratorRepo,
	NewGoodsClient,
	NewGoodsRepo,
)

// Data .
type Data struct {
	log                 *log.Helper
	administratorClient administratorClientV1.AdministratorClient
	goodsClient         goodsClientV1.GoodsClient
}

// NewData .
func NewData(
	conf *conf.Data,
	logger log.Logger,
	administratorClient administratorClientV1.AdministratorClient,
	goodsClient goodsClientV1.GoodsClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
		log:                 l,
		administratorClient: administratorClient,
		goodsClient:         goodsClient,
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

func NewAdministratorServiceClient(ac *conf.Auth, sr *conf.Service, r registry.Discovery, tp *tracesdk.TracerProvider) administratorClientV1.AdministratorClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Administrator.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			//jwt.Client(func(token *jwt2.Token) (interface{}, error) {
			//	return []byte(ac.ServiceKey), nil
			//}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := administratorClientV1.NewAdministratorClient(conn)
	return c
}

func NewGoodsClient(ac *conf.Auth, sr *conf.Service, r registry.Discovery, tp *tracesdk.TracerProvider) goodsClientV1.GoodsClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Goods.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			//jwt.Client(func(token *jwt2.Token) (interface{}, error) {
			//	return []byte(ac.ServiceKey), nil
			//}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := goodsClientV1.NewGoodsClient(conn)
	return c
}
