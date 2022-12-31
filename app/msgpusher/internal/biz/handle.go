package biz

import (
	"austin-v2/app/msgpusher/internal/sender"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type HandleUsecase struct {
	log *log.Helper
	hs  *sender.Handle
}

func NewHandleUsecase(logger log.Logger, hs *sender.Handle) *HandleUsecase {
	return &HandleUsecase{
		log: log.NewHelper(logger),
		hs:  hs,
	}
}
func (a HandleUsecase) Handle(ctx context.Context, str string) error {
	return a.hs.Handle(ctx, str)
}
