package types

import (
	"context"
)

type LimitService interface {
	LimitFilter(ctx context.Context, duplication DeduplicationService, taskInfo *TaskInfo, param DeduplicationConfigItem) (filterReceiver []string, err error)
	Name() string
}
type DeduplicationService interface {
	Deduplication(ctx context.Context, taskInfo *TaskInfo, param DeduplicationConfigItem) error
	DeduplicationSingleKey(taskInfo *TaskInfo, receiver string) string
}

const LimitSimple = "SimpleLimitService"
const LimitSlideWindow = "SlideWindowLimitService"
