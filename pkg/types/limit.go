package types

import (
	"context"
)

type ILimitService interface {
	LimitFilter(ctx context.Context, duplication IDeduplicationService, taskInfo *TaskInfo, param DeduplicationConfigItem) (filterReceiver []string, err error)
	Name() string
}
type IDeduplicationService interface {
	Deduplication(ctx context.Context, taskInfo *TaskInfo, param DeduplicationConfigItem) error
	DeduplicationSingleKey(taskInfo *TaskInfo, receiver string) string
	Name() string
}

const LimitSimple = "SimpleLimitService"
const LimitSlideWindow = "SlideWindowLimitService"
