package dto

import (
	"notification-service/internal/entity"
	"notification-service/internal/util/iface"
)

type ClientPermissionsRequest struct {
	iface.IRequestEntity[entity.ClientEntity]
	Permissions  []string `json:"permissions"`
}

func (ccr *ClientPermissionsRequest) Validate() iface.IErrorList {
	return nil
}

func (ccr *ClientPermissionsRequest) ToEntity() *entity.ClientEntity {
	var permissions int64 = 0
	
	for i := 0; i < len(ccr.Permissions); i++ {
		permission := entity.PermissionKeyToInt(ccr.Permissions[i])
		if permission == -1 {
			continue
		}
		permissions += int64(permission)
	}

	return &entity.ClientEntity{
		Permissions: permissions,
	}
}