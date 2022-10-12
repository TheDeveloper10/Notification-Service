package dto

import (
	"net/mail"
	"notification-service/internal/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"regexp"
)

type NotificationTarget struct {
	iface.IRequest
	Email                *string               `json:"email"`
	PhoneNumber          *string               `json:"phoneNumber"`
	FCMRegistrationToken *string       		   `json:"fcmRegistrationToken"`
	Placeholders         []TemplatePlaceholder `json:"placeholders"`
}

func (nt *NotificationTarget) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if nt.Email == nil && nt.PhoneNumber == nil && nt.FCMRegistrationToken == nil {
		errs.AddErrorFromString("You must pass an 'email', 'phoneNumber' or 'fcmRegistrationToken'")
	} else {
		if nt.Email != nil {
			if _, err := mail.ParseAddress(*nt.Email); err != nil {
				errs.AddErrorFromString("'email' is invalid")
			}
		}

		if nt.PhoneNumber != nil {
			rgx, err := regexp.Compile("^\\+\\d+$")
			if err != nil || !rgx.MatchString(*nt.PhoneNumber) {
				return errs.AddErrorFromString("'phoneNumber' is invalid")
			}
		}
	}

	return errs
}

func (nt *NotificationTarget) ValidateAgainstTemplate(template *entity.TemplateEntity) iface.IErrorList {
	errs := util.NewErrorList()

	if nt.Email != nil && template.Body.Email == nil {
		errs.AddErrorFromString("Email template is not set but an email was provided!")
	}

	if nt.PhoneNumber != nil && template.Body.SMS == nil {
		errs.AddErrorFromString("SMS template is not set but a phone number was provided!")
	}

	if nt.FCMRegistrationToken != nil && template.Body.Push == nil {
		errs.AddErrorFromString("Push template is not set but an FCM Registratation Token was provided!")
	}

	return errs
}