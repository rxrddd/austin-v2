package limit

import (
	"austin-v2/pkg/manager"
	"austin-v2/pkg/types"
)

type LimiterManager struct {
	manager *manager.Manager
}

func NewLimiterManager(
	sl *SimpleLimitService,
	sw *SlideWindowLimitService,
) *LimiterManager {
	return &LimiterManager{
		manager: manager.NewManager(
			sl,
			sw,
		),
	}
}
func (lm *LimiterManager) Get(key string) (resp types.ILimitService, err error) {
	if h, err := lm.manager.Get(key); err == nil {
		return h.(types.ILimitService), nil
	}
	return nil, err
}
