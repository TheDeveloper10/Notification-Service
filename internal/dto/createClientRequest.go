package dto

import (
	"notification-service/internal/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type CreateClientRequest struct {
	iface.IRequestEntity[entity.ClientEntity]
	Permissions  *int    `json:"permissions"`
}

func (ccr *CreateClientRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if ccr.Permissions == nil {
		errs.AddErrorFromString("Permissions must be given!")
	}

	return errs
}

func (ccr *CreateClientRequest) ToEntity() *entity.ClientEntity {
	return &entity.ClientEntity{
		Permissions: *ccr.Permissions,
	}
}