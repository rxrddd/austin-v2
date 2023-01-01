package server

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/groups"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/app/msgpusher-worker/internal/conf"
	"austin-v2/app/msgpusher-worker/internal/sender"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/pkg/types"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	brokermq "github.com/tx7do/kratos-transport/broker/rabbitmq"
	"github.com/tx7do/kratos-transport/transport/rabbitmq"
)

// NewMqServer new a MQ server.
func NewMqServer(
	c *conf.Data,
	logger log.Logger,
	bk broker.Broker,
	executors *sender.TaskExecutor,
	hs *sender.HandleManager,
	taskSvc *service.TaskService,
) *rabbitmq.Server {

	srv := rabbitmq.NewServer(
		rabbitmq.WithAddress([]string{c.Rabbitmq.URL}),
		rabbitmq.WithCodec("json"),
	)
	logic := NewMqHandler(logger, bk, executors, hs, taskSvc)

	for _, groupId := range groups.GetAllGroupIds() {
		fmt.Println(`subscriber`, fmt.Sprintf("austin.biz.%s", groupId))
		_ = srv.RegisterSubscriber(context.Background(),
			fmt.Sprintf("austin.biz.%s", groupId),
			logic.registerMessageHandler(logic.onMassage),
			logic.messageCreator,
			broker.WithQueueName(fmt.Sprintf("austin.biz.%s", groupId)),
			//brokermq.WithDurableQueue(), //queue不自动删除
			brokermq.WithAutoDeleteQueue(), //queue自动删除
			brokermq.WithAckOnSuccess(),
		)
	}

	return srv
}

type MessageHandler func(_ context.Context, topic string, headers broker.Headers, msg *[]types.TaskInfo) error

type MqHandler struct {
	logger   log.Logger
	broker   broker.Broker
	executor *sender.TaskExecutor
	hs       *sender.HandleManager
	taskSvc  *service.TaskService
}

func NewMqHandler(
	logger log.Logger,
	broker broker.Broker,
	executor *sender.TaskExecutor,
	hs *sender.HandleManager,
	taskSvc *service.TaskService,
) *MqHandler {
	return &MqHandler{
		logger:   logger,
		broker:   broker,
		executor: executor,
		hs:       hs,
		taskSvc:  taskSvc,
	}
}
func (m *MqHandler) messageCreator() broker.Any { return &[]types.TaskInfo{} }

func (m *MqHandler) onMassage(ctx context.Context, topic string, headers broker.Headers, taskList *[]types.TaskInfo) error {
	l := log.NewHelper(log.With(m.logger, "module", "MqHandler/onMassage"))

	for _, taskInfo := range *taskList {
		task := &taskInfo
		channel := channelType.TypeCodeEn[task.SendChannel]
		msgType := messageType.TypeCodeEn[task.MsgType]
		err := m.executor.Submit(ctx, fmt.Sprintf("%s.%s", channel, msgType), sender.NewTask(task, m.hs, m.logger, m.taskSvc))
		if err != nil {
			l.Errorf(topic+" on massage err: %v task_info: %s", err, taskInfo)
		}
	}
	return nil
}

func (m *MqHandler) registerMessageHandler(fnc MessageHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *[]types.TaskInfo:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}
