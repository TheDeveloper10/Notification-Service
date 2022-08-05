package dto

import "errors"

type TemplateIdRequest struct {
	AbstractRequestEntity[int]
	Id *int `json:"id"`
}

func (tir *TemplateIdRequest) Validate() []error {
	var errorsSlice []error

	if tir.Id == nil {
		errorsSlice = append(errorsSlice, errors.New("'id' must be given!"))
	} else if (*tir.Id) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'id' must be greater than 0!"))
	}
	
	return errorsSlice
}

func (tir *TemplateIdRequest) ToEntity() *int {
	return tir.Id
}