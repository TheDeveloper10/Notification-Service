package dto

import "errors"

type TemplateIdRequest struct {
	AbstractRequest
	Id *int `json:"id"`
}

func (ctr *TemplateIdRequest) Validate() (bool, error) {
	if ctr.Id == nil {
		return false, errors.New("'id' must be given!")
	} else if (*ctr.Id) <= 0 {
		return false, errors.New("'id' must be greater than 0")
	}
	return true, nil
}
