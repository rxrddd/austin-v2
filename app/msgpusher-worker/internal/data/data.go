package data

import (
	"austin-v2/app/msgpusher-worker/internal/conf"
	"austin-v2/pkg/mq"
	"context"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// ProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewMq,
	NewData,
	NewRedisCmd,
	NewMysqlCmd,
	NewMessageTemplateRepo,
	NewSendAccountRepo,
	NewSmsRecordRepo,
)

// Data .
type Data struct {
	rds redis.Cmdable
	db  *gorm.DB
}

// NewData .
func NewData(
	_ *conf.Data,
	mq mq.IMessagingClient,
	logger log.Logger,
	rds redis.Cmdable,
	db *gorm.DB,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		s, _ := db.DB()
		_ = s.Close()
		mq.Close()
	}
	return &Data{
		rds: rds,
		db:  db,
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

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	logs := log.NewHelper(log.With(logger, "module", "msgpusher-worker/data/redis"))
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
	logs := log.NewHelper(log.With(logger, "module", "msgpusher-worker/data/mysql"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: logger2.Default.LogMode(logger2.Info),
	})
	if err != nil {
		logs.Fatalf("mysql connect error: %v", err)
	}
	return db
}
