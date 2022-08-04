package dto

import (
	"errors"
	"strings"
)

type TemplatePlaceholder struct {
	AbstractRequest
	Key   *string `json:"key"`
	Value *string `json:"val"`
}

type SendNotificationRequest struct {
	AbstractRequest
	TemplateId   *int                  `json:"templateId"`
	UserId       *string               `json:"userId"`
	AppId        *string               `json:"appId"`
	ContactType  *string               `json:"contactType"`
	ContactInfo  *string               `json:"contactInfo"`
	Title        *string               `json:"title"`
	Placeholders []TemplatePlaceholder `json:"placeholders"`
}

func (snr *SendNotificationRequest) ContactTypeId() int8 {
	return convertStringContactTypeToInt(*snr.ContactType)
}

func (snr *SendNotificationRequest) Validate() (bool, error) {
	if snr.TemplateId == nil {
		return false, errors.New("'templateId' must be given!")
	} else if snr.UserId == nil || len(*snr.UserId) <= 0 {
		return false, errors.New("'userId' must be given!")
	} else if snr.AppId == nil || len(*snr.AppId) <= 0 {
		return false, errors.New("'appId' must be given!")
	} else if snr.ContactType == nil {
		return false, errors.New("'contactType' must be given!")
	} else if snr.ContactInfo == nil {
		return false, errors.New("'contactInfo' must be given!")
	} else if snr.Title == nil || len(*snr.Title) <= 0 {
		return false, errors.New("'title' must be given!")	
	} else if (*snr.TemplateId) <= 0 {
		return false, errors.New("'templateId' must be greater than 0!")
	} else if snr.ContactTypeId() <= 0 {
		return false, errors.New("'contactType' must be one of email/sms/push!")
	}

	for i := 0; i < len(snr.Placeholders); i++ {
		status, err := snr.Placeholders[i].Validate()
		if !status {
			return false, err
		}
	}

	return true, nil
}

func (tp *TemplatePlaceholder) Validate() (bool, error) {
	if tp.Key == nil || len(*tp.Key) <= 0 {
		return false, errors.New("'key' must be given!")
	} else if tp.Value == nil {
		return false, errors.New("'value' must be given!")
	} else if strings.HasPrefix(*tp.Key, "@{") || strings.HasSuffix(*tp.Key, "}") {
		return false, errors.New("'key' must not start with '@{' and must not end with '}'")
	}
	return true, nil
}