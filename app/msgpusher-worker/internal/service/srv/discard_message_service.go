package srv

import (
	"austin-v2/pkg/types"
	"austin-v2/utils/arrayHelper"
	"austin-v2/utils/transformHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type DiscardMessageService struct {
	logger *log.Helper
	rds    redis.Cmdable
}

const discardMessageKey = "discard_message"

func NewDiscardMessageService(logger log.Logger, rds redis.Cmdable) *DiscardMessageService {
	return &DiscardMessageService{
		logger: log.NewHelper(log.With(logger, "module", "service/discard-message-service")),
		rds:    rds,
	}
}

// IsDiscard 根据redis配置丢弃某个模板的所有消息
func (l *DiscardMessageService) IsDiscard(ctx context.Context, taskInfo *types.TaskInfo) bool {
	discardMessageTemplateIds, err := l.rds.SMembers(ctx, discardMessageKey).Result()
	if err != nil {
		l.logger.Errorf("discard message service smembers err: %v ", err)
		return false
	}
	if len(discardMessageTemplateIds) == 0 {
		return false
	}
	if arrayHelper.ArrayInt64In(transformHelper.ArrayStringToInt64(discardMessageTemplateIds), taskInfo.MessageTemplateId) {
		return true
	}
	return false
}
