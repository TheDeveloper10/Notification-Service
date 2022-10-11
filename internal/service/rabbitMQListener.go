package service

import (
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"notification-service/internal/controller"
	"notification-service/internal/helper"
	"time"
)

type RabbitMQListener struct {
	Service

	connection *amqp.Connection
	channel    *amqp.Channel

	templateController     *controller.TemplateV1Controller
	notificationController *controller.NotificationV1Controller
}

func (l *RabbitMQListener) Init(templateController *controller.TemplateV1Controller,
								notificationController *controller.NotificationV1Controller) {
	conn, err := amqp.Dial(helper.Config.RabbitMQ.URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	l.connection = conn

	l.channel, err = conn.Channel()
	if err != nil {
		helper.HandledClose(l.connection)
		log.Fatal(err.Error())
	}

	l.templateController = templateController
	l.notificationController = notificationController
}

func (l *RabbitMQListener) Close() {
	helper.HandledClose(l.channel)
	helper.HandledClose(l.connection)
}

func (l *RabbitMQListener) Run() {
	l.handleQueue(
		helper.Config.RabbitMQ.TemplatesQueueName,
		(*l.templateController).CreateTemplateFromBytes,
		helper.Config.RabbitMQ.TemplatesQueueMax,
	)
	l.handleQueue(
		helper.Config.RabbitMQ.NotificationsQueueName,
		(*l.notificationController).CreateNotificationFromBytes,
		helper.Config.RabbitMQ.NotificationsQueueMax,
	)
	log.Info("RabbitMQ listener is ON...")
}

func (l *RabbitMQListener) handleQueue(queue string, target func([]byte), maxRunningProcesses int) {
	requests, err := l.channel.Consume(
		queue,
		"", false, false, false, false, nil,
	)
	if err != nil {
		l.Close()
		log.Fatal(err.Error())
	}

	go l.processChannel(&requests, target, maxRunningProcesses)
}

func (l *RabbitMQListener) processChannel(requests *<-chan amqp.Delivery, target func([]byte), maxRunningProcesses int) {
	runningProcesses := 0

	for request := range *requests {
		runningProcesses++

		go l.processRequest(request, target, &runningProcesses)

		// sync.WaitGroup can be used instead of doing this but it cannot
		// Wait for a single process to finish - all of them have to finish
		// in order to free up for the next requests
		for runningProcesses >= maxRunningProcesses {
			time.Sleep(2 * time.Millisecond)
		}
	}
}

func (l *RabbitMQListener) processRequest(request amqp.Delivery,
										  target func([]byte),
										  runningProcesses *int) {
	target(request.Body)
	err := request.Ack(false)
	if err != nil {
		log.Error(err.Error())
	}
	(*runningProcesses)--
}