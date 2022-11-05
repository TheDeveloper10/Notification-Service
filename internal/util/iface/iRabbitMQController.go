package iface

type IRabbitMQController interface {
	QueueName() string
	QueueCapacity() int
	Handle([]byte) bool
}