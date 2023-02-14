package server

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/app/msgpusher-worker/internal/service/srv"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/taskHelper"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/robfig/cron/v3"
)

type CronTask struct {
	logger *log.Helper
	cli    *asynq.Client
	rds    redis.Cmdable
	c      *cron.Cron
}

func NewCronServer(
	logger log.Logger,
	cli *asynq.Client,
	rds redis.Cmdable,
) *CronTask {
	return &CronTask{
		logger: log.NewHelper(log.With(logger, "module", "server/cron")),
		cli:    cli,
		rds:    rds,
	}
}

func (l *CronTask) Start(context.Context) error {
	l.c = cron.New(cron.WithSeconds())
	//每早8点发送被屏蔽的消息
	l.c.AddFunc("0 0 8 * * ?", l.nightShieldHandler)
	l.c.Start()
	l.logger.Info("start the cron task")
	return nil
}
func (l *CronTask) Stop(context.Context) error {
	l.c.Stop()
	l.logger.Info("close the cron task")
	return nil
}

//夜间屏蔽凌晨8点开始发送
func (l *CronTask) nightShieldHandler() {
	ctx := context.Background()
	for {
		length, err := l.rds.LLen(ctx, srv.NightShieldButNextDaySendKey).Result()
		if err != nil {
			l.logger.Errorf("nightShieldHandler Llen err:%v", err)
			break
		}
		if length <= 0 {
			break
		}

		pop, err := l.rds.LPop(ctx, srv.NightShieldButNextDaySendKey).Result()
		if err != nil {
			l.logger.Errorf("nightShieldHandler Lpop err:%v", err)
			continue
		}
		var taskInfo types.TaskInfo
		err = json.Unmarshal([]byte(pop), &taskInfo)
		if err != nil {
			l.logger.Errorf("nightShieldHandler json.Unmarshal taskInfo err:%v", err)
			continue
		}
		channel := channelType.TypeCodeEn[taskInfo.SendChannel]
		msgType := messageType.TypeCodeEn[taskInfo.MsgType]
		marshal, _ := json.Marshal([]types.TaskInfo{taskInfo})
		_, err = l.cli.EnqueueContext(ctx, asynq.NewTask(taskHelper.GetMqKey(channel, msgType), marshal))
		if err != nil {
			l.logger.Errorf("nightShieldHandler Publish taskInfo err:%v taskInfo: %s ", err, string(marshal))
		}
	}
}
