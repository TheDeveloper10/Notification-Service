package controller

import (
	"net/http"

	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
)

type basicTemplateController struct {
	repository repository.TemplateRepository
}

func NewTemplateController(repository repository.TemplateRepository) Controller {
	return &basicTemplateController{
		repository,
	}
}

func (btc *basicTemplateController) Handle(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	switch req.Method {
		case http.MethodPost: {
			btc.create(brw, req)
		}
		case http.MethodGet: {
			btc.get(brw, req)
		}
		case http.MethodPut: {
			btc.update(brw, req)
		}
		case http.MethodDelete: {
			btc.delete(brw, req)
		}
		default: {
			brw.Status(http.StatusMethodNotAllowed)
		}
	}
}

func (btc *basicTemplateController) create(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.CreateTemplateRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	entity := reqObj.ToEntity()
	id := btc.repository.Insert(entity)
	if id == -1 {
		res.Status(http.StatusBadRequest).Text("Failed to add template to the database. Try again!")
	} else {
		metadata := dto.TemplateMetadata{
			Id: id,
			Placeholders: entity.GetPlaceholders(),
		}
		res.Status(http.StatusCreated).Json(metadata)
	}
}

func (btc *basicTemplateController) get(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.TemplateIdRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	record, statusCode := btc.repository.Get(*reqObj.ToEntity())
	if statusCode == 1 {
		res.Status(http.StatusBadRequest).Text("Failed to get the requested template. Try again!")
		return
	} else if statusCode == 2 {
		res.Status(http.StatusNotFound).Text("Couldn't find the template you were looking for!")
		return
	} else {
		res.Status(http.StatusOK).Json(record)
	}
}

func (btc *basicTemplateController) update(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.UpdateTemplateRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	status := btc.repository.Update(reqObj.ToEntity())
	if status == 0 {
		res.Status(http.StatusOK).Text("Updated successfully!")
	} else if status == 1 {
		res.Status(http.StatusBadRequest).Text("Failed to update it. Try again!")
	} else if status == 2 { 
		res.Status(http.StatusBadRequest).Text("Failed to find template to update. Try with another one!")
	}
}

func (btc *basicTemplateController) delete(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.TemplateIdRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	status := btc.repository.Delete(*reqObj.ToEntity())
	if status {
		res.Status(http.StatusOK).Text("Deleted successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to delete it. Try again!")
	}
}