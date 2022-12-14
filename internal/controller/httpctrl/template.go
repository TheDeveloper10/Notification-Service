package httpctrl

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/layer"
	dto2 "notification-service/internal/data/dto"
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"notification-service/internal/util/iface"
	"strconv"
)

type TemplateV1Controller interface {
	iface.IHTTPController
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
	reqObj := dto2.CreateTemplateRequest{}
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
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity2.PermissionReadTemplates) {
		return true
	}

	filter := entity2.TemplateFilterFromRequest(req, res)
	if filter == nil {
		return true
	}

	templates, status := btc.templateRepository.GetBulk(filter)
	if status == code.StatusSuccess {
		res.Status(http.StatusOK).JSON(*templates)
	} else {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to get anything. Try again!"))
	}
	return true
}

func (btc *basicTemplateV1Controller) create(res rem.IResponse, req rem.IRequest) bool {
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity2.PermissionCreateTemplates) {
		return true
	}

	reqObj := dto2.CreateTemplateRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	templateEntity := reqObj.ToEntity()
	id, status := btc.templateRepository.Insert(templateEntity)
	if status == code.StatusSuccess {
		metadata := dto2.TemplateMetadata{
			Id: id,
		}
		res.Status(http.StatusCreated).JSON(metadata)
	} else {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to add template to the database. Try again!"))
	}

	return true
}


func (btc *basicTemplateV1Controller) getTemplateIDFromURL(res rem.IResponse, req rem.IRequest) int {
	templateIdStr := req.GetURLParameters().Get("templateId")
	if templateIdStr == "" {
		res.
			Status(http.StatusBadRequest).
			JSON(util.ErrorListFromTextError("'templateId' must be provided via URL"))
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
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity2.PermissionReadTemplates) {
		return true
	}

	templateId := btc.getTemplateIDFromURL(res, req)
	if templateId == -1 {
		return true
	}

	record, status := btc.templateRepository.Get(templateId)
	if status == code.StatusError {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to get the requested template. Try again!"))
	} else if status == code.StatusNotFound {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Couldn't find the template you were looking for!"))
	} else {
		res.Status(http.StatusOK).JSON(record)
	}

	return true
}

func (btc *basicTemplateV1Controller) updateById(res rem.IResponse, req rem.IRequest) bool {
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity2.PermissionUpdateTemplates) {
		return true
	}

	templateId := btc.getTemplateIDFromURL(res, req)
	if templateId == -1 {
		return true
	}

	reqObj := dto2.UpdateTemplateRequest{TemplateIdRequest: dto2.TemplateIdRequest{Id: templateId}}
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
	if !layer.AccessTokenMiddleware(btc.clientRepository, res, req, entity2.PermissionDeleteTemplates) {
		return true
	}

	templateId := btc.getTemplateIDFromURL(res, req)
	if templateId == -1 {
		return true
	}

	status := btc.templateRepository.Delete(templateId)
	if status == code.StatusSuccess {
		res.Status(http.StatusOK)
	} else {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to delete it. Try again!"))
	}

	return true
}
