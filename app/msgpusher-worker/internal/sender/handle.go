package sender

import (
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/app/msgpusher-worker/internal/svc"
	"austin-v2/pkg/types"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
)

type Handle struct {
	hs     map[string]types.IHandler
	svcCtx *svc.ServiceContext
}

func NewHandle(
	logger log.Logger,
	broker broker.Broker,
) *Handle {

	svcCtx := &svc.ServiceContext{
		Logger: log.NewHelper(logger),
		Broker: broker,
	}

	h := &Handle{
		svcCtx: svcCtx,
	}

	h.registerHandler(handler.NewSmsHandler(h.svcCtx))

	return h
}

func (a *Handle) registerHandler(h types.IHandler) {
	if a.hs == nil {
		a.hs = make(map[string]types.IHandler)
	}
	a.hs[h.Name()] = h
}

func (a *Handle) Route(str string) (types.IHandler, error) {
	if h, ok := a.hs[str]; ok {
		return h, nil
	}
	return nil, errors.New("unknown handle")
}
