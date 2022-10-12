package dto

import (
	"notification-service/internal/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type ClientIdRequest struct {
	iface.IRequestEntity[entity.ClientEntity]
	ClientID 	string   `json:"clientId"`
}

func (cir *ClientIdRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if cir.ClientID == "" {
		errs.AddErrorFromString("'clientId' must be given")
	}

	return errs
}