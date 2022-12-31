package sender

import (
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/pkg/types"
	"errors"
)

type HandleManager struct {
	hsMap map[string]types.IHandler
}

func NewHandleManager(
	sms *handler.SmsHandler,
	email *handler.EmailHandler,
) *HandleManager {
	h := &HandleManager{}
	h.register(sms)
	h.register(email)
	return h
}

func (hs *HandleManager) Route(channel string) (types.IHandler, error) {
	if h, ok := hs.hsMap[channel]; ok {
		return h, nil
	}
	return nil, errors.New("unknown handle")
}

func (hs *HandleManager) register(h types.IHandler) {
	if hs.hsMap == nil {
		hs.hsMap = make(map[string]types.IHandler)
	}
	hs.hsMap[h.Name()] = h
}
