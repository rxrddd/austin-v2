package deduplicationService

//import (
//	limit "austin-v2/app/msgpusher-worker/internal/service/limiter"
//	"austin-v2/pkg/types"
//	"context"
//	"github.com/go-kratos/kratos/v2/log"
//	"github.com/go-redis/redis/v8"
//)
//
//type contentDeduplicationService struct {
//	logger log.Logger
//	rds    redis.Cmdable,
//}
//
//func NewContentDeduplicationService(
//	logger log.Logger,
//	rds redis.Cmdable,
//) types.DuplicationService {
//	return &contentDeduplicationService{
//		logger: logger,
//		rds:    rds,
//	}
//}
//
//func (c *contentDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param types.DeduplicationConfigItem) error {
//	return srv.NewContentDeduplicationService(c.svcCtx, limit.NewSimpleLimitService(
//		c.logger,
//		c.rds,
//	)).
//	Deduplication(ctx, taskInfo, param)
//}
