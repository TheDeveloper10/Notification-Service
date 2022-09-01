package service

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"

	log "github.com/sirupsen/logrus"
)

type HTTPServer struct {
	Service
	router *rem.Router
}

func (s *HTTPServer) Init(controllers ...iface.IController) {
	if s.router != nil {
		log.Fatal("Cannot initialize HTTP server again!")
		return
	}

	s.router = rem.NewRouter()

	for _, controller := range controllers {
		controller.CreateRoutes(s.router)
	}
}

func (s *HTTPServer) Run() {
	log.Info("HTTP Server is ON...")
	log.Fatal(http.ListenAndServe(helper.Config.HTTPServer.Addr, s.router))
}