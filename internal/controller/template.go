package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/layer"
	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"strconv"
)

type TemplateV1Controller interface {
	iface.IController
	CreateTemplateFromBytes(bytes []byte)
}

func NewTemplateV1Controller(
			templateRepository repository.ITemplateRepository,
			clientRepository repository.IClientRepository) TemplateV1Controller {
	return &basicTemplateV1Controller{
		templateRepository,
		clientRepository,
	}
}

type basicTemplateV1Controller struct {
	templateRepository repository.ITemplateRepository
	clientRepository   repository.IClientRepository
}

func (btc *basicTemplateV1Controller) CreateRoutes(router *rem.Router) {
	router.
		NewRoute("/v1/templates").
		Get(btc.getBulk).
		Post(btc.create)

	router.
		NewRoute("/v1/templates/:templateId").
		Get(btc.getById).
		Put(btc.updateById).
		Delete(btc.deleteById)
}

func (btc *basicTemplateV1Controller) CreateTemplateFromBytes(bytes []byte) {
	reqObj := dto.CreateTemplateRequest{}
	if !layer.JSONBytesConverterMiddleware(bytes, &reqObj) {
		return
	}

	templateEntity := reqObj.ToEntity()
	btc.templateRepository.Insert(templateEntity)
}



func (btc *basicTemplateV1Controller) getBulk(res rem.IResponse, req rem.IRequest) bool {
	// GET /templates
	// GET /templates?page=24 (size = default = 20)
	// GET /templates?size=50 (page = default = 1)
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity.PermissionReadTemplates) {
		return true
	}

	filter := entity.TemplateFilterFromRequest(req, res)
	if filter == nil {
		return true
	}

	templates := btc.templateRepository.GetBulk(filter)
	if templates == nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to get anything"))
	} else if len(*templates) > 0 {
		res.Status(http.StatusOK).JSON(*templates)
	} else {
		res.Status(http.StatusOK)
	}
	return true
}

func (btc *basicTemplateV1Controller) create(res rem.IResponse, req rem.IRequest) bool {
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity.PermissionCreateTemplates) {
		return true
	}

	reqObj := dto.CreateTemplateRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	templateEntity := reqObj.ToEntity()
	id := btc.templateRepository.Insert(templateEntity)
	if id == -1 {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to add template to the database. Try again!"))
	} else {
		metadata := dto.TemplateMetadata{
			Id: id,
		}
		res.Status(http.StatusCreated).JSON(metadata)
	}

	return true
}


func (btc *basicTemplateV1Controller) getTemplateIDFromURL(res rem.IResponse, req rem.IRequest) int {
	templateIdStr := req.GetURLParameters().Get("templateId")
	if templateIdStr == "" {
		res.
			Status(http.StatusBadRequest).
			JSON(util.ErrorListFromTextError("You must provide a 'templateId' via URL"))
		return -1
	}

	templateId, err := strconv.Atoi(templateIdStr)
	if err == nil && templateId > 0 {
		return templateId
	} else {
		res.
			Status(http.StatusBadRequest).
			JSON(util.ErrorListFromTextError("'templateId' must be a positive integer!"))
		return -1
	}
}

func (btc *basicTemplateV1Controller) getById(res rem.IResponse, req rem.IRequest) bool {
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity.PermissionReadTemplates) {
		return true
	}

	templateId := btc.getTemplateIDFromURL(res, req)
	if templateId == -1 {
		return true
	}

	record, statusCode := btc.templateRepository.Get(templateId)
	if statusCode == 1 {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to get the requested template. Try again!"))
	} else if statusCode == 2 {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Couldn't find the template you were looking for!"))
	} else {
		res.Status(http.StatusOK).JSON(record)
	}

	return true
}

func (btc *basicTemplateV1Controller) updateById(res rem.IResponse, req rem.IRequest) bool {
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity.PermissionUpdateTemplates) {
		return true
	}

	templateId := btc.getTemplateIDFromURL(res, req)
	if templateId == -1 {
		return true
	}

	reqObj := dto.UpdateTemplateRequest{TemplateIdRequest: dto.TemplateIdRequest{Id: templateId}}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	status := btc.templateRepository.Update(reqObj.ToEntity())
	if status == 0 {
		res.Status(http.StatusOK)
	} else if status == 1 {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to update it. Try again!"))
	} else if status == 2 {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to find template to update. Try with another one!"))
	}

	return true
}

func (btc *basicTemplateV1Controller) deleteById(res rem.IResponse, req rem.IRequest) bool {
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity.PermissionDeleteTemplates) {
		return true
	}

	templateId := btc.getTemplateIDFromURL(res, req)
	if templateId == -1 {
		return true
	}

	status := btc.templateRepository.Delete(templateId)
	if status {
		res.Status(http.StatusOK)
	} else {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to delete it. Try again!"))
	}

	return true
}
