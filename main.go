package main

import (
	"log"

	"notification-service.com/packages/internal/clients"
)



func main() {
	status := clients.EnvLoader("config.yaml")
	if !status {
		log.Fatal("Failed to load configuration varaibles!")
	}
}