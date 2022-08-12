package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"strconv"
)

type TemplateV1Controller interface {
	HandleAll(res http.ResponseWriter, req *http.Request)
	HandleById(res http.ResponseWriter, req *http.Request)
}

type basicTemplateV1Controller struct {
	repository repository.TemplateRepository
}

func NewTemplateV1Controller(repository repository.TemplateRepository) TemplateV1Controller {
	return &basicTemplateV1Controller{
		repository,
	}
}

func (btc *basicTemplateV1Controller) HandleAll(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	switch req.Method {
		case http.MethodGet: {
			btc.getBulk(brw, req)
		}
		case http.MethodPost: {
			btc.create(brw, req)
		}
		default: {
			brw.Status(http.StatusMethodNotAllowed)
		}
	}
}

func (btc *basicTemplateV1Controller) getBulk(res util.IResponseWriter, req *http.Request) {
	// GET /notifications
	// GET /notifications?page=24 (size = default = 20)
	// GET /notifications?size=50 (page = default = 1)

	queryParams := req.URL.Query()
	filter := entity.TemplateFilter{}

	if page := queryParams.Get("page"); page != "" {
		pageInt, err := strconv.Atoi(page)
		if helper.IsError(err) || pageInt <= 0 {
			res.Status(http.StatusBadRequest).Text("'page' must be a positive integer")
			return
		}
		filter.Page = pageInt
	}

	if size := queryParams.Get("size"); size != "" {
		sizeInt, err := strconv.Atoi(size)
		if helper.IsError(err) || sizeInt <= 0 {
			res.Status(http.StatusBadRequest).Text("'size' must be a positive integer")
			return
		}
		filter.Size = sizeInt
	}

	templates := btc.repository.GetBulk(&filter)
	if templates == nil {
		res.Status(http.StatusBadRequest).Text("Failed to get anything")
	} else {
		res.Status(http.StatusOK).Json(*templates)
	}
}

func (btc *basicTemplateV1Controller) create(res util.IResponseWriter, req *http.Request) {
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




func (btc *basicTemplateV1Controller) HandleById(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)
	templateId, _ := strconv.Atoi(mux.Vars(req)["templateId"])

	switch req.Method {
		case http.MethodGet: {
			btc.getById(brw, templateId)
		}
		case http.MethodPut: {
			btc.updateById(brw, req, templateId)
		}
		case http.MethodDelete: {
			btc.deleteById(brw, templateId)
		}
		default: {
			brw.Status(http.StatusMethodNotAllowed)
		}
	}
}

func (btc *basicTemplateV1Controller) getById(res util.IResponseWriter, templateId int) {
	record, statusCode := btc.repository.Get(templateId)
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

func (btc *basicTemplateV1Controller) updateById(res util.IResponseWriter, req *http.Request, templateId int) {
	reqObj := dto.UpdateTemplateRequest{ Id: &templateId }
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

func (btc *basicTemplateV1Controller) deleteById(res util.IResponseWriter, templateId int) {
	status := btc.repository.Delete(templateId)
	if status {
		res.Status(http.StatusOK).Text("Deleted successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to delete it. Try again!")
	}
}
