package data

import (
	"austin-v2/app/msgpusher/internal/conf"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/rabbitmq"
	etcdclient "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewBroker,
	NewData,
	NewGreeterRepo,
	NewDiscovery,
	NewRegistrar,
	NewRedisCmd,
	NewMysqlCmd,
)

// Data .
type Data struct {
	// TODO wrapped database client
	broker broker.Broker
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, broker broker.Broker) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		_ = broker.Disconnect()
	}
	return &Data{
		broker: broker,
	}, cleanup, nil
}

// NewBroker .
func NewBroker(c *conf.Data, logger log.Logger) broker.Broker {

	ctx := context.Background()
	b := rabbitmq.NewBroker(
		broker.WithOptionContext(ctx),
		broker.WithAddress(c.Rabbitmq.URL),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		log.Error(`err`, err)
		panic(err)
	}
	log.NewHelper(logger).Info("NewBroker " + b.Name())
	return b
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
	return client
}

func NewMysqlCmd(conf *conf.Data, logger log.Logger) *gorm.DB {
	logs := log.NewHelper(log.With(logger, "module", "administrator-service/data/mysql"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		logs.Fatalf("mysql connect error: %v", err)
	}
	return db
}
