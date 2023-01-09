package data

import (
	"austin-v2/app/msgpusher/internal/conf"
	"austin-v2/pkg/mq"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	etcdclient "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewMq,
	NewData,
	NewMessageTemplateRepo,
	NewRegistrar,
	NewRedisCmd,
	NewMysqlCmd,
)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(
	c *conf.Data,
	logger log.Logger,
	mq mq.IMessagingClient,
	db *gorm.DB,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		mq.Close()
	}
	return &Data{
		db: db,
	}, cleanup, nil
}

// NewMq .
func NewMq(c *conf.Data, logger log.Logger) mq.IMessagingClient {
	logs := log.NewHelper(log.With(logger, "module", "msgpusher-worker/data/mq"))
	client, err := mq.NewMessagingClientURL(c.Rabbitmq.URL)
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}
	return client
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
