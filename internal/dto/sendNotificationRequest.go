package dto

import (
	"errors"
)

type TemplatePlaceholder struct {
	AbstractRequest
	Key   *string `json:"key"`
	Value *string `json:"val"`
}

func (tp *TemplatePlaceholder) Validate() []error {
	var errorsSlice []error
	
	if tp.Key == nil || len(*tp.Key) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'key' on each placeholder must be given!"))
	}

	if tp.Value == nil {
		errorsSlice = append(errorsSlice, errors.New("'value' on each placeholder must be given!"))
	}

	return errorsSlice
}

type SendNotificationRequest struct {
	AbstractRequest
	TemplateID   *int                  `json:"templateId"`
	UserID       *string               `json:"userId"`
	AppID        *string               `json:"appId"`
	ContactType  *string               `json:"contactType"`
	ContactInfo  *string               `json:"contactInfo"`
	Title        *string               `json:"title"`
	Placeholders []TemplatePlaceholder `json:"placeholders"`
}

func (snr *SendNotificationRequest) Validate() []error {
	var errorsSlice []error

	if snr.TemplateID == nil {
		errorsSlice = append(errorsSlice, errors.New("'templateId' must be given!"))
	} else if (*snr.TemplateID) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'templateId' must be greater than 0!"))
	}
	
	if snr.UserID == nil || len(*snr.UserID) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'userId' must be given!"))
	} 
	
	if snr.AppID == nil || len(*snr.AppID) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'appId' must be given!"))
	} 
	
	if snr.ContactType == nil {
		errorsSlice = append(errorsSlice, errors.New("'contactType' must be given!"))
	} else if !validateContactType(snr.ContactType) {
		errorsSlice = append(errorsSlice, errors.New("'contactType' must be one of email/sms/push!"))
	}

	if snr.ContactInfo == nil || len(*snr.ContactInfo) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'contactInfo' must be given!"))
	}
	if snr.Title == nil || len(*snr.Title) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'title' must be given!"))
	}
	
	for i := 0; i < len(snr.Placeholders); i++ {
		errors := snr.Placeholders[i].Validate()
		if len(errors) > 0 {
			errorsSlice = append(errorsSlice, errors...)
			return errorsSlice
		}
	}

	return errorsSlice
}
