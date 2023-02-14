package handler

import (
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/jsonHelper"
	"austin-v2/pkg/utils/stringHelper"
	"austin-v2/pkg/utils/timeHelper"
	"context"
	"time"
)

type BaseHandler struct {
}

// Allow 限流方法 默认不限流
func (b BaseHandler) Allow(_ context.Context, _ *types.TaskInfo) bool {
	return true
}

func (b BaseHandler) getRecord(taskInfo *types.TaskInfo, receiver string) *model.MsgRecord {
	return &model.MsgRecord{
		ID:                stringHelper.NextID(),
		MessageTemplateID: taskInfo.MessageTemplateId,
		RequestID:         taskInfo.RequestId,
		CreateAt:          time.Now(),
		TaskInfo:          jsonHelper.MustToString(taskInfo),
		Receiver:          receiver,
		StartConsumeAt:    taskInfo.StartConsumeAt.Format(timeHelper.DateDefaultLayout),
		SendAt:            taskInfo.SendAt.Format(timeHelper.DateDefaultLayout),
		EndConsumeAt:      timeHelper.CurrentTimeYMDHIS(),
		ConsumeSinceTime:  time.Now().Sub(taskInfo.StartConsumeAt).String(),
		SendSinceTime:     time.Now().Sub(taskInfo.SendAt).String(),
	}
}
