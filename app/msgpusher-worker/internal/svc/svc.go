package svc

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
)

type ServiceContext struct {
	Logger *log.Helper
	Broker broker.Broker
}
