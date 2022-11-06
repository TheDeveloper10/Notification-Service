package service

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"time"
)

type RabbitMQListener struct {
	Service

	connection *amqp.Connection
	channel    *amqp.Channel

	controllers []iface.IRabbitMQController
}

func (l *RabbitMQListener) Init(controllers ...iface.IRabbitMQController) {
	conn, err := amqp.Dial(util.Config.RabbitMQ.URL)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	l.connection = conn

	l.channel, err = conn.Channel()
	if err != nil {
		util.HandledClose(l.connection)
		logrus.Fatal(err.Error())
	}

	l.controllers = controllers
}

func (l *RabbitMQListener) Close() {
	util.HandledClose(l.channel)
	util.HandledClose(l.connection)
}

func (l *RabbitMQListener) Run() {
	for _, controller := range l.controllers {
		l.handleQueue(controller)
	}

	logrus.Info("RabbitMQ listener is ON...")
}

func (l *RabbitMQListener) handleQueue(controller iface.IRabbitMQController) {
	requests, err := l.channel.Consume(
		controller.QueueName(),
		"", false, false, false, false, nil,
	)
	if err != nil {
		l.Close()
		logrus.Fatal(err.Error())
	}

	go l.processRequests(&requests, controller)
}

func (l *RabbitMQListener) processRequests(requests *<-chan amqp.Delivery, controller iface.IRabbitMQController) {
	runningProcesses := 0

	logrus.Infof("Listening RabbitMQ queue %s...", controller.QueueName())
	for request := range *requests {
		runningProcesses++

		go l.processRequest(request, controller, &runningProcesses)

		// sync.WaitGroup can be used instead of doing this but it cannot
		// Wait for a single process to finish - all of them have to finish
		// in order to free up for the next requests
		for runningProcesses >= controller.QueueCapacity() {
			time.Sleep(4 * time.Millisecond)
		}
	}
}

func (l *RabbitMQListener) processRequest(request amqp.Delivery, controller iface.IRabbitMQController, runningProcesses *int) {
	response, acknowledge := controller.Handle(request.Body)
	(*runningProcesses)--

	if acknowledge {
		err := request.Ack(false)
		if util.ManageError(err) {
			return
		}

		if request.ReplyTo != "" && request.CorrelationId != "" {
			// publish response to
			l.publishMessage(response, request.ReplyTo, request.CorrelationId)
		}
	}
}

func (l *RabbitMQListener) publishMessage(message any, queueName string, correlationId string) {
	// TODO: make resending
	contentType := ""
	var body []byte

	if _, ok := message.(string); ok {
		body = []byte(message.(string))
		contentType = "application/text"
	} else {
		temp, err := json.Marshal(message)
		if util.ManageError(err) {
			return
		} else {
			contentType = "application/json"
			body = temp
		}
	}

	err := l.channel.PublishWithContext(
		context.Background(),
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode:  amqp.Persistent,
			Timestamp:     time.Now(),
			ContentType:   contentType,
			CorrelationId: correlationId,
			Body: 		   body,
		},
	)

	util.ManageError(err)
}