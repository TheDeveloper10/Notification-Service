package dto

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

type TemplateBodyRequest struct {
	iface.IRequestEntity[entity.TemplateBody]
	Email *string `json:"email"`
	SMS   *string `json:"sms"`
	Push  *string `json:"push"`
}

func (tbr *TemplateBodyRequest) Validate() iface.IErrorList {
	errs := util.NewErrorList()

	if tbr.Email == nil && tbr.SMS == nil && tbr.Push == nil {
		errs.AddErrorFromString("At least one of 'email', 'sms', 'push' must be passed")
	}

	if tbr.Email != nil && len(*tbr.Email) > 2048 {
		errs.AddErrorFromString("'email' must be at most 2048 characters")
	}

	if tbr.SMS != nil && len(*tbr.SMS) > 2048 {
		errs.AddErrorFromString("'sms' must be at most 2048 characters")
	}

	if tbr.Push != nil && len(*tbr.Push) > 2048 {
		errs.AddErrorFromString("'push' must be at most 2048 characters")
	}

	return errs
}

func (tbr *TemplateBodyRequest) ToEntity() entity.TemplateBody {
	return entity.TemplateBody{
		Email: tbr.Email,
		SMS:   tbr.SMS,
		Push:  tbr.Push,
	}
}