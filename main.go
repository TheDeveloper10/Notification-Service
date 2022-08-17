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
	helper.LoadConfig("./config/service_config.yaml")

	client.InitializeSQLClient()
	client.InitializeMailClient()
	client.InitializePushClient("./config/adc_config.json")

	// Repositories
	templateRepository := repository.NewTemplateRepository()
	notificationRepository := repository.NewNotificationRepository()

	var wg sync.WaitGroup

	// Controllers
	testV1Controller := controller.NewTestV1Controller()
	templateV1Controller := controller.NewTemplateV1Controller(templateRepository)
	notificationV1Controller := controller.NewNotificationV1Controller(templateRepository, notificationRepository)

	// HTTP Server
	if helper.Config.Service.UseHTTP == "yes" {
		httpServer := service.HTTPServer{}
		httpServer.Init(&testV1Controller, &templateV1Controller, &notificationV1Controller)
		wg.Add(1)
		go httpServer.Run()
	}

	// RabbitMQ Listener
	if helper.Config.Service.UseRabbitMQ == "yes" {
		rabbitMQListener := service.RabbitMQListener{}
		rabbitMQListener.Init(&templateV1Controller, &notificationV1Controller)
		wg.Add(1)
		go rabbitMQListener.Run()
		defer rabbitMQListener.Close()
	}

	wg.Wait()
}
