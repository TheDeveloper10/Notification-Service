package dto

import (
	"strings"

	"notification-service/internal/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type CreateTemplateRequest struct {
	iface.IRequestEntity[entity.TemplateEntity]
	Body 	 TemplateBodyRequest `json:"body"`
	Language string 			 `json:"language"`
	Type     string 			 `json:"type"`
}

func (ctr *CreateTemplateRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()
	errs.Merge(ctr.Body.Validate())

	if ctr.Language == "" {
		errs.AddErrorFromString("'language' must be given")
	} else if !validateLanguage(&ctr.Language) {
		errs.AddErrorFromString("'language' must be one of " + allowedLanguages)
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

var allowedLanguages = "BG, EN, DE, ES, DA, CS"

func validateLanguage(language *string) bool {
	if strings.Contains(*language, " ") || strings.Contains(*language, ",") {
		return false
	}
	return strings.Contains(allowedLanguages, *language)
}
