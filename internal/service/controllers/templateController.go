package controllers

import (
	"net/http"

	"notification-service.com/packages/internal/service/dtos"
	"notification-service.com/packages/internal/service/utils"
	"notification-service.com/packages/internal/service/repositories"
)

func Template(res http.ResponseWriter, req *http.Request) {
	brw := &utils.BetterResponseWriter { RW: &res }

	switch(req.Method) {
		case "POST": {
			createTemplate(brw, req)
		}
		case "GET": {
			getTemplate(brw, req)
		}
		case "PATCH": {
			updateTemplate(brw, req)
		}
		case "DELETE": {
			deleteTemplate(brw, req)
		}
	}
}

func createTemplate(res *utils.BetterResponseWriter, req *http.Request) {
	reqObj := dtos.CreateTemplateRequest{}
	if !utils.JsonMiddleware(res, req, &reqObj) {
		return
	}

	result := repositories.InsertTemplate(&reqObj)
	if result {
		// Maybe return metadata such as id
		res.Status(http.StatusOK).Text("Created successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to add template to the database. Try again!")
	}
}

func getTemplate(res *utils.BetterResponseWriter, req *http.Request) {
	reqObj := dtos.TemplateIdRequest{}
	if !utils.JsonMiddleware(res, req, &reqObj) {
		return
	}

	record, statusCode := repositories.GetTemplate(&reqObj)
	if statusCode == 1 {
		res.Status(http.StatusBadRequest).Text("Failed to get the requested template. Try again!")
		return
	} else if statusCode == 2 {
		res.Status(http.StatusNotFound).Text("Couldn't find the template you were looking for!")
		return
	} else {
		res.Status(http.StatusOK).Json(record.ToReadable())
	}
}

func updateTemplate(res *utils.BetterResponseWriter, req *http.Request) {
	reqObj := dtos.UpdateTemplateRequest{}
	if !utils.JsonMiddleware(res, req, &reqObj) {
		return
	}

	status := repositories.UpdateTemplate(&reqObj)
	if status {
		res.Status(http.StatusOK).Text("Updated successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to update it. Try again!")
	}
}

func deleteTemplate(res *utils.BetterResponseWriter, req *http.Request) {
	reqObj := dtos.TemplateIdRequest{}
	if !utils.JsonMiddleware(res, req, &reqObj) {
		return
	}

	status := repositories.DeleteTemplate(&reqObj)
	if status {
		res.Status(http.StatusOK).Text("Deleted successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to delete it. Try again!")
	}
}