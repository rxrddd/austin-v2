package handler

import (
	"austin-v2/pkg/manager"
	"austin-v2/pkg/types"
)

type HandleManager struct {
	manager *manager.Manager
}

func NewHandleManager(
	sms *SmsHandler,
	email *EmailHandler,
	officialAccount *OfficialAccountHandler,
) *HandleManager {
	return &HandleManager{
		manager: manager.NewManager(
			sms,
			email,
			officialAccount,
		),
	}
}

func (hm *HandleManager) Get(key string) (resp types.IHandler, err error) {
	if h, err := hm.manager.Get(key); err == nil {
		return h.(types.IHandler), nil
	}
	return nil, err
}
