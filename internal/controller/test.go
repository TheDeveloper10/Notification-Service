package controller

import (
	"net/http"
	"notification-service/internal/util"
)

type TestV1Controller interface {
	Handle (res http.ResponseWriter, req *http.Request)
}

type basicTestV1Controller struct { }

func NewTestV1Controller() TestV1Controller {
	return &basicTestV1Controller{}
}

func (tc *basicTestV1Controller) Handle(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(res)

	brw.Status(http.StatusOK).Text("Request method: " + req.Method)
}