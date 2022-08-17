package service

import (
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"

	"notification-service/internal/controller"
	"notification-service/internal/helper"
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
	l.handleQueue("templates",     (*l.templateController).CreateTemplateFromBytes)
	l.handleQueue("notifications", (*l.notificationController).CreateNotificationFromBytes)
}

func (l *RabbitMQListener) handleQueue(queue string, target func([]byte)bool) {
	requests, err := l.channel.Consume(
		queue,
		"", false, false, false, false, nil,
	)
	if err != nil {
		l.Close()
		log.Fatal(err.Error())
	}

	go func() {
		for req := range requests {
			res := target(req.Body)
			if res {
				err := req.Ack(false)
				if err != nil {
					log.Error(err.Error())
				}
			}
		}
	}()
}
