package dto

import (
	"strings"

	"notification-service/internal/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type CreateTemplateRequest struct {
	iface.IRequestEntity[entity.TemplateEntity]
	ContactType *string `json:"contactType"`
	Template    *string `json:"template"`
	Language    *string `json:"language"`
	Type        *string `json:"type"`
}

func (ctr *CreateTemplateRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if ctr.ContactType == nil {
		errs.AddErrorFromString("'contactType' must be given")
	} else if !validateContactType(ctr.ContactType) {
		errs.AddErrorFromString("'contactType' must be one of email/sms/push")
	}

	if ctr.Language == nil || len(*ctr.Language) <= 0 {
		errs.AddErrorFromString("'language' must be given")
	} else if !validateLanguage(ctr.Language) {
		errs.AddErrorFromString("'language' must be one of " + allowedLanguages)
	}

	if ctr.Type == nil || len(*ctr.Type) <= 0 {
		errs.AddErrorFromString("'type' must be given")
	} else if len(*ctr.Type) > 8 {
		errs.AddErrorFromString("'type' must be at max 8 characters long")
	}

	if ctr.Template == nil {
		errs.AddErrorFromString("'template' must be given")
	} else if len(*ctr.Template) <= 0 || len(*ctr.Template) > 2048 {
		errs.AddErrorFromString("'template' must have a length greater than 0 and lesser than 2048")
	}

	return errs
}

func (ctr *CreateTemplateRequest) ToEntity() *entity.TemplateEntity {
	return &entity.TemplateEntity{
		ContactType: *ctr.ContactType,
		Template:    *ctr.Template,
		Language:    *ctr.Language,
		Type:        *ctr.Type,
	}
}

// TODO: Move these validations out of here
func validateContactType(contactType *string) bool {
	return *contactType == entity.ContactTypeEmail || *contactType == entity.ContactTypeSMS || *contactType == entity.ContactTypePush
}

var allowedLanguages = "BG, EN, DE, ES, DA, CS"

func validateLanguage(language *string) bool {
	if strings.Contains(*language, " ") || strings.Contains(*language, ",") {
		return false
	}
	return strings.Contains(allowedLanguages, *language)
}
