package deduplication

import (
	limit "austin-v2/app/msgpusher-worker/internal/service/limiter"
	"austin-v2/pkg/types"
	"austin-v2/utils/encrypt"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

type ContentDeduplicationService struct {
	limit types.ILimitService

	deduplicationService
}

func NewContentDeduplicationService(manager *limit.LimiterManager) *ContentDeduplicationService {
	l, err := manager.Get(types.LimitSimple)
	if err != nil {
		panic(err)
	}
	return &ContentDeduplicationService{
		limit: l,
	}
}

func (c ContentDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param types.DeduplicationConfigItem) error {
	return c.deduplicationService.Deduplication(ctx, c.limit, c, taskInfo, param)
}

func (c ContentDeduplicationService) DeduplicationSingleKey(taskInfo *types.TaskInfo, receiver string) string {
	str, _ := json.Marshal(taskInfo.ContentModel)
	return encrypt.MD5(cast.ToString(taskInfo.MessageTemplateId) + receiver + string(str))
}
func (c ContentDeduplicationService) Name() string {
	return deduplicationPrefix + Content
}
