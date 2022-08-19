package service

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"notification-service/internal/controller"
	"notification-service/internal/helper"
)

type HTTPServer struct {
	Service
	router *mux.Router
}

func (s *HTTPServer) Init(testController *controller.TestV1Controller,
						  authController *controller.AuthV1Controller,
						  templateController *controller.TemplateV1Controller,
						  notificationController *controller.NotificationV1Controller) {
	if s.router != nil {
		log.Fatal("Cannot initialize HTTP server again!")
		return
	}

	s.router = &mux.Router{}

	s.router.HandleFunc("/v1/test", (*testController).Handle)

	s.router.HandleFunc("/v1/oauth/token", (*authController).HandleToken)

	s.router.HandleFunc("/v1/templates", (*templateController).HandleAll)
	s.router.HandleFunc("/v1/templates/{templateId:\\d+}", (*templateController).HandleById)

	s.router.HandleFunc("/v1/notifications", (*notificationController).HandleAll)
}

func (s *HTTPServer) Run() {
	log.Info("HTTP Server is ON...")
	log.Fatal(http.ListenAndServe(helper.Config.HTTPServer.Addr, s.router))
}