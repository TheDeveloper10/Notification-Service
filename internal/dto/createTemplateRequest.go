package dto

import (
	"errors"
	"strings"

	"notification-service/internal/entity"
	"notification-service/internal/util/iface"
)

type CreateTemplateRequest struct {
	iface.IRequestEntity[entity.TemplateEntity]
	ContactType *string `json:"contactType"`
	Template    *string `json:"template"`
	Language    *string `json:"language"`
	Type        *string `json:"type"`
}

func (ctr *CreateTemplateRequest) Validate() []error {
	var errorsSlice []error

	if ctr.ContactType == nil {
		errorsSlice = append(errorsSlice, errors.New("'contactType' must be given"))
	} else if !validateContactType(ctr.ContactType) {
		errorsSlice = append(errorsSlice, errors.New("'contactType' must be one of email/sms/push"))
	}

	if ctr.Language == nil {
		errorsSlice = append(errorsSlice, errors.New("'language' must be given"))
	} else if !validateLanguage(ctr.Language) {
		errorsSlice = append(errorsSlice, errors.New("'language' must be one of " + allowedLanguages))
	}

	if ctr.Type == nil || len(*ctr.Type) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'type' must be given"))
	} else if len(*ctr.Type) > 8 {
		errorsSlice = append(errorsSlice, errors.New("'type' must be at max 8 characters long"))
	}

	if ctr.Template == nil {
		errorsSlice = append(errorsSlice, errors.New("'template' must be given"))
	} else if len(*ctr.Template) <= 0 || len(*ctr.Template) > 2048 {
		errorsSlice = append(errorsSlice, errors.New("'template' must have a length greater than 0 and lesser than 2048"))
	}

	return errorsSlice
}

func (ctr *CreateTemplateRequest) ToEntity() *entity.TemplateEntity {
	return &entity.TemplateEntity{
		ContactType: *ctr.ContactType,
		Template:	 *ctr.Template,
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