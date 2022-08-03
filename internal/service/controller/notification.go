package controller

import (
	"net/http"
	"strings"

	"notification-service.com/packages/internal/dto"
	"notification-service.com/packages/internal/repository"
	"notification-service.com/packages/internal/util"
)

type notification struct {
	util.Controller
}

func GetNotification() *notification {
	return &notification{}
}

func (n *notification) Handle(res http.ResponseWriter, req *http.Request) {
	brw := util.ConvertResponseWriter(&res)

	switch (req.Method) {
		case http.MethodPost: {
			n.send(brw, req)
		}
	}
}

func (n *notification) send(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.SendNotificationRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	record, status := repository.GetTemplate().Get(&dto.TemplateIdRequest{Id: reqObj.TemplateId})
	if status == 1 {
		res.Status(http.StatusNotFound).Text("Something was wrong with the database. Try again!")
		return
	} else if status == 2 {
		res.Status(http.StatusNotFound).Text("Template was not found!")
		return
	}

	for i := 0; i < len(reqObj.Placeholders); i++ {
		placeholder := &(reqObj.Placeholders[i])
		key := "@{" + (*placeholder.Key) + "}"
		record.Template = strings.ReplaceAll(record.Template, key, *placeholder.Value)
	}
	
	repository.GetNotification().Insert(&reqObj, &record.Template)
}