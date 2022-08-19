package controller

import (
	"net/http"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type AuthV1Controller interface {
	HandleToken(http.ResponseWriter, *http.Request)
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
		boac.token(brw, req)
	default:
		brw.Status(http.StatusMethodNotAllowed)
	}
}

func (boac *basicAuthV1Controller) token(res iface.IResponseWriter, req *http.Request) {
	// TODO: move AuthRequest to header Authentication
	reqObj := dto.AuthRequest{}
	if !util.ConvertFromJsonRequest(res, req, &reqObj) {
		return
	}

	client := boac.repository.GetClient(reqObj.ToEntity())
	if client == nil {
		res.Status(http.StatusUnauthorized)
		return
	}

	accessToken := boac.repository.GenerateAccessToken(client)
	if accessToken == nil {
		res.Status(http.StatusBadRequest).TextError("Failed to generate a token!")
		return
	}

	res.Status(http.StatusOK).Json(*accessToken)
}