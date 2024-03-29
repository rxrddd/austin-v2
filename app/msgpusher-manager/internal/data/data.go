package data

import (
	"austin-v2/app/msgpusher-manager/internal/conf"
	"austin-v2/common/dal/query"
	"austin-v2/pkg/transaction"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	etcdclient "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	transaction.NewTranMgr,
	NewData,
	NewRegistrar,
	NewRedisCmd,
	NewMysqlCmd,
	NewMsgRecordRepo,
	NewMessageTemplateRepo,
	NewSmsRecordRepo,
	NewSendAccountRepo,
	NewWxTemplateRepo,
)

// Data .
type Data struct {
	db    *gorm.DB
	rds   redis.Cmdable
	txMgr transaction.ITranMgr
}

// NewData .
func NewData(
	c *conf.Data,
	logger log.Logger,
	db *gorm.DB,
	rds redis.Cmdable,
	txMgr transaction.ITranMgr,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:    db,
		rds:   rds,
		txMgr: txMgr,
	}, cleanup, nil
}

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) DB(ctx context.Context) *gorm.DB {
	return d.txMgr.DB(ctx)
}

// Query 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) Query(ctx context.Context) *query.Query {
	return d.txMgr.Query(ctx)
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
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: logger2.Default.LogMode(logger2.Info),
	})
	if err != nil {
		logs.Fatalf("mysql connect error: %v", err)
	}
	return db
}
