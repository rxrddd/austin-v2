package limit

import (
	"austin-v2/pkg/types"
	"errors"
)

type LimiterManager struct {
	mMap map[string]types.ILimitService
}

func NewLimiterManager(
	sl *SimpleLimitService,
	sw *SlideWindowLimitService,
) *LimiterManager {
	h := &LimiterManager{}
	h.register(sl)
	h.register(sw)
	return h
}

func (hs *LimiterManager) Route(code string) (types.ILimitService, error) {
	if h, ok := hs.mMap[code]; ok {
		return h, nil
	}
	return nil, errors.New("unknown limiter " + code)
}

func (hs *LimiterManager) register(h types.ILimitService) {
	if hs.mMap == nil {
		hs.mMap = make(map[string]types.ILimitService)
	}
	hs.mMap[h.Name()] = h
}
