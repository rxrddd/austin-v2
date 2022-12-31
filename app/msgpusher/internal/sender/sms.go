package sender

import (
	"context"
	"fmt"
)

type Sms struct {
	svcCtx *ServiceContext
}

func NewSms(svcCtx *ServiceContext) Handler {
	return &Sms{
		svcCtx: svcCtx,
	}
}
func (a Sms) Name() string {
	return "sms"
}
func (a Sms) Handle(ctx context.Context) error {
	fmt.Println("sms send")
	return nil
}
