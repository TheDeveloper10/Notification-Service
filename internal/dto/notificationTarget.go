package dto

import (
	"errors"
	"notification-service/internal/entity"
)

type NotificationTarget struct {
	Email                *string               `json:"email"`
	PhoneNumber          *string               `json:"phoneNumber"`
	FCMRegistrationToken *string       		   `json:"fcmRegistrationToken"`
	Placeholders         []TemplatePlaceholder `json:"placeholders"`
}

func (nt *NotificationTarget) Validate(contactType *string) error {
	switch *contactType {
		case entity.ContactTypeEmail: {
			if nt.Email == nil {
				return errors.New("'email' must be given")
			}
			break
		}
		case entity.ContactTypeSMS: {
			if nt.PhoneNumber == nil {
				return errors.New("'phoneNumber' must be given")
			}
			break
		}
		case entity.ContactTypePush: {
			if nt.FCMRegistrationToken == nil {
				return errors.New("'fcmRegistrationToken' must be given")
			}
			break
		}
	}

	return nil
}

func (nt *NotificationTarget) GetContactInfo() *string {
	if nt.Email != nil {
		return nt.Email
	} else if nt.PhoneNumber != nil {
		return nt.PhoneNumber
	} else if nt.FCMRegistrationToken != nil {
		return nt.FCMRegistrationToken
	}
	return nil
}

func (nt *NotificationTarget) GetContactType() *string {
	if nt.Email != nil {
		return &entity.ContactTypeEmail
	} else if nt.PhoneNumber != nil {
		return &entity.ContactTypeSMS
	} else if nt.FCMRegistrationToken != nil {
		return &entity.ContactTypePush
	}
	return nil
}