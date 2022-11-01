package dto

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"
)

type CreateTemplateRequest struct {
	iface.IRequestEntity[entity.TemplateEntity]
	Body     TemplateBodyRequest `json:"body"`
	Language string              `json:"language"`
	Type     string 			 `json:"type"`
}

func (ctr *CreateTemplateRequest) Validate() iface.IErrorList {
	errs := ctr.Body.Validate()

	if ctr.Language == "" {
		errs.AddErrorFromString("'language' must be given")
	} else if !helper.Config.Service.AllowedLanguages.Has(ctr.Language) {
		errs.AddErrorFromString("'language' must be one of " + helper.Config.Service.AllowedLanguages.Join(", "))
	}

	if ctr.Type == "" {
		errs.AddErrorFromString("'type' must be given")
	} else if len(ctr.Type) > 8 {
		errs.AddErrorFromString("'type' must be at max 8 characters long")
	}

	return errs
}

func (ctr *CreateTemplateRequest) ToEntity() *entity.TemplateEntity {
	return &entity.TemplateEntity{
		Body: 	  ctr.Body.ToEntity(),
		Language: ctr.Language,
		Type:     ctr.Type,
	}
}
