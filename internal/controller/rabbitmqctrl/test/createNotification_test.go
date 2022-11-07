package test

import (
	"notification-service/internal/controller/rabbitmqctrl"
	"notification-service/internal/data/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"testing"
)

func TestCreateNotificationV1Controller(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	notificationRepository := repository.NewMockNotificationRepository()
	ctrl := rabbitmqctrl.NewCreateNotificationV1Controller(templateRepository, notificationRepository)

	s := func(str string) *string { return &str }
	testCases := []ControllerTestCase{
		{
			Controller: ctrl,
			NoResponseComparison: true,
			ExpectedAck: true,
			Body: dto.SendNotificationRequest{},
		},
		{
			Controller: ctrl,
			NoResponseComparison: false,
			ExpectedAck: true,
			ExpectedResponse: util.ErrorListFromTextError("Template not found"),
			Body: dto.SendNotificationRequest{
				AppID: "test",
				TemplateID: 1,
				Title: "magic",
				Targets: []dto.NotificationTarget{
					{
						Email: s("test@example.com"),
					},
				},
			},
		},
		{
			Controller: ctrl,
			NoResponseComparison: false,
			ExpectedAck: true,
			ExpectedResponse: &dto.SendNotificationsError{
				Errors: []dto.SendNotificationErrorData{
					{
						TargetId: 0,
						Messages: []string{
							"Email template is not set but an email was provided!",
						},
					},
				},
				SuccessfullySentNotifications: 0,
				FailedNotifications: 0,
			},
			Body: dto.SendNotificationRequest{
				AppID: "test",
				TemplateID: 2,
				Title: "magic",
				Targets: []dto.NotificationTarget{
					{
						Email: s("test@example.com"),
					},
				},
			},
		},
	}

	RunControllerTestCases(&testCases, t)
}