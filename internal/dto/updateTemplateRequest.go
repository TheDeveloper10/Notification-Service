package dto

import "errors"

type UpdateTemplateRequest struct {
	AbstractRequest
	Id 		 *int `json:"id"`
	Template *string `json:"template"`
}

func (utr *UpdateTemplateRequest) Validate() []error {
	var errorsSlice []error

	if utr.Id == nil {
		errorsSlice = append(errorsSlice, errors.New("'id' must be given!"))
	} else if (*utr.Id) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'id' must be greater than 0"))
	}
	
	if utr.Template == nil {
		errorsSlice = append(errorsSlice, errors.New("'template' must be given!"))
	} else if len(*utr.Template) <= 0 || len(*utr.Template) > 2048 {
		errorsSlice = append(errorsSlice, errors.New("'template' must have a length greater than 0 and lesser than 2048!"))
	}

	return errorsSlice
}
