package process

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/idType"
	"austin-v2/app/msgpusher/internal/data/model"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/taskHelper"
	"context"
	"regexp"
)

type AfterParamCheckAction struct {
}

func NewAfterParamCheckAction() *AfterParamCheckAction {
	return &AfterParamCheckAction{}
}

func (p *AfterParamCheckAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel, _ model.MessageTemplate) error {
	// 1. 过滤掉不合法的手机号
	if sendTaskModel.TaskInfo[0].IdType == idType.Phone &&
		sendTaskModel.TaskInfo[0].SendChannel == channelType.Sms {
		var newTask []types.TaskInfo
		for _, item := range sendTaskModel.TaskInfo {
			for _, tel := range item.Receiver {
				matched, _ := regexp.Match(taskHelper.PhoneRegex, []byte(tel))
				if matched {
					newTask = append(newTask, item)
				}
			}
		}
		if len(newTask) <= 0 {
			return pb.ErrorPreParamAllFilter("没有可以发送的消息")
		}
		sendTaskModel.TaskInfo = newTask
	}

	return nil
}
