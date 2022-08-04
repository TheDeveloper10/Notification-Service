package main

import (
	"log"
	"net/http"
	"os"

	"notification-service.com/packages/internal/helper"
	"notification-service.com/packages/internal/controller"
)



func main() {
	status := helper.LoadConfig("config.yaml")
	if !status {
		log.Fatal("Failed to load configuration varaibles!")
	}

	http.HandleFunc("/template", controller.NewTemplateRepository().Handle)
	http.HandleFunc("/notification", controller.GetNotification().Handle)
	
	log.Fatal(http.ListenAndServe(os.Getenv("server.addr"), nil))
}