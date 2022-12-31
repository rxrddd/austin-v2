package service

import (
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/pkg/types"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

const Content = "10"   //N分钟相同内容去重
const Frequency = "20" //一天内N次相同渠道去重
const deduplicationPrefix = "deduplication_"

type DeduplicationRuleService struct {
	logger *log.Helper
	rds    redis.Cmdable
	uc     *biz.MessageTemplateUseCase
}

func NewDeduplicationRuleService(
	logger log.Logger,
	rds redis.Cmdable,
	uc *biz.MessageTemplateUseCase,
) *DeduplicationRuleService {
	return &DeduplicationRuleService{
		logger: log.NewHelper(log.With(logger, "module", "service/deduplication-rule-service")),
		rds:    rds,
		uc:     uc,
	}
}

func (l DeduplicationRuleService) Duplication(ctx context.Context, taskInfo *types.TaskInfo) {

	// 配置样例：{"deduplication_10":{"num":1,"time":300},"deduplication_20":{"num":5}}
	one, err := l.uc.One(ctx, taskInfo.MessageTemplateId)
	if err != nil {
		//logx.Errorw("DeduplicationRuleService 查询模板错误 err", logx.Field("err", err))
		return
	}
	if one.DeduplicationConfig == "" {
		//没有配置去重策略直接不管
		return
	}
	var deduplicationConfig = make(map[string]types.DeduplicationConfigItem)
	err = json.Unmarshal([]byte(one.DeduplicationConfig), &deduplicationConfig)
	if err != nil {
		//logx.Errorw("DeduplicationRuleService jsonx.Unmarshal err", logx.Field("err", err))
		return
	}
	if len(deduplicationConfig) <= 0 {
		//没配置限流策略
		return
	}

	//for key, value := range deduplicationConfig {
	//	exec, flag := getExec(key, l.svcCtx)
	//	//表示没匹配到对于的执行器
	//	if !flag {
	//		continue
	//	}
	//	err := exec.Deduplication(ctx, taskInfo, value)
	//	if err != nil {
	//		logx.Errorw("exec.Deduplication err", logx.Field("err", err))
	//	}
	//}

}

//func getExec(exec string, svcCtx *svc.ServiceContext) (structs.DuplicationService, bool) {
//	var duplicationExec = map[string]structs.DuplicationService{
//		deduplicationPrefix + Content:   deduplicationService.NewContentDeduplicationService(svcCtx),
//		deduplicationPrefix + Frequency: deduplicationService.NewFrequencyDeduplicationService(svcCtx),
//	}
//	v, ok := duplicationExec[exec]
//	return v, ok
//}
