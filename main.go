package main

import (
	"notification-service/internal/client"
	"notification-service/internal/controller"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"sync"
)

func main() {
	// Configuration
	helper.LoadConfig(helper.ServiceConfigPath)

	client.InitializeSQLClient()
	client.InitializeMailClient()
	client.InitializePushClient("./config/adc_config.json")
	client.InitializeSMSClient()

	// Repositories
	clientRepository := repository.NewClientRepository()
	templateRepository := repository.NewTemplateRepository()
	notificationRepository := repository.NewNotificationRepository()

	var wg sync.WaitGroup

	// Controllers
	testV1Controller := controller.NewTestV1Controller()
	authV1Controller := controller.NewAuthV1Controller(clientRepository)
	templateV1Controller := controller.NewTemplateV1Controller(templateRepository, clientRepository)
	notificationV1Controller := controller.NewNotificationV1Controller(templateRepository, notificationRepository, clientRepository)

	// HTTP Server
	if helper.Config.Service.Services.Has("http") {
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

	// RabbitMQ Listener
	if helper.Config.Service.Services.Has("rabbitmq") {
		rabbitMQListener := service.RabbitMQListener{}
		rabbitMQListener.Init(&templateV1Controller, &notificationV1Controller)
		wg.Add(1)
		go rabbitMQListener.Run()
		defer rabbitMQListener.Close()
	}

	wg.Wait()
}
