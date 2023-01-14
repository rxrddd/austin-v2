package sender

import (
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/pkg/manager"
	"austin-v2/pkg/types"
)

type HandleManager struct {
	manager *manager.Manager
}

func NewHandleManager(
	sms *handler.SmsHandler,
	email *handler.EmailHandler,
	officialAccount *handler.OfficialAccountHandler,
) *HandleManager {
	hm := &HandleManager{
		manager: manager.NewManager(
			sms,
			email,
			officialAccount,
		),
	}
	return hm
}

func (hm *HandleManager) Get(key string) (resp types.IHandler, err error) {
	if h, err := hm.manager.Get(key); err != nil {
		return h.(types.IHandler), nil
	}
	return nil, err
}
