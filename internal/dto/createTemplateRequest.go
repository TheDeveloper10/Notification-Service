package dto

import "errors"

type CreateTemplateRequest struct {
	AbstractRequest
	ContactType *string `json:"contactType"`
	Template    *string `json:"template"`
}

func (ctr *CreateTemplateRequest) Validate() (bool, error) {
	if ctr.ContactType == nil || ctr.Template == nil {
		return false, errors.New("'contactType' must be given!")
	} else if ctr.Template == nil {
		return false, errors.New("'template' muts be given!")
	} else if ctr.ContactTypeId() < 0 {
		return false, errors.New("'contactType' must be one of email/sms/push!")
	} else if len(*ctr.Template) <= 0 || len(*ctr.Template) > 2048 {
		return false, errors.New("'template' must have a length greater than 0 and lesser than 2048!")
	}
	return true, nil
}

func (ctr *CreateTemplateRequest) ContactTypeId() int8 {
	return convertStringContactTypeToInt(*ctr.ContactType)
}

func convertStringContactTypeToInt(contactType string) int8 {
	switch(contactType) {
	case "email":
		return 1
	case "sms":
		return 2
	case "push":
		return 3
	default:
		return -1
	}
}
