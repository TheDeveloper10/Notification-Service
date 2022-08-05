package dto

import (
	"errors"

	"notification-service.com/packages/internal/entity"
)

type CreateTemplateRequest struct {
	AbstractRequestEntity[entity.TemplateEntity]
	ContactType *string `json:"contactType"`
	Template    *string `json:"template"`
}

func (ctr *CreateTemplateRequest) Validate() []error {
	var errorsSlice []error

	if ctr.ContactType == nil {
		errorsSlice = append(errorsSlice, errors.New("'contactType' must be given!"))
	} else if !validateContactType(ctr.ContactType) {
		errorsSlice = append(errorsSlice, errors.New("'contactType' must be one of email/sms/push!"))
	}
	
	if ctr.Template == nil {
		errorsSlice = append(errorsSlice, errors.New("'template' muts be given!"))
	} else if len(*ctr.Template) <= 0 || len(*ctr.Template) > 2048 {
		errorsSlice = append(errorsSlice, errors.New("'template' must have a length greater than 0 and lesser than 2048!"))
	}

	return errorsSlice
}

func (ctr *CreateTemplateRequest) ToEntity() *entity.TemplateEntity {
	return &entity.TemplateEntity{
		ContactType: *ctr.ContactType,
		Template:	 *ctr.Template,
	}
}

func validateContactType(contactType *string) bool {
	return *contactType == "email" || *contactType == "sms" || *contactType == "push"
}