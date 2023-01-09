package server

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-common/enums/groups"
	"austin-v2/app/msgpusher-common/enums/messageType"
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-worker/internal/conf"
	"austin-v2/app/msgpusher-worker/internal/sender"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/pkg/mq"
	"austin-v2/pkg/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/streadway/amqp"
	"github.com/tx7do/kratos-transport/broker"
	"gorm.io/gorm"
)

type RabbitMqServer struct {
	logic    *MqHandler
	mqHelper mq.IMessagingClient
}

func NewRabbitMqServer(
	c *conf.Data,
	logic *MqHandler,
	mqHelper mq.IMessagingClient,
) *RabbitMqServer {
	return &RabbitMqServer{
		mqHelper: mqHelper,
		logic:    logic,
	}
}

func (l *RabbitMqServer) Start(context.Context) error {
	fmt.Println(`RabbitMqServer Start`)
	for _, groupId := range groups.GetAllGroupIds() {
		_ = l.mqHelper.Subscribe(fmt.Sprintf("austin.biz.%s", groupId), l.logic.onMassage)
	}
	_ = l.mqHelper.Subscribe("sms.record", l.logic.smsRecord)

	return nil
}
func (l *RabbitMqServer) Stop(context.Context) error {
	fmt.Println(`RabbitMqServer Stop`)
	return nil
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
	executor *sender.TaskExecutor,
	hs *sender.HandleManager,
	taskSvc *service.TaskService,
	db *gorm.DB,
) *MqHandler {
	return &MqHandler{
		logger:   logger,
		executor: executor,
		hs:       hs,
		taskSvc:  taskSvc,
		db:       db,
	}
}

func (m *MqHandler) onMassage(delivery amqp.Delivery) {
	l := log.NewHelper(log.With(m.logger, "module", "MqHandler/onMassage"))
	var taskList []*types.TaskInfo
	_ = json.Unmarshal(delivery.Body, &taskList)
	for _, task := range taskList {
		channel := channelType.TypeCodeEn[task.SendChannel]
		msgType := messageType.TypeCodeEn[task.MsgType]
		err := m.executor.Submit(context.Background(), fmt.Sprintf("%s.%s", channel, msgType), sender.NewTask(task, m.hs, m.logger, m.taskSvc))
		if err != nil {
			l.Errorf(" on massage err: %v task_info: %s", err, task)
		}
	}
	delivery.Ack(false)
}

func (m *MqHandler) smsRecord(delivery amqp.Delivery) {
	var smsRecord []*model.SmsRecord
	_ = json.Unmarshal(delivery.Body, &smsRecord)
	m.db.Model(smsRecord).CreateInBatches(smsRecord, 500)
	delivery.Ack(false)
}
