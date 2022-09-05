package dto

import (
	"errors"
	"net/mail"
	"notification-service/internal/entity"
	"regexp"
)

type NotificationTarget struct {
	Email                string               `json:"email"`
	PhoneNumber          string               `json:"phoneNumber"`
	FCMRegistrationToken string       		   `json:"fcmRegistrationToken"`
	Placeholders         []TemplatePlaceholder `json:"placeholders"`
}

func (nt *NotificationTarget) Validate(contactType *string) error {
	if contactType == nil {
		return errors.New("'contactType' must be given")
	}
	switch *contactType {
		case entity.ContactTypeEmail: {
			if nt.Email == "" {
				return errors.New("'email' must be given")
			} else if _, err := mail.ParseAddress(nt.Email); err != nil {
				return errors.New("'email' is invalid")
			}

			break
		}
		case entity.ContactTypeSMS: {
			if nt.PhoneNumber == "" {
				return errors.New("'phoneNumber' must be given")
			} else {
				rgx, err := regexp.Compile("^\\+\\d+$")
				if err != nil || !rgx.MatchString(nt.PhoneNumber) {
					return errors.New("'phoneNumber' is invalid")
				}
			}

			break
		}
		case entity.ContactTypePush: {
			if nt.FCMRegistrationToken == "" {
				return errors.New("'fcmRegistrationToken' must be given")
			}
			break
		}
	default:
		return errors.New("'contactType' must be one of email/push/sms")
	}

	return nil
}

func (nt *NotificationTarget) GetContactInfo() *string {
	if nt.Email != "" {
		return &nt.Email
	} else if nt.PhoneNumber != "" {
		return &nt.PhoneNumber
	} else if nt.FCMRegistrationToken != "" {
		return &nt.FCMRegistrationToken
	}
	return nil
}

func (nt *NotificationTarget) GetContactType() *string {
	if nt.Email != "" {
		return &entity.ContactTypeEmail
	} else if nt.PhoneNumber != "" {
		return &entity.ContactTypeSMS
	} else if nt.FCMRegistrationToken != "" {
		return &entity.ContactTypePush
	}
	return nil
}