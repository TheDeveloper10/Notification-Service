package controller

import (
	"net/http"
	"strings"

	"notification-service.com/packages/internal/dto"
	"notification-service.com/packages/internal/repository"
	"notification-service.com/packages/internal/util"
)

func Notification(res http.ResponseWriter, req *http.Request) {
	brw := &util.ResponseWriterWrapper { RW: &res }
	switch (req.Method) {
		case "POST": {
			sendNotification(brw, req)
		}
	}
}

func sendNotification(res *util.ResponseWriterWrapper, req *http.Request) {
	reqObj := dto.SendNotificationRequest{}
	if !util.JsonMiddleware(res, req, &reqObj) {
		return
	}

	record, status := repository.GetTemplate(&dto.TemplateIdRequest{Id: reqObj.TemplateId})
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
	
	repository.InsertNotification(&reqObj, &record.Template)
}