package controller

import (
	"net/http"
	"notification-service/internal/util"
)

type testV1Controller struct { }

func NewTestV1Controller() Controller {
	return &testV1Controller{}
}

func (tc *testV1Controller) Handle(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	brw.Status(200).Text("Request method: " + req.Method)
}