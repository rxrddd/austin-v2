package process

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/app/msgpusher/internal/data/model"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/mqHelper"
	"austin-v2/pkg/utils/taskHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SendMqAction struct {
	mqHelper *mqHelper.MqHelper
	logger   *log.Helper
}

func NewSendMqAction(
	mqHelper *mqHelper.MqHelper,
	logger log.Logger,
) *SendMqAction {
	return &SendMqAction{
		mqHelper: mqHelper,
		logger:   log.NewHelper(log.With(logger, "module", "msgpusher-worker/biz/process-send-action")),
	}
}

func (p *SendMqAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel, _ model.MessageTemplate) error {

	channel := channelType.TypeCodeEn[sendTaskModel.TaskInfo[0].SendChannel]
	msgType := messageType.TypeCodeEn[sendTaskModel.TaskInfo[0].MsgType]

	return p.mqHelper.PublishTopic(taskHelper.GetMqKey(channel, msgType), sendTaskModel.TaskInfo)
}
