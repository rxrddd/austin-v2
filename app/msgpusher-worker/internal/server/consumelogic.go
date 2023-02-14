package server

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/app/msgpusher-worker/internal/sender"
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/app/msgpusher-worker/internal/service/srv"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/taskHelper"
	"austin-v2/pkg/utils/timeHelper"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"time"
)

type ConsumeLogic struct {
	loggerHelper *log.Helper
	logger       log.Logger
	executor     *sender.TaskExecutor
	hs           *handler.HandleManager
	taskSvc      *service.TaskService
	suc          *biz.SmsRecordUseCase
	cli          *asynq.Client
	rds          redis.Cmdable
}

func NewLogic(
	logger log.Logger,
	executor *sender.TaskExecutor,
	hs *handler.HandleManager,
	taskSvc *service.TaskService,
	suc *biz.SmsRecordUseCase,
	cli *asynq.Client,
	rds redis.Cmdable,
) *ConsumeLogic {
	return &ConsumeLogic{
		loggerHelper: log.NewHelper(log.With(logger, "module", "ConsumeLogic")),
		logger:       logger,
		executor:     executor,
		hs:           hs,
		taskSvc:      taskSvc,
		suc:          suc,
		cli:          cli,
		rds:          rds,
	}
}

func (m *ConsumeLogic) onMassage(msg []byte) error {
	var taskList []*types.TaskInfo
	_ = json.Unmarshal(msg, &taskList)
	for _, task := range taskList {
		channel := channelType.TypeCodeEn[task.SendChannel]
		msgType := messageType.TypeCodeEn[task.MsgType]
		task.StartConsumeAt = time.Now()
		err := m.executor.Submit(context.Background(), fmt.Sprintf("%s.%s", channel, msgType), sender.NewTask(task, m.hs, m.logger, m.taskSvc))
		if err != nil {
			m.loggerHelper.Errorf("on massage err: %v task_info: %s", err, task)
		}
	}
	return nil
}

func (m *ConsumeLogic) smsRecord(msg []byte) error {
	var smsRecord []*model.SmsRecord
	_ = json.Unmarshal(msg, &smsRecord)
	if err := m.suc.Create(context.Background(), smsRecord); err != nil {
		m.loggerHelper.Errorf(" sms record err: %v body: %s", err, string(msg))
	}
	return nil
}

func (m *ConsumeLogic) test(msg []byte) error {
	fmt.Println(timeHelper.CurrentTimeYMDHIS())
	return nil
}

//夜间屏蔽凌晨8点开始发送
func (m *ConsumeLogic) nightShieldHandler(_ []byte) error {
	ctx := context.Background()
	for {
		length, err := m.rds.LLen(ctx, srv.NightShieldButNextDaySendKey).Result()
		if err != nil {
			m.loggerHelper.Errorf("nightShieldHandler Llen err:%v", err)
			break
		}
		if length <= 0 {
			break
		}

		pop, err := m.rds.LPop(ctx, srv.NightShieldButNextDaySendKey).Result()
		if err != nil {
			m.loggerHelper.Errorf("nightShieldHandler Lpop err:%v", err)
			continue
		}
		var taskInfo types.TaskInfo
		err = json.Unmarshal([]byte(pop), &taskInfo)
		if err != nil {
			m.loggerHelper.Errorf("nightShieldHandler json.Unmarshal taskInfo err:%v", err)
			continue
		}
		channel := channelType.TypeCodeEn[taskInfo.SendChannel]
		msgType := messageType.TypeCodeEn[taskInfo.MsgType]
		marshal, _ := json.Marshal([]types.TaskInfo{taskInfo})
		_, err = m.cli.EnqueueContext(ctx, asynq.NewTask(taskHelper.GetMqKey(channel, msgType), marshal))
		if err != nil {
			m.loggerHelper.Errorf("nightShieldHandler Publish taskInfo err:%v taskInfo: %s ", err, string(marshal))
		}
	}
	return nil
}
