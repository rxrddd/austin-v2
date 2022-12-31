package deduplication

import (
	"austin-v2/pkg/types"
	"errors"
)

type DeduplicationManager struct {
	lMap map[string]types.IDeduplicationService
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
	if h, ok := hs.lMap[code]; ok {
		return h, nil
	}
	return nil, errors.New("unknown handle")
}

func (hs *DeduplicationManager) register(h types.IDeduplicationService) {
	if hs.lMap == nil {
		hs.lMap = make(map[string]types.IDeduplicationService)
	}
	hs.lMap[h.Name()] = h
}
