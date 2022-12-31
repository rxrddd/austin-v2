package process

import (
	"austin-v2/app/msgpusher/internal/data/model"
	"austin-v2/pkg/types"
	"context"
)

type BusinessProcess struct {
	process []Process
}

func NewBusinessProcess(
	apc *AfterParamCheckAction,
	ass *AssembleAction,
	ppc *PreParamCheckAction,
	sma *SendMqAction,
) *BusinessProcess {
	return &BusinessProcess{
		process: []Process{
			ppc, //前置参数校验
			ass, //拼装参数
			apc, //后置参数检查
			sma, //发送到mq
		},
	}
}
func (p *BusinessProcess) Process(ctx context.Context, sendTaskModel *types.SendTaskModel, messageTemplate model.MessageTemplate) error {
	for _, pr := range p.process {
		err := pr.Process(ctx, sendTaskModel, messageTemplate)
		if err != nil {
			return err
		}
	}
	return nil
}

type Process interface {
	Process(ctx context.Context, sendTaskModel *types.SendTaskModel, messageTemplate model.MessageTemplate) error
}
