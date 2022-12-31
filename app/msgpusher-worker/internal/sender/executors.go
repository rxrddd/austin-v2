package sender

import (
	"austin-v2/app/msgpusher-common/enums/groups"
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"runtime"
)

type TaskExecutor struct {
	pool map[string]*ants.Pool
}

func NewTaskExecutor() *TaskExecutor {
	//初始化所有的链接池
	groupIds := groups.GetAllGroupIds()
	pool := make(map[string]*ants.Pool)
	size := runtime.NumCPU() * 2
	for _, value := range groupIds {
		var pushWorkerPool *ants.Pool
		if wp, err := ants.NewPool(size); err != nil {
			panic(fmt.Errorf("error occurred when creating push worker: %w", err))
		} else {
			pushWorkerPool = wp
		}
		pool[value] = pushWorkerPool
	}
	return &TaskExecutor{pool: pool}
}

//把任务提交到对应的池子内
func (t *TaskExecutor) Submit(ctx context.Context, groupId string, run TaskRun) error {
	return t.pool[groupId].Submit(func() {
		run.Run(ctx)
	})
}

type TaskRun interface {
	Run(ctx context.Context)
}
