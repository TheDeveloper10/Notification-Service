package controllers

import (
	"net/http"
	"strings"

	"notification-service.com/packages/internal/service/dtos"
	"notification-service.com/packages/internal/service/repositories"
	"notification-service.com/packages/internal/service/utils"
)

func Notification(res http.ResponseWriter, req *http.Request) {
	brw := &utils.BetterResponseWriter { RW: &res }
	switch (req.Method) {
		case "POST": {
			sendNotification(brw, req)
		}
	}
}

func sendNotification(res *utils.BetterResponseWriter, req *http.Request) {
	reqObj := dtos.SendNotificationRequest{}
	if !utils.JsonMiddleware(res, req, &reqObj) {
		return
	}

	record, status := repositories.GetTemplate(&dtos.TemplateIdRequest{Id: reqObj.TemplateId})
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
	
	repositories.InsertNotification(&reqObj, &record.Template)
}