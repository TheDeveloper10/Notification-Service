package main

import (
	"notification-service/internal/client"
	"notification-service/internal/controller/httpctrl"
	"notification-service/internal/controller/rabbitmq"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"notification-service/internal/util"
	"sync"
)

func main() {
	// Configuration
	util.LoadConfig(util.ServiceConfigPath)

	client.InitializeSQLClient()
	client.InitializeMailClient()
	client.InitializePushClient("./config/adc_config.json")
	client.InitializeSMSClient()

	// Repositories
	clientRepository := repository.NewClientRepository()
	templateRepository := repository.NewTemplateRepository()
	notificationRepository := repository.NewNotificationRepository()

	var wg sync.WaitGroup

	// HTTP Service
	if util.Config.Service.Services.Has("http") {
		// Controllers
		testV1Controller := httpctrl.NewTestV1Controller()
		authV1Controller := httpctrl.NewAuthV1Controller(clientRepository)
		templateV1Controller := httpctrl.NewTemplateV1Controller(templateRepository, clientRepository)
		notificationV1Controller := httpctrl.NewNotificationV1Controller(templateRepository, notificationRepository, clientRepository)

		// HTTP Server
		httpServer := service.HTTPServer{}
		httpServer.Init(
			testV1Controller,
			authV1Controller,
			templateV1Controller,
			notificationV1Controller,
		)
		wg.Add(1)
		go httpServer.Run()
	}

	// RabbitMQ Service
	if util.Config.Service.Services.Has("rabbitmq") {
		// Controllers
		createNotificationController := rabbitmq.NewCreateNotificationV1Controller(templateRepository, notificationRepository)

		// RabbitMQ Listener
		rabbitMQListener := service.RabbitMQListener{}
		rabbitMQListener.Init(
			createNotificationController,
		)
		wg.Add(1)
		go rabbitMQListener.Run()
		defer rabbitMQListener.Close()
	}

	wg.Wait()
}
