package service

import (
	"austin-v2/pkg/types"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"time"
)

type ShieldService struct {
	logger *log.Helper
	rds    redis.Cmdable
}

const (
	NightNoShield                = 10                  //夜间不屏蔽
	NightShield                  = 20                  //夜间屏蔽
	NightShieldButNextDaySend    = 30                  //夜间屏蔽(次日早上9点发送)
	NightShieldButNextDaySendKey = "night_shield_send" //夜间屏蔽redis key
)

//屏蔽服务
func NewShieldService(logger log.Logger, rds redis.Cmdable) *ShieldService {
	return &ShieldService{
		logger: log.NewHelper(log.With(logger, "module", "service/discard-message-service")),
		rds:    rds,
	}
}

func (s *ShieldService) Shield(ctx context.Context, taskInfo *types.TaskInfo) {
	if taskInfo.ShieldType == NightNoShield {
		return
	}

	if isNight() {
		if taskInfo.ShieldType == NightShield {
			//夜间屏蔽
			//发送到mq
			taskInfo.Receiver = []string{} //置空发送人
		}
		if taskInfo.ShieldType == NightShieldButNextDaySend {
			//夜间屏蔽,次日9点发送 扔到redis list里面 定时任务消费
			//发送到mq
			marshal, _ := json.Marshal(taskInfo)

			_, err := s.rds.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
				s.rds.LPush(ctx, NightShieldButNextDaySendKey, marshal)

				expire := int(time.Now().AddDate(0, 0, 1).Unix() - time.Now().Unix())

				s.rds.Expire(ctx, NightShieldButNextDaySendKey, time.Duration(expire)*time.Second)
				return nil
			})

			if err != nil {
				s.logger.Error(
					"msg", "夜间屏蔽(次日早上9点发送)模式 写入redis错误",
					"task_info", taskInfo,
					"err", err)
			}
			taskInfo.Receiver = []string{} //置空发送人
		}
	}
}

func isNight() bool {
	//return true
	return time.Now().Hour() < 8
}
