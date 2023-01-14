package srv

import (
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/app/msgpusher-worker/internal/service/deduplication"
	"austin-v2/pkg/types"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type DeduplicationRuleService struct {
	logger               *log.Helper
	rds                  redis.Cmdable
	uc                   *biz.MessageTemplateUseCase
	deduplicationManager *deduplication.DeduplicationManager
}

func NewDeduplicationRuleService(
	logger log.Logger,
	rds redis.Cmdable,
	uc *biz.MessageTemplateUseCase,
	deduplicationManager *deduplication.DeduplicationManager,
) *DeduplicationRuleService {
	return &DeduplicationRuleService{
		logger:               log.NewHelper(log.With(logger, "module", "service/deduplication-rule-service")),
		rds:                  rds,
		uc:                   uc,
		deduplicationManager: deduplicationManager,
	}
}

func (l *DeduplicationRuleService) Duplication(ctx context.Context, taskInfo *types.TaskInfo) {

	// 配置样例：{"deduplication_10":{"num":1,"time":300},"deduplication_20":{"num":5}}
	one, err := l.uc.One(ctx, taskInfo.MessageTemplateId)
	if err != nil {
		l.logger.Errorf("deduplication rule 查询模板错误 err: %v", err)
		return
	}
	if one.DeduplicationConfig == "" {
		//没有配置去重策略直接不管
		return
	}
	var deduplicationConfig = make(map[string]types.DeduplicationConfigItem)
	err = json.Unmarshal([]byte(one.DeduplicationConfig), &deduplicationConfig)
	if err != nil {
		l.logger.Errorf("deduplication json 解析config err: %v", err)
		return
	}
	if len(deduplicationConfig) <= 0 {
		//没配置限流策略
		return
	}

	for key, value := range deduplicationConfig {
		route, err := l.deduplicationManager.Get(key)
		//表示没匹配到对于的执行器
		if err != nil {
			continue
		}

		if err = route.Deduplication(ctx, taskInfo, value); err != nil {
			l.logger.Errorf("deduplication rule exec err: %v", err)
		}
	}

}
