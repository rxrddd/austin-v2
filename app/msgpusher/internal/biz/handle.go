package biz

import (
	"austin-v2/app/msgpusher/internal/sender"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type HandleUsecase struct {
	log *log.Helper
	hs  []sender.Handler
}

func NewHandleUsecase(logger log.Logger, hs []sender.Handler) *HandleUsecase {
	return &HandleUsecase{
		log: log.NewHelper(logger),
		hs:  hs,
	}
}
func (a HandleUsecase) Handle(ctx context.Context, str string) error {
	for _, h := range a.hs {
		fmt.Println(h.Name())
	}
	return nil
}
