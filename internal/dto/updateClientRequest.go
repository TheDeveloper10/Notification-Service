package dto

import (
	"notification-service/internal/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type UpdateClientRequest struct {
	iface.IRequestEntity[entity.ClientEntity]
	Permissions []string `json:"permissions"`
	ClientID 	string   `json:"clientId"`
}

func (ucr *UpdateClientRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if ucr.ClientID == "" {
		errs.AddErrorFromString("'clientId' must be given")
	}

	return errs
}

func (ucr *UpdateClientRequest) ToEntity() *entity.ClientEntity {
	var permissions int64 = 0

	for i := 0; i < len(ucr.Permissions); i++ {
		permission := entity.PermissionKeyToInt(ucr.Permissions[i])
		if permission == -1 {
			continue
		}
		permissions += int64(permission)
	}

	return &entity.ClientEntity{
		Permissions: permissions,
	}
}