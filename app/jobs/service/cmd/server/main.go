package main

import (
	"flag"
	"github.com/ZQCard/kratos-base-project/app/jobs/service/conf"
	"github.com/ZQCard/kratos-base-project/app/jobs/service/initDB"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	// 启动监听
	go func() {
		initDB.ListenClearSignal(bc.Data)
	}()
	// 发起任务
	go func() {
		initDB.CallInitDB(bc.Data)
	}()
	select {}
}
