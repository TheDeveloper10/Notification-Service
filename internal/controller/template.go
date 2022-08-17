package controller

import (
	"net/http"
	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"strconv"

	"github.com/gorilla/mux"
)

type TemplateV1Controller interface {
	HandleAll(http.ResponseWriter, *http.Request)
	HandleById(http.ResponseWriter, *http.Request)
	CreateTemplateFromBytes([]byte) bool
}

type basicTemplateV1Controller struct {
	repository repository.TemplateRepository
}

func NewTemplateV1Controller(repository repository.TemplateRepository) TemplateV1Controller {
	return &basicTemplateV1Controller{
		repository,
	}
}




func (btc *basicTemplateV1Controller) CreateTemplateFromBytes(bytes []byte) bool {
	reqObj := dto.CreateTemplateRequest{}
	if !util.ConvertFromJsonBytes(bytes, &reqObj) {
		return false
	}

	templateEntity := reqObj.ToEntity()
	id := btc.repository.Insert(templateEntity)
	return id != -1
}




func (btc *basicTemplateV1Controller) HandleAll(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	switch req.Method {
	case http.MethodGet:
		btc.getBulk(brw, req)
	case http.MethodPost:
		btc.create(brw, req)
	default:
		brw.Status(http.StatusMethodNotAllowed)
	}
}

func (btc *basicTemplateV1Controller) getBulk(res iface.IResponseWriter, req *http.Request) {
	// GET /templates
	// GET /templates?page=24 (size = default = 20)
	// GET /templates?size=50 (page = default = 1)

	filter := entity.TemplateFilterFromRequest(req, res)
	if filter == nil {
		return
	}

	templates := btc.repository.GetBulk(filter)
	if templates == nil {
		res.Status(http.StatusBadRequest).TextError("Failed to get anything")
	} else if len(*templates) > 0 {
		res.Status(http.StatusOK).Json(*templates)
	} else {
		res.Status(http.StatusOK)
	}
}

func (btc *basicTemplateV1Controller) create(res iface.IResponseWriter, req *http.Request) {
	reqObj := dto.CreateTemplateRequest{}
	if !util.ConvertFromJsonRequest(res, req, &reqObj) {
		return
	}

	templateEntity := reqObj.ToEntity()
	id := btc.repository.Insert(templateEntity)
	if id == -1 {
		res.Status(http.StatusBadRequest).TextError("Failed to add template to the database. Try again!")
	} else {
		placeholders := dto.GetPlaceholders(&templateEntity.Template)
		metadata := dto.TemplateMetadata{
			Id:           id,
			Placeholders: placeholders,
		}
		res.Status(http.StatusCreated).Json(metadata)
	}
}




func (btc *basicTemplateV1Controller) HandleById(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)
	templateId, _ := strconv.Atoi(mux.Vars(req)["templateId"])

	switch req.Method {
	case http.MethodGet:
		btc.getById(brw, templateId)
	case http.MethodPut:
		btc.updateById(brw, req, templateId)
	case http.MethodDelete:
		btc.deleteById(brw, templateId)
	default:
		brw.Status(http.StatusMethodNotAllowed)
	}
}

func (btc *basicTemplateV1Controller) getById(res iface.IResponseWriter, templateId int) {
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

func (btc *basicTemplateV1Controller) updateById(res iface.IResponseWriter, req *http.Request, templateId int) {
	reqObj := dto.UpdateTemplateRequest{Id: &templateId}
	if !util.ConvertFromJsonRequest(res, req, &reqObj) {
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

func (btc *basicTemplateV1Controller) deleteById(res iface.IResponseWriter, templateId int) {
	status := btc.repository.Delete(templateId)
	if status {
		res.Status(http.StatusOK).Text("Deleted successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to delete it. Try again!")
	}
}
