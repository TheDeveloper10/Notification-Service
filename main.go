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

	http.HandleFunc("/v1/test", controller.NewTestV1Controller().Handle)
	http.HandleFunc("/v1/templates", controller.NewTemplateV1Controller(templateRepository).Handle)
	http.HandleFunc("/v1/notifications", controller.NewNotificationV1Controller(templateRepository, notificationRepository).Handle)
	
	log.Fatal(http.ListenAndServe(helper.Config.Server.Addr, nil))
}