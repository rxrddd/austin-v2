package server

import (
	"austin-v2/app/mgr/internal/biz"
	"austin-v2/app/mgr/internal/conf"
	"context"
	"fmt"
	"github.com/hibiken/asynq"
)

type WorkerServer struct {
	logic  *biz.ConsumeLogic
	server *asynq.Server
}

func NewWorkerServer(
	_ *conf.Data,
	logic *biz.ConsumeLogic,
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

func (l *WorkerServer) Start(context.Context) error {
	mux := asynq.NewServeMux()
	mux.HandleFunc("test", withRun(l.logic.Test))
	return l.server.Start(mux)
}

func (l *WorkerServer) Stop(context.Context) error {
	fmt.Println(`WorkerServer Stop`)
	l.server.Stop()
	return nil
}
