package limit

import (
	"austin-v2/pkg/types"
	"errors"
)

type LimiterManager struct {
	lMap map[string]types.ILimitService
}

func NewLimiterManager(
	sl *SimpleLimitService,
) *LimiterManager {
	h := &LimiterManager{}
	h.register(sl)
	return h
}

func (hs *LimiterManager) Route(code string) (types.ILimitService, error) {
	if h, ok := hs.lMap[code]; ok {
		return h, nil
	}
	return nil, errors.New("unknown handle")
}

func (hs *LimiterManager) register(h types.ILimitService) {
	if hs.lMap == nil {
		hs.lMap = make(map[string]types.ILimitService)
	}
	hs.lMap[h.Name()] = h
}
