package dto

import (
	"notification-service/internal/util"
	"notification-service/internal/util/iface"

	"notification-service/internal/entity"
)

type UpdateTemplateRequest struct {
	iface.IRequestEntity[entity.TemplateEntity]
	CreateTemplateRequest
	TemplateIdRequest
}

func (utr *UpdateTemplateRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	errs.Merge(utr.CreateTemplateRequest.Validate())
	errs.Merge(utr.TemplateIdRequest.Validate())

	return errs
}

func (utr *UpdateTemplateRequest) ToEntity() *entity.TemplateEntity {
	return &entity.TemplateEntity{
		Id:          utr.Id,
		ContactType: utr.ContactType,
		Template:    utr.Template,
		Language:    utr.Language,
		Type:        utr.Type,
	}
}
