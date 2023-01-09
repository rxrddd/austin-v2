package mq

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// Defines our interface for connecting and consuming messages.
type IMessagingClient interface {
	ConnectToBroker(connectionString string) error
	PublishToExchange(msg []byte, exchangeName string, exchangeType string) error
	Publish(msg []byte, queueName string) error
	SubscribeToExchange(exchangeName string, exchangeType string, handlerFunc func(amqp.Delivery)) error
	Subscribe(queueName string, handlerFunc func(amqp.Delivery)) error
	Close()
}

// Real implementation, encapsulates a pointer to an amqp.Connection
type MessagingClient struct {
	conn *amqp.Connection
}

func NewMessagingClientURL(url string) (IMessagingClient, error) {
	messagingClient := &MessagingClient{}

	err := messagingClient.ConnectToBroker(url)
	if err != nil {
		return messagingClient, err
	}

	return messagingClient, nil
}

func NewMessagingClient(hostname, port, username, password string) (IMessagingClient, error) {
	connectionString := fmt.Sprintf("amqp://%v:%v@%v:%v", username, password, hostname, port)
	return NewMessagingClientURL(connectionString)
}

func (m *MessagingClient) ConnectToBroker(connectionString string) error {
	if connectionString == "" {
		return errors.New("Cannot initialize connection to broker, connectionString not set. Have you initialized?")
	}

	var err error
	m.conn, err = amqp.Dial(fmt.Sprintf("%s/", connectionString))
	if err != nil {
		return errors.New("Failed to connect to AMQP compatible broker at: " + connectionString)
	}

	return nil
}

func (m *MessagingClient) PublishToExchange(body []byte, exchangeName string, exchangeType string) error {
	if m.conn == nil {
		return errors.New("Tried to send message before connection was initialized. Don't do that.")
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	if err != nil {
		return err
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return err
	}
	//failOnError(err, "Failed to register an Exchange")

	queue, err := ch.QueueDeclare( // Declare a queue that will be created if not exists with some args
		"",    // our queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		queue.Name,   // name of the queue
		exchangeName, // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	err = ch.Publish( // Publishes a message onto the queue.
		exchangeName, // exchange
		exchangeName, // routing key      q.Name
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Body: body, // Our JSON body as []byte
		})
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"message": string(body),
	}).Debug("A message was sent")

	return nil
}

func (m *MessagingClient) Publish(body []byte, queueName string) error {
	if m.conn == nil {
		return errors.New("Tried to send message before connection was initialized. Don't do that.")
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare( // Declare a queue that will be created if not exists with some args
		queueName, // our queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	// Publishes a message onto the queue.
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body, // Our JSON body as []byte
		})
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"queueName": queueName,
		"message":   string(body),
	}).Debug("A message was sent to queue")

	return nil
}

func (m *MessagingClient) SubscribeToExchange(exchangeName string, exchangeType string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	if err != nil {
		//failOnError(err, "Failed to open a channel")
		return err
	}
	// defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return err
	}
	//failOnError(err, "Failed to register an Exchange")

	//log.Debugf("declared Exchange, declaring Queue (%s)", "")
	queue, err := ch.QueueDeclare(
		"",    // name of the queue
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	//failOnError(err, "Failed to register an Queue")

	log.Debugf("declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')",
		queue.Messages, queue.Consumers, exchangeName)

	err = ch.QueueBind(
		queue.Name,   // name of the queue
		exchangeName, // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}
	//failOnError(err, "Failed to register a consumer")

	go consumeLoop(msgs, handlerFunc)
	return nil
}

func (m *MessagingClient) Subscribe(queueName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	if err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	go consumeLoop(msgs, handlerFunc)

	log.Debugf("Succesfully subscribed to queue %v", queueName)
	return nil
}

func (m *MessagingClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		// Invoke the handlerFunc func we passed as parameter.
		handlerFunc(d)
	}
}
