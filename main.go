package main

import (
	"log"
	"net/http"

	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/controller"
	"notification-service.com/packages/internal/helper"
	"notification-service.com/packages/internal/repository"
)

func main() {
	status := helper.LoadConfig("config.yaml")
	if !status {
		log.Fatal("Failed to load configuration variables!")
	}

	clients.InitializeSQLClient()

	templateRepository := repository.NewTemplateRepository()
	notificationRepository := repository.NewNotificationRepository()

	http.HandleFunc("/template", controller.NewTemplateController(templateRepository).Handle)
	http.HandleFunc("/notification", controller.NewNotificationController(templateRepository, notificationRepository).Handle)
	
	log.Fatal(http.ListenAndServe(helper.Config.Server.Addr, nil))
}