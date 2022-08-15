package dto

import (
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type TemplateIdRequest struct {
	iface.IRequestEntity[int]
	Id *int `json:"id"`
}

func (tir *TemplateIdRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if tir.Id == nil {
		errs.AddErrorFromString("'id' must be given")
	} else if (*tir.Id) <= 0 {
		errs.AddErrorFromString("'id' must be greater than 0")
	}

	return errs
}

func (tir *TemplateIdRequest) ToEntity() *int {
	return tir.Id
}
