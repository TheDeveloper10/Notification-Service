package controller

import (
	"net/http"

	"notification-service.com/packages/internal/dto"
	"notification-service.com/packages/internal/util"
	"notification-service.com/packages/internal/repository"
)

type basicTemplateController struct {
	repository repository.TemplateRepository
}

func NewTemplateController(repository repository.TemplateRepository) util.Controller {
	return &basicTemplateController{
		repository,
	}
}

func (btc *basicTemplateController) Handle(res http.ResponseWriter, req *http.Request) {
	brw := util.ConvertResponseWriter(&res)

	switch(req.Method) {
		case "POST": {
			btc.create(brw, req)
		}
		case "GET": {
			btc.get(brw, req)
		}
		case "PATCH": {
			btc.update(brw, req)
		}
		case "DELETE": {
			btc.delete(brw, req)
		}
	}
}

func (btc *basicTemplateController) create(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.CreateTemplateRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	result := btc.repository.Insert(&reqObj)
	if result {
		// Maybe return metadata such as id
		res.Status(http.StatusOK).Text("Created successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to add template to the database. Try again!")
	}
}

func (btc *basicTemplateController) get(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.TemplateIdRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	record, statusCode := btc.repository.Get(&reqObj)
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

func (btc *basicTemplateController) update(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.UpdateTemplateRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	status := btc.repository.Update(&reqObj)
	if status {
		res.Status(http.StatusOK).Text("Updated successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to update it. Try again!")
	}
}

func (btc *basicTemplateController) delete(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.TemplateIdRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	status := btc.repository.Delete(&reqObj)
	if status {
		res.Status(http.StatusOK).Text("Deleted successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to delete it. Try again!")
	}
}