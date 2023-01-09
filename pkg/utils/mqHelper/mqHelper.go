package mqHelper

import (
	"encoding/json"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/rabbitmq"
)

type MqHelper struct {
	broker broker.Broker
}

func NewMqHelper(broker broker.Broker) *MqHelper {
	return &MqHelper{broker: broker}
}
func (m *MqHelper) PublishTopic(topic string, data interface{}) error {
	return m.broker.Publish(topic, m.convert(data),
		rabbitmq.WithPublishDeclareQueue(
			topic,
			false,
			true,
			nil,
			nil,
		),
	)
}

func (m *MqHelper) GetBroker() broker.Broker {
	return m.broker
}

func (m *MqHelper) convert(data interface{}) []byte {
	var msg []byte
	switch v := data.(type) {
	case []byte:
		msg = v
	case string:
		msg = []byte(v)
	default:
		msg, _ = json.Marshal(data)
	}
	return msg
}
