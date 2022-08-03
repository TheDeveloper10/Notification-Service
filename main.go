package main

import (
	"log"
	"net/http"
	"os"

	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/service"
)



func main() {
	status := clients.EnvLoader("config.yaml")
	if !status {
		log.Fatal("Failed to load configuration varaibles!")
	}

	http.HandleFunc("/template", service.Template)
	
	log.Fatal(http.ListenAndServe(os.Getenv("server.addr"), nil))
}