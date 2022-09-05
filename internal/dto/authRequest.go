package dto

import (
	"notification-service/internal/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type AuthRequest struct {
	iface.IRequestEntity[entity.ClientCredentials]
	ClientId 	 string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

func (ar *AuthRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if ar.ClientId == "" {
		errs.AddErrorFromString("'clientId' must be given")
	} else if len(ar.ClientId) != 16 {
		errs.AddErrorFromString("'clientId' must be exactly 16 characters long")
	}

	if ar.ClientSecret == "" {
		errs.AddErrorFromString("'clientSecret' must be given")
	} else if len(ar.ClientSecret) != 128 {
		errs.AddErrorFromString("'clientSecret' must be exactly 128 characters")
	}

	return errs
}

func (ar *AuthRequest) ToEntity() *entity.ClientCredentials {
	return &entity.ClientCredentials{
		Id:     ar.ClientId,
		Secret: ar.ClientSecret,
	}
}