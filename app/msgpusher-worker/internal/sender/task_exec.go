package sender

import (
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/common/enums/channelType"
	"austin-v2/pkg/types"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Task struct {
	taskInfo *types.TaskInfo
	manager  *handler.HandleManager
	logger   *log.Helper

	svc *service.TaskService
}

func NewTask(
	taskInfo *types.TaskInfo,
	hm *handler.HandleManager,
	logger log.Logger,
	svc *service.TaskService,
) *Task {

	return &Task{
		taskInfo: taskInfo,
		manager:  hm,
		logger:   log.NewHelper(log.With(logger, "module", "sender/task-run")),
		svc:      svc,
	}

}

func (t *Task) Run(ctx context.Context) {

	// 0. 丢弃消息 根据redis配置丢弃某个模板的所有消息
	if t.svc.DiscardMessageService.IsDiscard(ctx, t.taskInfo) {
		t.logger.Infof("requestId:[%s] 消息被丢弃 task_info: %s", t.taskInfo.RequestId, t.taskInfo)
		return
	}
	//// 1.屏蔽消息 1. 夜间屏蔽直接丢弃, 2.夜间屏蔽次日八点发送
	t.svc.ShieldService.Shield(ctx, t.taskInfo)
	//// 2.平台通用去重 1. N分钟相同内容去重, 2. 一天内N次相同渠道去重
	if len(t.taskInfo.Receiver) > 0 {
		t.svc.DeduplicationService.Duplication(ctx, t.taskInfo)

		if len(t.taskInfo.Receiver) <= 0 {
			t.logger.Infof("requestId:[%s] 平台通用去重后没有可以发送的人了 task_info: %s", t.taskInfo.RequestId, t.taskInfo)
		}
	}
	// 3. 真正发送消息

	if len(t.taskInfo.Receiver) > 0 {
		h, err := t.manager.Get(channelType.TypeCodeEn[t.taskInfo.SendChannel])
		if err != nil {
			t.logger.Errorf("requestId:[%s] handle [%s] manager route  channel: %s task_info: %s  err: %v", t.taskInfo.RequestId, h.Name(), channelType.TypeCodeEn[t.taskInfo.SendChannel], t.taskInfo, err)
			return
		}
		for {
			if h.Allow(ctx, t.taskInfo) {
				err := h.Execute(ctx, t.taskInfo)
				if err != nil {
					t.logger.Errorf("requestId:[%s] handle [%s] execute task_info: %s err: %v ", t.taskInfo.RequestId, h.Name(), t.taskInfo, err)
				}
				return
			}
			t.logger.Infof("requestId:[%s] handle [%s] 触发限流 ", t.taskInfo.RequestId, h.Name())
			time.Sleep(200 * time.Millisecond)
		}
	}
}
