package main

import (
	"log"
	"net/http"
	"os"

	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/service/controllers"
)



func main() {
	status := clients.EnvLoader("config.yaml")
	if !status {
		log.Fatal("Failed to load configuration varaibles!")
	}

	http.HandleFunc("/template", controllers.Template)
	http.HandleFunc("/notification", controllers.Notification)
	
	log.Fatal(http.ListenAndServe(os.Getenv("server.addr"), nil))
}