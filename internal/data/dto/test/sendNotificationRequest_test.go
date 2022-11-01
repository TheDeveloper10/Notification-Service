package test

import (
	dto2 "notification-service/internal/data/dto"
	"testing"
)

func TestSendNotificationRequest_Validate(t *testing.T) {
	s := func(str string) *string { return &str }
	targets := []dto2.NotificationTarget{
		{ Email: s("test@example.com") },
	}

	testCases := []RequestTestCase{
		{ ExpectedErrors: 4, Data: &dto2.SendNotificationRequest{}},
		{ ExpectedErrors: 3, Data: &dto2.SendNotificationRequest{Targets: targets}},
		{ ExpectedErrors: 3, Data: &dto2.SendNotificationRequest{ AppID: "", TemplateID: 0, Title: "", Targets: targets } },
		{ ExpectedErrors: 2, Data: &dto2.SendNotificationRequest{ AppID: "q", TemplateID: 0, Title: "", Targets: targets } },
		{ ExpectedErrors: 2, Data: &dto2.SendNotificationRequest{ AppID: "w", TemplateID: -5, Title: "", Targets: targets } },
		{ ExpectedErrors: 1, Data: &dto2.SendNotificationRequest{ AppID: "w", TemplateID: 5, Title: "", Targets: targets } },
		{ ExpectedErrors: 0, Data: &dto2.SendNotificationRequest{ AppID: "w", TemplateID: 5, Title: "rt", Targets: targets } },
	}

	RunRequestTestCases(&testCases, t)
}
