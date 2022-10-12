package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/layer"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

func NewAuthV1Controller(repository repository.IClientRepository) iface.IController {
	return &basicAuthV1Controller{
		repository,
	}
}

type basicAuthV1Controller struct {
	repository repository.IClientRepository
}

func (boac *basicAuthV1Controller) CreateRoutes(router *rem.Router) {
	router.
		NewRoute("/v1/oauth/client").
		Post(boac.createClient).
		Put(boac.updateClient).
		Delete(boac.deleteClient)

	router.
		NewRoute("/v1/oauth/token").
		Post(boac.createAccessToken)
}

func (boac *basicAuthV1Controller) createClient(res rem.IResponse, req rem.IRequest) bool {
	if !layer.MasterTokenMiddleware(res, req) {
		return true
	}

	reqObj := dto.CreateClientRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	clientEntity := reqObj.ToEntity()
	credentials := boac.repository.CreateClient(clientEntity)

	if credentials == nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to create user!"))
	} else {
		res.Status(http.StatusCreated).JSON(credentials)
	}
	return true
}

func (boac *basicAuthV1Controller) updateClient(res rem.IResponse, req rem.IRequest) bool {
	if !layer.MasterTokenMiddleware(res, req) {
		return true
	}

	reqObj := dto.UpdateClientRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	clientEntity := reqObj.ToEntity()
	status := boac.repository.UpdateClient(&reqObj.ClientID, clientEntity)
	if status == 0 {
		res.Status(http.StatusOK)
	} else if status == 1 {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Client not found!"))
	} else {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Something went wrong. Try again!"))
	}

	return true
}

func (boac *basicAuthV1Controller) deleteClient(res rem.IResponse, req rem.IRequest) bool {
	if !layer.MasterTokenMiddleware(res, req) {
		return true
	}

	reqObj := dto.ClientIdRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	status := boac.repository.DeleteClient(&reqObj.ClientID)
	if status == 0 {
		res.Status(http.StatusOK)
	} else if status == 1 {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Client not found!"))
	} else {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Something went wrong. Try again!"))
	}

	return true
}

func (boac *basicAuthV1Controller) createAccessToken(res rem.IResponse, req rem.IRequest) bool {
	client := layer.ClientInfoMiddleware(boac.repository, res, req)
	if client == nil {
		return true
	}

	accessToken := boac.repository.GenerateAccessToken(client)
	if accessToken == nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to generate a token!"))
		return true
	}

	res.Status(http.StatusOK).JSON(*accessToken)
	return true
}
