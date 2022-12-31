package process

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/taskHelper"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/rabbitmq"
)

type SendMqAction struct {
	b      broker.Broker
	logger *log.Helper
}

func NewSendMqAction(b broker.Broker,
	logger log.Logger) *SendMqAction {
	return &SendMqAction{
		b:      b,
		logger: log.NewHelper(log.With(logger, "module", "msgpusher-worker/biz/process-send-action")),
	}
}

func (p *SendMqAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel) error {
	marshal, err := json.Marshal(sendTaskModel.TaskInfo)
	if err != nil {
		return err
	}
	channel := channelType.TypeCodeEn[sendTaskModel.TaskInfo[0].SendChannel]
	msgType := messageType.TypeCodeEn[sendTaskModel.TaskInfo[0].MsgType]
	return p.b.Publish("austin.biz.routing", marshal,
		rabbitmq.WithPublishDeclareQueue(
			taskHelper.GetMqKey(channel, msgType),
			true,
			false,
			nil,
			nil,
		),
	)
}
