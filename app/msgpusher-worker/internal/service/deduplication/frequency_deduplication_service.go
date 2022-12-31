package deduplication

import (
	limit "austin-v2/app/msgpusher-worker/internal/service/limiter"
	"austin-v2/pkg/types"
	"context"
	"fmt"
)

const frequencyDeduplicationServicePrefix = "FRE"

type FrequencyDeduplicationService struct {
	deduplicationService

	limit types.ILimitService
}

func NewFrequencyDeduplicationService(manager *limit.LimiterManager) *FrequencyDeduplicationService {
	l, err := manager.Route(types.LimitSlideWindow)
	if err != nil {
		panic(err)
	}
	return &FrequencyDeduplicationService{
		limit: l,
	}
}

func (c FrequencyDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param types.DeduplicationConfigItem) error {
	return c.deduplicationService.Deduplication(ctx, c.limit, c, taskInfo, param)
}

func (c FrequencyDeduplicationService) DeduplicationSingleKey(taskInfo *types.TaskInfo, receiver string) string {
	return fmt.Sprintf("%s_%s_%d_%d", frequencyDeduplicationServicePrefix, receiver, taskInfo.MessageTemplateId, taskInfo.SendChannel)
}
func (c FrequencyDeduplicationService) Name() string {
	return deduplicationPrefix + Frequency
}
