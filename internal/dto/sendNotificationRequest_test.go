package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestSendNotificationRequest_Validate(t *testing.T) {
	s := func(str string) *string { return &str }
	targets := []NotificationTarget{
		{ Email: s("test@example.com") },
	}

	testCases := []test.RequestTestCase{
		{ ExpectedErrors: 4, Data: &SendNotificationRequest{}},
		{ ExpectedErrors: 3, Data: &SendNotificationRequest{Targets: targets}},
		{ ExpectedErrors: 3, Data: &SendNotificationRequest{ AppID: "", TemplateID: 0, Title: "", Targets: targets } },
		{ ExpectedErrors: 2, Data: &SendNotificationRequest{ AppID: "q", TemplateID: 0, Title: "", Targets: targets } },
		{ ExpectedErrors: 2, Data: &SendNotificationRequest{ AppID: "w", TemplateID: -5, Title: "", Targets: targets } },
		{ ExpectedErrors: 1, Data: &SendNotificationRequest{ AppID: "w", TemplateID: 5, Title: "", Targets: targets } },
		{ ExpectedErrors: 0, Data: &SendNotificationRequest{ AppID: "w", TemplateID: 5, Title: "rt", Targets: targets } },
	}

	test.RunRequestTestCases(&testCases, t)
}
