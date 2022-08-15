package dto

import (
	"notification-service/internal/util"
	"notification-service/internal/util/iface"

	"notification-service/internal/entity"
)

type UpdateTemplateRequest struct {
	iface.IRequestEntity[entity.TemplateEntity]
	CreateTemplateRequest
	Id *int
}

func (utr *UpdateTemplateRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()
	errs2 := utr.CreateTemplateRequest.Validate()
	errs.Merge(errs2)

	if utr.Id == nil {
		errs.AddErrorFromString("'id' must be given")
	} else if (*utr.Id) <= 0 {
		errs.AddErrorFromString("'id' must be greater than 0")
	}

	return errs
}

func (utr *UpdateTemplateRequest) ToEntity() *entity.TemplateEntity {
	return &entity.TemplateEntity{
		Id:          *utr.Id,
		ContactType: *utr.ContactType,
		Template:    *utr.Template,
		Language:    *utr.Language,
		Type:        *utr.Type,
	}
}
