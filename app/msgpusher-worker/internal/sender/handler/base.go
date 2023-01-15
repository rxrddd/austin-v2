package handler

import (
	"austin-v2/app/msgpusher-common/model/mongo_model"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/jsonHelper"
	"austin-v2/pkg/utils/stringHelper"
	"austin-v2/pkg/utils/timeHelper"
	"context"
)

type BaseHandler struct {
}

// Allow 限流方法 默认不限流
func (b BaseHandler) Allow(_ context.Context, _ *types.TaskInfo) bool {
	return true
}

func (b BaseHandler) getRecord(taskInfo *types.TaskInfo, receiver string) mongo_model.MsgRecord {
	return mongo_model.MsgRecord{
		ID:                stringHelper.NextID(),
		MessageTemplateID: taskInfo.MessageTemplateId,
		RequestID:         taskInfo.RequestId,
		CreateAt:          timeHelper.CurrentTimeYMDHIS(),
		TaskInfo:          jsonHelper.MustToString(taskInfo),
		Receiver:          receiver,
	}
}
