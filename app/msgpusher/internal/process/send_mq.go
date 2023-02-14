package process

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/app/msgpusher/internal/data/model"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/jsonHelper"
	"austin-v2/pkg/utils/taskHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
)

type SendMqAction struct {
	logger *log.Helper
	cli    *asynq.Client
}

func NewSendMqAction(
	logger log.Logger,

	cli *asynq.Client,

) *SendMqAction {
	return &SendMqAction{
		logger: log.NewHelper(log.With(logger, "module", "msgpusher-worker/biz/process-send-action")),
		cli:    cli,
	}
}

func (p *SendMqAction) Process(ctx context.Context, sendTaskModel *types.SendTaskModel, _ model.MessageTemplate) error {
	channel := channelType.TypeCodeEn[sendTaskModel.TaskInfo[0].SendChannel]
	msgType := messageType.TypeCodeEn[sendTaskModel.TaskInfo[0].MsgType]
	_, err := p.cli.EnqueueContext(ctx, asynq.NewTask(taskHelper.GetMqKey(channel, msgType), jsonHelper.MustToByte(sendTaskModel.TaskInfo)))
	if err != nil {
		return pb.ErrorSystem("系统异常:%v", err)
	}
	return nil
	//return p.mqHelper.Publish(jsonHelper.MustToByte(sendTaskModel.TaskInfo), taskHelper.GetMqKey(channel, msgType))
}
