package data

import (
	"austin-v2/app/msgpusher-worker/internal/conf"
	"context"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewData,
	NewRedisCmd,
	NewMysqlCmd,
	NewMessageTemplateRepo,
	NewSendAccountRepo,
	NewSmsRecordRepo,
	NewMysqlMsgRecordRepo,
	NewAsynqServer,
	NewAsynqClient,
	NewAsynqScheduler,
)

// Data .
type Data struct {
	rds redis.Cmdable
	db  *gorm.DB
}

// NewData .
func NewData(
	_ *conf.Data,
	logger log.Logger,
	rds redis.Cmdable,
	db *gorm.DB,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		s, _ := db.DB()
		_ = s.Close()
	}
	return &Data{
		rds: rds,
		db:  db,
	}, cleanup, nil
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
		//Logger: logger2.Default.LogMode(logger2.Info),
	})
	if err != nil {
		logs.Fatalf("mysql connect error: %v", err)
	}
	return db
}

func NewAsynqServer(conf *conf.Data) *asynq.Server {
	// 首先定义一个client
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password,
		},
		asynq.Config{
			Concurrency: 10, // Concurrency表示最大并发处理任务数。
		},
	)
	return srv
}
func NewAsynqClient(conf *conf.Data) *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
	})
	return client
}

func NewAsynqScheduler(conf *conf.Data) *asynq.Scheduler {
	scheduler := asynq.NewScheduler(asynq.RedisClientOpt{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
	}, &asynq.SchedulerOpts{})
	return scheduler
}
