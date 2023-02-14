package server

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/groups"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/app/msgpusher-worker/internal/conf"
	"austin-v2/app/msgpusher-worker/internal/sender"
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/pkg/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"time"
)

type WorkerServer struct {
	logic  *Logic
	server *asynq.Server
}

func NewWorkerServer(
	_ *conf.Data,
	logic *Logic,
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
	fmt.Println(`WorkerServer Start`)
	mux := asynq.NewServeMux()
	for _, groupId := range groups.GetAllGroupIds() {
		mux.HandleFunc(fmt.Sprintf("austin.biz.%s", groupId), withRun(l.logic.onMassage))
	}
	mux.HandleFunc("sms.record", withRun(l.logic.smsRecord))
	return l.server.Start(mux)
}
func (l *WorkerServer) Stop(context.Context) error {
	fmt.Println(`WorkerServer Stop`)
	l.server.Stop()
	return nil
}

type Logic struct {
	logger   log.Logger
	executor *sender.TaskExecutor
	hs       *handler.HandleManager
	taskSvc  *service.TaskService
	suc      *biz.SmsRecordUseCase
}

func NewLogic(
	logger log.Logger,
	executor *sender.TaskExecutor,
	hs *handler.HandleManager,
	taskSvc *service.TaskService,
	suc *biz.SmsRecordUseCase,
) *Logic {
	return &Logic{
		logger:   logger,
		executor: executor,
		hs:       hs,
		taskSvc:  taskSvc,
		suc:      suc,
	}
}

func (m *Logic) onMassage(msg []byte) error {
	l := log.NewHelper(log.With(m.logger, "module", "Logic/onMassage"))
	var taskList []*types.TaskInfo
	_ = json.Unmarshal(msg, &taskList)
	for _, task := range taskList {
		channel := channelType.TypeCodeEn[task.SendChannel]
		msgType := messageType.TypeCodeEn[task.MsgType]
		task.StartConsumeAt = time.Now()
		err := m.executor.Submit(context.Background(), fmt.Sprintf("%s.%s", channel, msgType), sender.NewTask(task, m.hs, m.logger, m.taskSvc))
		if err != nil {
			l.Errorf(" on massage err: %v task_info: %s", err, task)
		}
	}
	return nil
}

func (m *Logic) smsRecord(msg []byte) error {
	var smsRecord []*model.SmsRecord
	_ = json.Unmarshal(msg, &smsRecord)
	l := log.NewHelper(log.With(m.logger, "module", "Logic/sms-record"))
	if err := m.suc.Create(context.Background(), smsRecord); err != nil {
		l.Errorf(" sms record err: %v body: %s", err, string(msg))
	}
	return nil
}
