package controller

import (
	"net/http"

	"notification-service.com/packages/internal/dto"
	"notification-service.com/packages/internal/util"
	"notification-service.com/packages/internal/repository"
)

func Template(res http.ResponseWriter, req *http.Request) {
	brw := &util.ResponseWriterWrapper { RW: &res }

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

func createTemplate(res *util.ResponseWriterWrapper, req *http.Request) {
	reqObj := dto.CreateTemplateRequest{}
	if !util.JsonMiddleware(res, req, &reqObj) {
		return
	}

	result := repository.InsertTemplate(&reqObj)
	if result {
		// Maybe return metadata such as id
		res.Status(http.StatusOK).Text("Created successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to add template to the database. Try again!")
	}
}

func getTemplate(res *util.ResponseWriterWrapper, req *http.Request) {
	reqObj := dto.TemplateIdRequest{}
	if !util.JsonMiddleware(res, req, &reqObj) {
		return
	}

	record, statusCode := repository.GetTemplate(&reqObj)
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

func updateTemplate(res *util.ResponseWriterWrapper, req *http.Request) {
	reqObj := dto.UpdateTemplateRequest{}
	if !util.JsonMiddleware(res, req, &reqObj) {
		return
	}

	status := repository.UpdateTemplate(&reqObj)
	if status {
		res.Status(http.StatusOK).Text("Updated successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to update it. Try again!")
	}
}

func deleteTemplate(res *util.ResponseWriterWrapper, req *http.Request) {
	reqObj := dto.TemplateIdRequest{}
	if !util.JsonMiddleware(res, req, &reqObj) {
		return
	}

	status := repository.DeleteTemplate(&reqObj)
	if status {
		res.Status(http.StatusOK).Text("Deleted successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to delete it. Try again!")
	}
}