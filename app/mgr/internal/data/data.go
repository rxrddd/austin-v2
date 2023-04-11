package data

import (
	msgpushermanagerV1 "austin-v2/api/msgpusher-manager/v1"
	"austin-v2/common/dal/query"
	"austin-v2/pkg/transaction"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hibiken/asynq"
	etcdclient "go.etcd.io/etcd/client/v3"
	"time"

	"austin-v2/app/mgr/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var RedisCli redis.Cmdable

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	transaction.NewTranMgr,
	NewMsgPusherManagerClient,
	NewData,
	NewRedisCmd,
	NewMysqlCmd,
	NewAsynqServer,
	NewAsynqClient,
	NewDiscovery,

	NewMsgPusherManagerRepo,
	NewMenuRepo,
	NewRoleRepo,
	NewAdminRepo,
)

// Data .
type Data struct {
	db       *gorm.DB
	redisCli redis.Cmdable
	log      *log.Helper
	txMgr    transaction.ITranMgr

	msgPusherManagerClient msgpushermanagerV1.MsgPusherManagerClient
}

// NewData .
func NewData(db *gorm.DB,
	redisCmd redis.Cmdable,
	logger log.Logger,
	txMgr transaction.ITranMgr,
	msgPusherManagerClient msgpushermanagerV1.MsgPusherManagerClient,
) (*Data, func(), error) {
	logs := log.NewHelper(log.With(logger, "module", "mgr/data"))

	d := &Data{
		db:                     db,
		redisCli:               redisCmd,
		txMgr:                  txMgr,
		msgPusherManagerClient: msgPusherManagerClient,
	}
	return d, func() {
		logs.Error("mgr/data run failed")
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

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) DB(ctx context.Context) *gorm.DB {
	return d.txMgr.DB(ctx)
}

// Query 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) Query(ctx context.Context) *query.Query {
	return d.txMgr.Query(ctx)
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

func NewMysqlCmd(conf *conf.Data, logger log.Logger) *gorm.DB {
	logs := log.NewHelper(log.With(logger, "module", "gvs-mgr/data/mysql"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		logs.Fatalf("mysql connect error: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logs.Fatalf("mysql connect error: %v", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(100)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db.Debug()
}

func NewAsynqServer(conf *conf.Data) *asynq.Server {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password,
			DB:       0,
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
