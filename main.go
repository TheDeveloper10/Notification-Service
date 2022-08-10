package main

import (
	"github.com/gorilla/mux"
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

	testV1Controller := controller.NewTestV1Controller()
	templateV1Controller := controller.NewTemplateV1Controller(templateRepository)
	notificationV1Controller := controller.NewNotificationV1Controller(templateRepository, notificationRepository)

	r := mux.NewRouter()

	r.HandleFunc("/v1/test", testV1Controller.Handle)
	r.HandleFunc("/v1/templates", templateV1Controller.Handle)
	r.HandleFunc("/v1/templates/{templateId:\\d+}", templateV1Controller.HandleSpecific)
	r.HandleFunc("/v1/notifications", notificationV1Controller.Handle)

	log.Fatal(http.ListenAndServe(helper.Config.Server.Addr, r))
}