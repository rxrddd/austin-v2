package smsScript

import (
	"austin-v2/pkg/manager"
	"austin-v2/pkg/types"
)

type SmsManager struct {
	manager *manager.Manager
}

func NewSmsManager(
	yunpian *YunPian,
	aliyun *AliyunSms,
) *SmsManager {
	sm := &SmsManager{
		manager: manager.NewManager(
			yunpian,
			aliyun,
		),
	}
	return sm
}

func (hm *SmsManager) Get(key string) (resp types.ISmsScript, err error) {
	if h, err := hm.manager.Get(key); err != nil {
		return h.(types.ISmsScript), nil
	}
	return nil, err
}
