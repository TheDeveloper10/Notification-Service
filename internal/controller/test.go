package controller

import (
	"net/http"
	"notification-service/internal/util"
)

type basicTestController struct { }

func NewTestController() Controller {
	return &basicTestController{}
}

func (btc *basicTestController) Handle(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	brw.Status(200).Text("Request method: " + req.Method)
}