package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"notification-service/internal/clients"
	"notification-service/internal/controller"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
)

func main() {
	// Configuration
	helper.LoadConfig("./config/service_config.yaml")

	clients.InitializeSQLClient()
	clients.InitializeFCMClient("./config/adc_config.json")

	// Repositories
	templateRepository := repository.NewTemplateRepository()
	notificationRepository := repository.NewNotificationRepository()

	// Controllers
	testV1Controller := controller.NewTestV1Controller()
	templateV1Controller := controller.NewTemplateV1Controller(templateRepository)
	notificationV1Controller := controller.NewNotificationV1Controller(templateRepository, notificationRepository)

	// Routing
	r := mux.NewRouter()

	r.HandleFunc("/v1/test", testV1Controller.Handle)
	r.HandleFunc("/v1/templates", templateV1Controller.HandleAll)
	r.HandleFunc("/v1/templates/{templateId:\\d+}", templateV1Controller.HandleById)
	r.HandleFunc("/v1/notifications", notificationV1Controller.HandleAll)

	// Starting http server
	log.Info("Listening...")
	log.Fatal(http.ListenAndServe(helper.Config.Server.Addr, r))
}
