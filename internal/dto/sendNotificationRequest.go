package dto

import (
	"errors"
	"notification-service/internal/util/iface"
)

type SendNotificationRequest struct {
	iface.IRequest
	AppID        		 *string       		   `json:"appId"`
	TemplateID   		 *int          		   `json:"templateId"`
	ContactType          *string			   `json:"contactType"`

	Targets 			  []NotificationTarget `json:"targets"`

	Title        		  *string       		`json:"title"`
	UniversalPlaceholders []TemplatePlaceholder `json:"universalPlaceholders"`
}

func (snr *SendNotificationRequest) Validate() []error {
	var errorsSlice []error

	if snr.TemplateID == nil {
		errorsSlice = append(errorsSlice, errors.New("'templateId' must be given"))
	} else if (*snr.TemplateID) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'templateId' must be greater than 0"))
	}

	err := basicStringValidation("appId", snr.AppID)
	if err != nil {
		errorsSlice = append(errorsSlice, err)
	}

	err = basicStringValidation("title", snr.Title)
	if err != nil {
		errorsSlice = append(errorsSlice, err)
	}

	return errorsSlice
}

func basicStringValidation(propertyName string, property *string) error {
	if property == nil || len(*property) <= 0 {
		return errors.New("'" + propertyName + "' must be given!")
	}

	return nil
}