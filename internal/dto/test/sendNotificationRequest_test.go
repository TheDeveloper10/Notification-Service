package test

import (
	"notification-service/internal/dto"
	"notification-service/internal/util/testutils"
	"testing"
)

func TestSendNotificationRequest_Validate(t *testing.T) {
	s := func(str string) *string { return &str }
	targets := []dto.NotificationTarget{
		{ Email: s("testutils@example.com") },
	}

	testCases := []testutils.RequestTestCase{
		{ ExpectedErrors: 4, Data: &dto.SendNotificationRequest{}},
		{ ExpectedErrors: 3, Data: &dto.SendNotificationRequest{Targets: targets}},
		{ ExpectedErrors: 3, Data: &dto.SendNotificationRequest{ AppID: "", TemplateID: 0, Title: "", Targets: targets } },
		{ ExpectedErrors: 2, Data: &dto.SendNotificationRequest{ AppID: "q", TemplateID: 0, Title: "", Targets: targets } },
		{ ExpectedErrors: 2, Data: &dto.SendNotificationRequest{ AppID: "w", TemplateID: -5, Title: "", Targets: targets } },
		{ ExpectedErrors: 1, Data: &dto.SendNotificationRequest{ AppID: "w", TemplateID: 5, Title: "", Targets: targets } },
		{ ExpectedErrors: 0, Data: &dto.SendNotificationRequest{ AppID: "w", TemplateID: 5, Title: "rt", Targets: targets } },
	}

	testutils.RunRequestTestCases(&testCases, t)
}
