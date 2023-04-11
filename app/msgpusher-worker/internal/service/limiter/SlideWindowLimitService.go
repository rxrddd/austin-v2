package limit

import (
	"austin-v2/app/msgpusher-worker/internal/conf"
	"austin-v2/pkg/types"
	"austin-v2/utils/timeHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"time"
)

const slideWindowLimitServiceTag = "SW_"
const slideWindowLimitServicePrefix = "slideWindowLimitServicePrefix_"

//滑动窗口去重服务
type SlideWindowLimitService struct {
	logger *log.Helper
	data   *conf.Data
	rds    redis.Cmdable
}

func NewSlideWindowLimitService(
	logger log.Logger,
	data *conf.Data,
	rds redis.Cmdable,
) *SlideWindowLimitService {
	return &SlideWindowLimitService{
		data:   data,
		rds:    rds,
		logger: log.NewHelper(log.With(logger, "module", "service/simple-limit-service")),
	}
}

func (s *SlideWindowLimitService) LimitFilter(ctx context.Context, duplication types.IDeduplicationService, taskInfo *types.TaskInfo, param types.DeduplicationConfigItem) (filterReceiver []string, err error) {
	filterReceiver = make([]string, 0)
	end := timeHelper.GetDisTodayEnd()
	for _, receiver := range taskInfo.Receiver {
		key := slideWindowLimitServiceTag + deduplicationSingleKey(duplication, taskInfo, receiver)

		num, err := s.rds.Get(ctx, key).Int()
		if err != nil {
			s.logger.Errorf("slide_window_limit_service get key:%s err:%v", key, err)
			continue
		}
		if num == 0 {
			s.rds.Expire(ctx, key, time.Duration(end)*time.Second)
		}

		if err = s.rds.Incr(ctx, key).Err(); err != nil {
			s.logger.Errorf("slide_window_limit_service incr key:%s err:%v", key, err)
			continue
		}
		//表示到了上限 直接过滤掉
		if num > param.Num {
			filterReceiver = append(filterReceiver, receiver)
		}
	}
	return filterReceiver, nil
}

func (s *SlideWindowLimitService) Name() string {
	return types.LimitSlideWindow
}
