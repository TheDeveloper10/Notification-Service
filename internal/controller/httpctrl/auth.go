package httpctrl

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/layer"
	"notification-service/internal/data/dto"
	"notification-service/internal/data/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"notification-service/internal/util/iface"
)

func NewAuthV1Controller(repository repository.IClientRepository) iface.IHTTPController {
	return &basicAuthV1Controller{
		repository,
	}
}

type basicAuthV1Controller struct {
	repository repository.IClientRepository
}

func (boac *basicAuthV1Controller) CreateRoutes(router *rem.Router) {
	// TODO: send refresher token**
	router.
		NewRoute("/v1/oauth/client").
		Post(boac.createClient)

	router.
		NewRoute("/v1/oauth/client/:clientId").
		Put(boac.updateClient).
		Delete(boac.deleteClient)

	router.
		NewRoute("/v1/oauth/token").
		Post(boac.createAuthTokens)
}

func (boac *basicAuthV1Controller) createClient(res rem.IResponse, req rem.IRequest) bool {
	if !layer.MasterTokenMiddleware(res, req) {
		return true
	}

	reqObj := dto.ClientPermissionsRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	clientEntity := reqObj.ToEntity()
	credentials, status := boac.repository.CreateClient(clientEntity)

	if status == code.StatusSuccess {
		res.Status(http.StatusCreated).JSON(credentials)
	} else if status == code.StatusError {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to create user. Try again!"))
	}

	return true
}

func (boac *basicAuthV1Controller) updateClient(res rem.IResponse, req rem.IRequest) bool {
	if !layer.MasterTokenMiddleware(res, req) {
		return true
	}

	clientID := req.GetURLParameters().Get("clientId")
	if clientID == "" {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("'clientId' must be provided via URL"))
		return true
	}

	reqObj := dto.ClientPermissionsRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	clientEntity := reqObj.ToEntity()
	status := boac.repository.UpdateClient(&clientID, clientEntity)
	if status == code.StatusSuccess {
		res.Status(http.StatusOK)
	} else if status == code.StatusNotFound {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Client not found!"))
	} else if status == code.StatusError {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Something went wrong. Try again!"))
	}

	return true
}

func (boac *basicAuthV1Controller) deleteClient(res rem.IResponse, req rem.IRequest) bool {
	if !layer.MasterTokenMiddleware(res, req) {
		return true
	}

	clientID := req.GetURLParameters().Get("clientId")
	if clientID == "" {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("'clientId' must be provided via URL"))
		return true
	}

	status := boac.repository.DeleteClient(&clientID)
	if status == 0 {
		res.Status(http.StatusOK)
	} else if status == 1 {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Client not found!"))
	} else {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Something went wrong. Try again!"))
	}

	return true
}



func (boac *basicAuthV1Controller) createAuthTokens(res rem.IResponse, req rem.IRequest) bool {
	client := layer.ClientInfoMiddleware(boac.repository, res, req)
	if client == nil {
		return true
	}

	at, status := boac.repository.GenerateToken(client, &util.Config.Service.AccessTokenSecret, util.Config.Service.AccessTokenExpiry)
	if status != code.StatusSuccess {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to generate an access token!"))
		return true
	}

	rt, status := boac.repository.GenerateToken(client, &util.Config.Service.RefresherTokenSecret, util.Config.Service.RefresherTokenExpiry)
	if status != code.StatusSuccess {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to generate a refresher token!"))
		return true
	}

	res.Status(http.StatusCreated).JSON(entity.AuthTokens{
		AccessToken: *at,
		RefresherToken: *rt,
	})

	return true
}
