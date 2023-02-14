package main

import (
	"austin-v2/app/project/admin/internal/conf"
	"austin-v2/pkg/mode"
	"austin-v2/pkg/utils/httpHelper"
	"austin-v2/pkg/utils/stringHelper"
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"

	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// FlagConf is the config flag.
	FlagConf string

	ID string
)

func init() {
	flag.StringVar(&FlagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")

	json.MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true, //默认值不忽略
		UseProtoNames:   true, //使用proto name返回http字段
	}
}

func newApp(logger log.Logger, debug *conf.Debug, hs *http.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
		kratos.AfterStart(func(ctx context.Context) error {
			if debug.Mode == mode.DEV {
				httpHelper.PrintRoute(hs)
			}
			return nil
		}),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(FlagConf),
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

	Name = bc.Service.Name
	Version = bc.Service.Version

	hostname, _ := os.Hostname()
	ID = hostname + "." + bc.Service.Name + "." + Version + "." + stringHelper.GenerateUUID()

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}
	logger := log.With(log.NewStdLogger(os.Stdout),
		"service.name", Name,
		"service.version", Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	app, cleanup, err := wireApp(bc.Server, bc.Debug, &rc, bc.Data, bc.Auth, bc.Service, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
