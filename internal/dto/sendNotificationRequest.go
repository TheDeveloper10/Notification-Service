package dto

import (
	"errors"
	"notification-service/internal/entity"
)

type TemplatePlaceholder struct {
	AbstractRequest
	Key   *string `json:"key"`
	Value *string `json:"val"`
}

func (tp *TemplatePlaceholder) Validate() []error {
	var errorsSlice []error
	
	if tp.Key == nil || len(*tp.Key) <= 0 {
		errorsSlice = append(errorsSlice, errors.New("'key' on each placeholder must be given"))
	}

	if tp.Value == nil {
		errorsSlice = append(errorsSlice, errors.New("'val' on each placeholder must be given"))
	}

	return errorsSlice
}

type SendNotificationRequest struct {
	AbstractRequest
	TemplateID   		 *int          		   `json:"templateId"`

	ContactType          *string			   `json:"contactType"`
	Email                *string               `json:"email"`
	PhoneNumber          *string               `json:"phoneNumber"`
	FCMRegistrationToken *string       		   `json:"fcmRegistrationToken"`

	AppID        		 *string       		   `json:"appId"`
	Title        		 *string       		   `json:"title"`
	Placeholders         []TemplatePlaceholder `json:"placeholders"`
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

	switch *snr.ContactType {
		case entity.ContactTypeEmail: {
			err = basicStringValidation("email", snr.Email)
			if err != nil {
				errorsSlice = append(errorsSlice, err)
			} else if snr.PhoneNumber != nil || snr.FCMRegistrationToken != nil {
				errorsSlice = append(errorsSlice, errors.New("'email' must be given only"))
			}
		}
		case entity.ContactTypeSMS: {
			err = basicStringValidation("phoneNumber", snr.PhoneNumber)
			if err != nil {
				errorsSlice = append(errorsSlice, err)
			} else if snr.Email != nil || snr.FCMRegistrationToken != nil {
				errorsSlice = append(errorsSlice, errors.New("'phoneNumber' must be given only"))
			}
		}
		case entity.ContactTypePush: {
			err = basicStringValidation("fcmRegistrationToken", snr.FCMRegistrationToken)
			if err != nil {
				errorsSlice = append(errorsSlice, err)
			} else if snr.Email != nil || snr.PhoneNumber != nil {
				errorsSlice = append(errorsSlice, errors.New("'fcmRegistrationToken' must be given only"))
			}
		}
		default: {
			errorsSlice = append(errorsSlice, errors.New("'contactType' must be one of email/sms/push"))
		}
	}
	
	for i := 0; i < len(snr.Placeholders); i++ {
		errs := snr.Placeholders[i].Validate()
		if len(errs) > 0 {
			errorsSlice = append(errorsSlice, errs...)
			return errorsSlice
		}
	}

	return errorsSlice
}

func (snr *SendNotificationRequest) GetContactInfo() *string {
	switch *snr.ContactType {
		case entity.ContactTypeEmail:
			return snr.Email
		case entity.ContactTypeSMS:
			return snr.PhoneNumber
		case entity.ContactTypePush:
			return snr.FCMRegistrationToken
		default:
			return nil
	}
}

func basicStringValidation(propertyName string, property *string) error {
	if property == nil || len(*property) <= 0 {
		return errors.New("'" + propertyName + "' must be given!")
	}

	return nil
}