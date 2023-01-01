package limit

import (
	"austin-v2/pkg/types"
)

func deduplicationAllKey(service types.IDeduplicationService, taskInfo *types.TaskInfo) []string {
	var newRows []string
	for _, receiver := range taskInfo.Receiver {
		newRows = append(newRows, deduplicationSingleKey(service, taskInfo, receiver))
	}
	return newRows
}
func deduplicationSingleKey(service types.IDeduplicationService, taskInfo *types.TaskInfo, receiver string) string {
	return service.DeduplicationSingleKey(taskInfo, receiver)
}
