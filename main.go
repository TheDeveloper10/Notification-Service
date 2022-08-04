package main

import (
	"log"
	"net/http"

	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/controller"
	"notification-service.com/packages/internal/helper"
)

func main() {
	status := helper.LoadConfig("config.yaml")
	if !status {
		log.Fatal("Failed to load configuration varaibles!")
	}

	clients.InitializeSQLClient()

	http.HandleFunc("/template", controller.NewTemplateRepository().Handle)
	http.HandleFunc("/notification", controller.GetNotification().Handle)
	
	log.Fatal(http.ListenAndServe(helper.Config.Server.Addr, nil))
}