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
	"github.com/robfig/cron/v3"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/rabbitmq"
)

type CronTask struct {
	logger *log.Helper
	broker broker.Broker
	rds    redis.Cmdable
	c      *cron.Cron
}

func NewCronServer(
	logger log.Logger,
	broker broker.Broker,
	rds redis.Cmdable,
) *CronTask {
	return &CronTask{
		logger: log.NewHelper(log.With(logger, "module", "server/cron")),
		broker: broker,
		rds:    rds,
	}
}

func (l *CronTask) Start(context.Context) error {
	l.c = cron.New(cron.WithSeconds())
	//每早8点发送被屏蔽的消息
	l.c.AddFunc("0 0 8 * * ?", l.nightShieldHandler)
	//l.c.AddFunc("*/1 * * * * ?", func() {
	//	fmt.Println("cron task 1s" + timeHelper.CurrentTimeYMDHIS())
	//})
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

		//自己调试queue是否需要自动删除
		durableQueue := false
		autoDelete := true

		err = l.broker.Publish(taskHelper.GetMqKey(channel, msgType), marshal,
			rabbitmq.WithPublishDeclareQueue(
				taskHelper.GetMqKey(channel, msgType),
				durableQueue,
				autoDelete,
				nil,
				nil,
			),
		)
		if err != nil {
			l.logger.Errorf("nightShieldHandler Publish taskInfo err:%v taskInfo: %s ", err, string(marshal))
		}
	}
}
