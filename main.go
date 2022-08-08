package main

import (
	"log"
	"net/http"

	"notification-service/internal/clients"
	"notification-service/internal/controller"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
)

func main() {
	helper.LoadConfig("config.yaml")
	clients.InitializeSQLClient()

	templateRepository := repository.NewTemplateRepository()
	notificationRepository := repository.NewNotificationRepository()

	http.HandleFunc("/test", controller.NewTestController().Handle)
	http.HandleFunc("/template", controller.NewTemplateController(templateRepository).Handle)
	http.HandleFunc("/notification", controller.NewNotificationController(templateRepository, notificationRepository).Handle)
	
	log.Fatal(http.ListenAndServe(helper.Config.Server.Addr, nil))
}