package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/util/iface"
)

func NewTestV1Controller() iface.IController {
	return &basicTestV1Controller{}
}

type basicTestV1Controller struct { }

func (btc *basicTestV1Controller) CreateRoutes(router *rem.Router) {
	router.
		NewRoute("/v1/testutils").
		MultiMethod([]string{ http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete }, btc.Handle)
}

func (btc *basicTestV1Controller) Handle(res rem.IResponse, req rem.IRequest) bool {
	res.
		Status(http.StatusOK).
		Text("Request method: " + req.GetMethod())
	return true
}