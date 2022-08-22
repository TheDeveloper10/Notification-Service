package controller

import (
	"net/http"
	"notification-service/internal/controller/layer"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type AuthV1Controller interface {
	HandleToken(http.ResponseWriter, *http.Request)
	HandleClient(http.ResponseWriter, *http.Request)
}

type basicAuthV1Controller struct {
	repository repository.ClientRepository
}

func NewAuthV1Controller(repository repository.ClientRepository) AuthV1Controller {
	return &basicAuthV1Controller{
		repository,
	}
}

func (boac *basicAuthV1Controller) HandleToken(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	switch req.Method {
	case http.MethodPost:
		boac.createAccessToken(brw, req)
	default:
		brw.Status(http.StatusMethodNotAllowed)
	}
}

func (boac *basicAuthV1Controller) createAccessToken(res iface.IResponseWriter, req *http.Request) {
	client := layer.ClientInfoMiddleware(boac.repository, res, req)
	if client == nil {
		return
	}

	accessToken := boac.repository.GenerateAccessToken(client)
	if accessToken == nil {
		res.Status(http.StatusBadRequest).TextError("Failed to generate a token!")
		return
	}

	res.Status(http.StatusOK).Json(*accessToken)
}

func (boac *basicAuthV1Controller) HandleClient(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	switch req.Method {
	case http.MethodPost:
		boac.createClient(brw, req)
	default:
		brw.Status(http.StatusMethodNotAllowed)
	}
}

func (boac *basicAuthV1Controller) createClient(res iface.IResponseWriter, req *http.Request) {
	if !layer.MasterTokenMiddleware(res, req) {
		return
	}

	reqObj := dto.CreateClientRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return
	}

	clientEntity := reqObj.ToEntity()
	credentials := boac.repository.CreateClient(clientEntity)

	res.Status(http.StatusCreated).Json(credentials)
}