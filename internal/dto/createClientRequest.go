package dto

import (
	"notification-service/internal/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type CreateClientRequest struct {
	iface.IRequestEntity[entity.ClientEntity]
	Permissions  *[]string `json:"permissions"`
}

func (ccr *CreateClientRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if ccr.Permissions == nil {
		errs.AddErrorFromString("Permissions must be given!")
	}

	return errs
}

func (ccr *CreateClientRequest) ToEntity() *entity.ClientEntity {
	var permissions int64 = 0
	
	for i := 0; i < len(*ccr.Permissions); i++ {
		permission := entity.PermissionKeyToInt((*ccr.Permissions)[i])
		if permission == -1 {
			continue
		}
		permissions += int64(permission)
	}

	return &entity.ClientEntity{
		Permissions: permissions,
	}
}