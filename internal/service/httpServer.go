package service

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"

	"github.com/sirupsen/logrus"
)

type HTTPServer struct {
	Service
	router *rem.Router
}

func (s *HTTPServer) Init(controllers ...iface.IHTTPController) {
	if s.router != nil {
		logrus.Fatal("Cannot initialize HTTP server again!")
		return
	}

	s.router = rem.NewRouter()

	for _, controller := range controllers {
		controller.CreateRoutes(s.router)
	}
}

func (s *HTTPServer) Run() {
	logrus.Info("HTTP Server is ON...")
	logrus.Fatal(http.ListenAndServe(util.Config.HTTPServer.Addr, s.router))
}