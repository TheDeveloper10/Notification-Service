package dto

import (
	"errors"

	"notification-service.com/packages/internal/entity"
)

type UpdateTemplateRequest struct {
	AbstractRequestEntity[entity.TemplateEntity]
	Id 		    *int `json:"id"`
	ContactType *string `json:"contactType"`
	Template    *string `json:"template"`
}

func (utr *UpdateTemplateRequest) Validate() []error {
	var errorsSlice []error

	if utr.Id == nil {
		errorsSlice = append(errorsSlice, errors.New("'id' must be given!"))
	} else if (*utr.Id) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'id' must be greater than 0"))
	}

	if utr.ContactType == nil {
		errorsSlice = append(errorsSlice, errors.New("'contactType' must be given!"))
	} else if !validateContactType(utr.ContactType) {
		errorsSlice = append(errorsSlice, errors.New("'contactType' must be one of email/sms/push!"))
	}
	
	if utr.Template == nil {
		errorsSlice = append(errorsSlice, errors.New("'template' must be given!"))
	} else if len(*utr.Template) <= 0 || len(*utr.Template) > 2048 {
		errorsSlice = append(errorsSlice, errors.New("'template' must have a length greater than 0 and lesser than 2048!"))
	}

	return errorsSlice
}

func (utr *UpdateTemplateRequest) ToEntity() *entity.TemplateEntity {
	return &entity.TemplateEntity{
		Id: 		 *utr.Id,
		ContactType: *utr.ContactType,
		Template: 	 *utr.Template,
	}
}