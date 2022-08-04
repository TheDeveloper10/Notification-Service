package dto

import "errors"

type TemplateIdRequest struct {
	AbstractRequest
	Id *int `json:"id"`
}

func (ctr *TemplateIdRequest) Validate() []error {
	var errorsSlice []error

	if ctr.Id == nil {
		errorsSlice = append(errorsSlice, errors.New("'id' must be given!"))
	} else if (*ctr.Id) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'id' must be greater than 0!"))
	}
	
	return errorsSlice
}
