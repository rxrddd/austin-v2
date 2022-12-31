package deduplication

import (
	"austin-v2/pkg/types"
	"errors"
)

type DeduplicationManager struct {
	mMap map[string]types.IDeduplicationService
}

func NewDeduplicationManager(
	fds *FrequencyDeduplicationService,
	cds *ContentDeduplicationService,
) *DeduplicationManager {
	h := &DeduplicationManager{}
	h.register(fds)
	h.register(cds)
	return h
}

func (hs *DeduplicationManager) Route(code string) (types.IDeduplicationService, error) {
	if h, ok := hs.mMap[code]; ok {
		return h, nil
	}
	return nil, errors.New("unknown deduplication " + code)
}

func (hs *DeduplicationManager) register(h types.IDeduplicationService) {
	if hs.mMap == nil {
		hs.mMap = make(map[string]types.IDeduplicationService)
	}
	hs.mMap[h.Name()] = h
}
