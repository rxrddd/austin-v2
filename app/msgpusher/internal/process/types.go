package process

import (
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
			apc,
			ass,
			ppc,
			sma,
		},
	}
}
func (p *BusinessProcess) Process(ctx context.Context, sendTaskModel *types.SendTaskModel) error {
	for _, pr := range p.process {
		err := pr.Process(ctx, sendTaskModel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *BusinessProcess) AddProcess(pr ...Process) error {
	if len(pr) > 0 {
		p.process = append(p.process, pr...)
	}
	return nil
}

type Process interface {
	Process(ctx context.Context, sendTaskModel *types.SendTaskModel) error
}
