package sender

import (
	"context"
	"fmt"
)

type Sms struct {
}

func NewSms() *Sms {
	return &Sms{}
}
func (a Sms) Name() string {
	return "sms"
}
func (a Sms) Handle(ctx context.Context) error {
	fmt.Println("sms send")
	return nil
}

type Sms2 struct {
}

func NewSms2() *Sms2 {
	return &Sms2{}
}
func (a Sms2) Name() string {
	return "sms2"
}
func (a Sms2) Handle(ctx context.Context) error {
	fmt.Println("sms send")
	return nil
}
