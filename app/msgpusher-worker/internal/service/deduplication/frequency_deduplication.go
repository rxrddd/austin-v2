package deduplicationService

//
//import (
//	"austin-go/app/austin-common/types"
//	"austin-go/app/austin-job/internal/handler/services/deduplication/limit"
//	"austin-go/app/austin-job/internal/handler/services/deduplication/srv"
//	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
//	"austin-go/app/austin-job/internal/svc"
//	"context"
//)
//
//type frequencyDeduplicationService struct {
//	svcCtx *svc.ServiceContext
//}
//
//func NewFrequencyDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
//	return &frequencyDeduplicationService{svcCtx: svcCtx}
//}
//
//func (c frequencyDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
//	return srv.NewFrequencyDeduplicationService(c.svcCtx, limit.NewSlideWindowLimitService(c.svcCtx)).
//		Deduplication(ctx, taskInfo, param)
//}
