package data

import (
	"austin-v2/app/msgpusher-worker/internal/conf"
	"context"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/rabbitmq"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// ProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewBroker,
	NewData,
	NewRedisCmd,
	NewMysqlCmd,
	NewMessageTemplateRepo,
)

// Data .
type Data struct {
	broker broker.Broker
	rds    redis.Cmdable
	db     *gorm.DB
}

// NewData .
func NewData(
	c *conf.Data,
	logger log.Logger,
	broker broker.Broker,
	rds redis.Cmdable,
	db *gorm.DB,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		_ = broker.Disconnect()
		s, _ := db.DB()
		s.Close()
	}
	return &Data{
		broker: broker,
		rds:    rds,
		db:     db,
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
