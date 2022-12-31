package limit

import (
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/redisHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

const simpleLimitServiceTag = "SP_"

//简单去重器（目前承载着 N分钟相同内容去重）
type SimpleLimitService struct {
	logger *log.Helper
	rds    redis.Cmdable
}

func NewSimpleLimitService(
	logger log.Logger,
	rds redis.Cmdable,
) *SimpleLimitService {
	return &SimpleLimitService{
		logger: log.NewHelper(log.With(logger, "module", "service/simple-limit-service")),
		rds:    rds,
	}
}
func (s *SimpleLimitService) Name() string {
	return types.LimitSimple
}
func (s *SimpleLimitService) LimitFilter(ctx context.Context, duplication types.DeduplicationService, taskInfo *types.TaskInfo, param types.DeduplicationConfigItem) (filterReceiver []string, err error) {
	filterReceiver = make([]string, 0)
	readyPutRedisReceiver := make(map[string]string, len(taskInfo.Receiver))
	keys := each(deduplicationAllKey(duplication, taskInfo), simpleLimitServiceTag)
	inRedisValue, err := redisHelper.MGet(ctx, s.rds, keys)
	if err != nil {
		//logx.Errorw("SimpleLimitService inRedisValue MGet err", logx.Field("err", err))
		return filterReceiver, nil
	}
	for _, receiver := range taskInfo.Receiver {
		key := simpleLimitServiceTag + deduplicationSingleKey(duplication, taskInfo, receiver)
		if v, ok := inRedisValue[key]; ok {
			if cast.ToInt(v) > param.Num {
				filterReceiver = append(filterReceiver, receiver)
			} else {
				readyPutRedisReceiver[receiver] = key
			}
		}
	}
	err = s.putInRedis(ctx, readyPutRedisReceiver, inRedisValue, param.Time)
	if err != nil {
		//logx.Errorw("SimpleLimitService putInRedis err", logx.Field("err", err))
		return filterReceiver, nil
	}
	return filterReceiver, nil
}

func (s *SimpleLimitService) putInRedis(ctx context.Context, readyPutRedisReceiver, inRedisValue map[string]string, deduplicationTime int64) error {
	keyValues := make(map[string]string, len(readyPutRedisReceiver))
	for _, value := range readyPutRedisReceiver {
		if val, ok := inRedisValue[value]; ok {
			keyValues[value] = cast.ToString(cast.ToInt(val) + 1)
		} else {
			keyValues[value] = "1"
		}
	}

	if len(keyValues) > 0 {
		return redisHelper.PipelineSetEx(ctx, s.rds, keyValues, deduplicationTime)
	}
	return nil
}

func each(keys []string, tag string) []string {
	newRows := make([]string, len(keys))
	for i, key := range keys {
		newRows[i] = tag + key
	}
	return newRows
}
