package server

import (
	"austin-v2/app/msgpusher-worker/internal/conf"
	"austin-v2/common/enums/groups"
	"context"
	"fmt"
	"github.com/hibiken/asynq"
)

type WorkerServer struct {
	logic  *ConsumeLogic
	server *asynq.Server
}

func NewWorkerServer(
	_ *conf.Data,
	logic *ConsumeLogic,
	server *asynq.Server,
) *WorkerServer {
	return &WorkerServer{
		logic:  logic,
		server: server,
	}
}

func withRun(fn func(msg []byte) error) func(ctx context.Context, task *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		return fn(task.Payload())
	}
}

const smsRecordKey = "sms.record"
const austinKeyPrefix = "austin.biz.%s"
const nightShieldHandlerKey = "night.shield.handler"

func (l *WorkerServer) Start(context.Context) error {
	fmt.Println(`WorkerServer Start`)
	mux := asynq.NewServeMux()
	for _, groupId := range groups.GetAllGroupIds() {
		mux.HandleFunc(fmt.Sprintf(austinKeyPrefix, groupId), withRun(l.logic.onMassage))
	}
	mux.HandleFunc(smsRecordKey, withRun(l.logic.smsRecord))
	mux.HandleFunc(nightShieldHandlerKey, withRun(l.logic.nightShieldHandler))
	mux.HandleFunc("test", withRun(l.logic.test))
	return l.server.Start(mux)
}

func (l *WorkerServer) Stop(context.Context) error {
	fmt.Println(`WorkerServer Stop`)
	l.server.Stop()
	return nil
}
