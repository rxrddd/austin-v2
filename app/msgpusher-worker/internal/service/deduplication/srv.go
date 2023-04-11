package deduplication

import (
	"austin-v2/pkg/types"
	"austin-v2/utils/arrayHelper"
	"context"
)

type deduplicationService struct {
}

func (c deduplicationService) Deduplication(ctx context.Context,
	limit types.ILimitService,
	service types.IDeduplicationService,
	taskInfo *types.TaskInfo,
	param types.DeduplicationConfigItem) error {

	var newRows []string
	filter, err := limit.LimitFilter(ctx, service, taskInfo, param)
	if err != nil {
		return err
	}
	for _, s := range taskInfo.Receiver {
		if !arrayHelper.ArrayStringIn(filter, s) {
			newRows = append(newRows, s)
		}
	}
	taskInfo.Receiver = newRows
	return nil
}

const Content = "10"   //N分钟相同内容去重
const Frequency = "20" //一天内N次相同渠道去重
const deduplicationPrefix = "deduplication_"
