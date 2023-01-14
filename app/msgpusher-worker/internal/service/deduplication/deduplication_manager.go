package deduplication

import (
	"austin-v2/pkg/manager"
	"austin-v2/pkg/types"
)

type DeduplicationManager struct {
	manager *manager.Manager
	mMap    map[string]types.IDeduplicationService
}

func NewDeduplicationManager(
	fds *FrequencyDeduplicationService,
	cds *ContentDeduplicationService,
) *DeduplicationManager {
	dm := &DeduplicationManager{
		manager: manager.NewManager(
			fds,
			cds,
		),
	}
	return dm
}

func (dm *DeduplicationManager) Get(key string) (resp types.IDeduplicationService, err error) {
	if h, err := dm.manager.Get(key); err == nil {
		return h.(types.IDeduplicationService), nil
	}
	return nil, err
}
