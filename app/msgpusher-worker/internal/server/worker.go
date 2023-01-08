package server

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/groups"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/app/msgpusher-worker/internal/conf"
	"austin-v2/app/msgpusher-worker/internal/data/model"
	"austin-v2/app/msgpusher-worker/internal/sender"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/pkg/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	brokermq "github.com/tx7do/kratos-transport/broker/rabbitmq"
	"github.com/tx7do/kratos-transport/transport/rabbitmq"
	"gorm.io/gorm"
)

// NewMqServer new a MQ server.
func NewMqServer(
	c *conf.Data,
	logic *MqHandler,
) *rabbitmq.Server {

	srv := rabbitmq.NewServer(
		rabbitmq.WithAddress([]string{c.Rabbitmq.URL}),
	)
	for _, groupId := range groups.GetAllGroupIds() {
		fmt.Println(`subscriber`, fmt.Sprintf("austin.biz.%s", groupId))
		_ = srv.RegisterSubscriber(context.Background(),
			fmt.Sprintf("austin.biz.%s", groupId),
			logic.registerRawHandler(logic.onMassage),
			nil,
			broker.WithQueueName(fmt.Sprintf("austin.biz.%s", groupId)),
			brokermq.WithAutoDeleteQueue(), //queue自动删除
			brokermq.WithAckOnSuccess(),
		)
	}

	_ = srv.RegisterSubscriber(context.Background(),
		"sms.record",
		logic.registerRawHandler(logic.smsRecord),
		nil,
		broker.WithQueueName("sms.record"),
		brokermq.WithAutoDeleteQueue(), //queue自动删除
		brokermq.WithAckOnSuccess(),
	)

	return srv
}

type RawHandler func(_ context.Context, topic string, headers broker.Headers, msg []byte) error

func (m *MqHandler) registerRawHandler(fnc RawHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		var msg []byte
		switch t := event.Message().Body.(type) {
		case []byte:
			msg = t
		case string:
			msg = []byte(t)
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}

		if err := fnc(ctx, event.Topic(), event.Message().Headers, msg); err != nil {
			return err
		}

		return nil
	}
}

type MqHandler struct {
	logger   log.Logger
	broker   broker.Broker
	executor *sender.TaskExecutor
	hs       *sender.HandleManager
	taskSvc  *service.TaskService
	db       *gorm.DB
}

func NewMqHandler(
	logger log.Logger,
	broker broker.Broker,
	executor *sender.TaskExecutor,
	hs *sender.HandleManager,
	taskSvc *service.TaskService,
	db *gorm.DB,
) *MqHandler {
	return &MqHandler{
		logger:   logger,
		broker:   broker,
		executor: executor,
		hs:       hs,
		taskSvc:  taskSvc,
		db:       db,
	}
}

func (m *MqHandler) onMassage(ctx context.Context, topic string, headers broker.Headers, msg []byte) error {
	l := log.NewHelper(log.With(m.logger, "module", "MqHandler/onMassage"))
	var taskList []*types.TaskInfo
	_ = json.Unmarshal(msg, &taskList)
	for _, task := range taskList {
		channel := channelType.TypeCodeEn[task.SendChannel]
		msgType := messageType.TypeCodeEn[task.MsgType]
		err := m.executor.Submit(ctx, fmt.Sprintf("%s.%s", channel, msgType), sender.NewTask(task, m.hs, m.logger, m.taskSvc))
		if err != nil {
			l.Errorf(topic+" on massage err: %v task_info: %s", err, task)
		}
	}
	return nil
}

func (m *MqHandler) smsRecord(ctx context.Context, topic string, headers broker.Headers, msg []byte) error {
	var smsRecord []*model.SmsRecord
	_ = json.Unmarshal(msg, &smsRecord)
	m.db.Model(smsRecord).CreateInBatches(smsRecord, 500)
	return nil
}
