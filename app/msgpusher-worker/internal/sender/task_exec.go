package sender

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/pkg/types"
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
		logger:   log.NewHelper(log.With(logger, "module", "sender/task")),
		svc:      svc,
	}

}

func (t *Task) Run(ctx context.Context) {

	// 0. 丢弃消息 根据redis配置丢弃某个模板的所有消息
	if t.svc.DiscardMessageService.IsDiscard(ctx, t.taskInfo) {
		t.logger.Infof("消息被丢弃 task_info: %s", t.taskInfo)
		return
	}
	//// 1.屏蔽消息 1. 夜间屏蔽直接丢弃, 2.夜间屏蔽次日八点发送
	t.svc.ShieldService.Shield(ctx, t.taskInfo)
	//// 2.平台通用去重 1. N分钟相同内容去重, 2. 一天内N次相同渠道去重
	if len(t.taskInfo.Receiver) > 0 {
		t.svc.DeduplicationService.Duplication(ctx, t.taskInfo)

		if len(t.taskInfo.Receiver) <= 0 {
			t.logger.Infof("平台通用去重后没有可以发送的人了 task_info: %s", t.taskInfo)
		}
	}
	// 3. 真正发送消息

	if len(t.taskInfo.Receiver) > 0 {
		h, err := t.manager.Get(channelType.TypeCodeEn[t.taskInfo.SendChannel])
		if err != nil {
			t.logger.Errorf("handle manager route  channel: %s task_info: %s  err: %v", channelType.TypeCodeEn[t.taskInfo.SendChannel], t.taskInfo, err)
			return
		}
		for {
			if h.Allow(ctx, t.taskInfo) {
				err := h.Execute(ctx, t.taskInfo)
				if err != nil {
					t.logger.Errorf("handle execute task_info: %s err: %v ", t.taskInfo, err)
				}
				return
			}
		}
	}
}
