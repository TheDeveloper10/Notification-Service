package dto

import (
	"errors"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type SendNotificationRequest struct {
	iface.IRequest
	AppID       *string `json:"appId"`
	TemplateID  *int    `json:"templateId"`
	ContactType *string `json:"contactType"`

	Targets []NotificationTarget `json:"targets"`

	Title                 *string               `json:"title"`
	UniversalPlaceholders []TemplatePlaceholder `json:"universalPlaceholders"`
}

func (snr *SendNotificationRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if snr.TemplateID == nil {
		errs.AddErrorFromString("'templateId' must be given")
	} else if (*snr.TemplateID) <= 0 {
		errs.AddErrorFromString("'templateId' must be greater than 0")
	}

	err := basicStringValidation("appId", snr.AppID)
	if err != nil {
		errs.AddError(err)
	}

	err = basicStringValidation("title", snr.Title)
	if err != nil {
		errs.AddError(err)
	}

	err = basicStringValidation("contactType", snr.ContactType)
	if err != nil {
		errs.AddError(err)
	}

	if len(snr.Targets) <= 0 {
		errs.AddErrorFromString("at least one target must be given in 'targets'")
	}

	return errs
}

func basicStringValidation(propertyName string, property *string) error {
	if property == nil || len(*property) <= 0 {
		return errors.New("'" + propertyName + "' must be given!")
	}

	return nil
}
