package dto

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
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
	template := utr.CreateTemplateRequest.ToEntity()
	template.Id = utr.TemplateIdRequest.Id
	return template
}
