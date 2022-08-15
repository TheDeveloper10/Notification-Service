package clients

import (
	"notification-service/internal/helper"
	"notification-service/internal/util"
)

var MailClient *util.MailClient

func InitializeMailClient() {
	if MailClient != nil {
		return
	}

	MailClient = &util.MailClient{}
	MailClient.Init(helper.Config.Smtp.Host, helper.Config.Smtp.Port, helper.Config.Smtp.FromEmail, helper.Config.Smtp.FromPassword)
}
