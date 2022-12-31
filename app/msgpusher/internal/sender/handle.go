package sender

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
)

type Handle struct {
	hs map[string]Handler

	svcCtx *ServiceContext
}

type ServiceContext struct {
	log    *log.Helper
	broker broker.Broker
}

func NewHandle(
	logger log.Logger,
	broker broker.Broker,
) *Handle {

	h := &Handle{
		svcCtx: &ServiceContext{
			log:    log.NewHelper(logger),
			broker: broker,
		},
	}

	h.registerHandler(NewSms())

	return h
}

func (a *Handle) registerHandler(h Handler) {
	if a.hs == nil {
		a.hs = make(map[string]Handler)
	}
	a.hs[h.Name()] = h
}

func (a *Handle) Handle(ctx context.Context, str string) error {
	if h, ok := a.hs[str]; ok {
		return h.Handle(ctx)
	}
	return errors.New("unknown handle")
}

type Handler interface {
	Name() string
	Handle(ctx context.Context) error
}
