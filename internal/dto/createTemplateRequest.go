package dto

import "errors"

type CreateTemplateRequest struct {
	AbstractRequest
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

func validateContactType(contactType *string) bool {
	return *contactType == "email" || *contactType == "sms" || *contactType == "push"
}