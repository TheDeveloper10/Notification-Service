package dto

import (
	"errors"

	"notification-service/internal/entity"
)

type UpdateTemplateRequest struct {
	AbstractRequestEntity[entity.TemplateEntity]
	CreateTemplateRequest
	Id 		    *int `json:"id"`
}

func (utr *UpdateTemplateRequest) Validate() []error {
	errorsSlice := utr.CreateTemplateRequest.Validate()

	if utr.Id == nil {
		errorsSlice = append(errorsSlice, errors.New("'id' must be given"))
	} else if (*utr.Id) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'id' must be greater than 0"))
	}

	return errorsSlice
}

func (utr *UpdateTemplateRequest) ToEntity() *entity.TemplateEntity {
	return &entity.TemplateEntity{
		Id: 		 *utr.Id,
		ContactType: *utr.ContactType,
		Template: 	 *utr.Template,
		Language:    *utr.Language,
		Type:        *utr.Type,
	}
}