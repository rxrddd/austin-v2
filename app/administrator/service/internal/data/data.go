package data

import (
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-redis/redis"
	etcdclient "go.etcd.io/etcd/client/v3"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewRedisCmd,
	NewMysqlCmd,
	NewAdministratorRepo,
)

// Data .
type Data struct {
	db       *gorm.DB
	redisCli redis.Cmdable
	log      *log.Helper
}

// NewData .
func NewData(db *gorm.DB, redisCmd redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	logs := log.NewHelper(log.With(logger, "module", "administrator-service/data"))

	d := &Data{
		db:       db.Debug(),
		redisCli: redisCmd,
	}
	return d, func() {
		logs.Error("administrator-service/data failed")
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
	err := client.Ping().Err()
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
