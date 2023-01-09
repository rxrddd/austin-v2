package smsScript

import (
	"austin-v2/pkg/types"
	"errors"
)

type SmsManager struct {
	mMap map[string]types.ISmsScript
}

func NewSmsManager(
	yunpian *YunPian,
	aliyun *AliyunSms,
) *SmsManager {
	h := &SmsManager{}
	h.register(yunpian)
	h.register(aliyun)
	return h
}

func (hs *SmsManager) Route(code string) (types.ISmsScript, error) {
	if h, ok := hs.mMap[code]; ok {
		return h, nil
	}
	return nil, errors.New("unknown sms script " + code)
}

func (hs *SmsManager) register(h types.ISmsScript) {
	if hs.mMap == nil {
		hs.mMap = make(map[string]types.ISmsScript)
	}
	hs.mMap[h.Name()] = h
}
