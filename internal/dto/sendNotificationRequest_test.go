package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestSendNotificationRequest_Validate(t *testing.T) {
	targets := []NotificationTarget{
		{ Email: "test@example.com" },
	}

	testCases := []test.RequestTestCase{
		{ ExpectedErrors: 5, Data: &SendNotificationRequest{}},
		{ ExpectedErrors: 4, Data: &SendNotificationRequest{Targets: targets}},
		{ ExpectedErrors: 4, Data: &SendNotificationRequest{ AppID: "", TemplateID: 0, ContactType: "", Title: "", Targets: targets } },
		{ ExpectedErrors: 3, Data: &SendNotificationRequest{ AppID: "q", TemplateID: 0, ContactType: "", Title: "", Targets: targets } },
		{ ExpectedErrors: 3, Data: &SendNotificationRequest{ AppID: "w", TemplateID: -5, ContactType: "", Title: "", Targets: targets } },
		{ ExpectedErrors: 2, Data: &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "", Title: "", Targets: targets } },
		{ ExpectedErrors: 2, Data: &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "r", Title: "", Targets: targets } },
		{ ExpectedErrors: 1, Data: &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "email", Title: "", Targets: targets } },
		{ ExpectedErrors: 1, Data: &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "sms", Title: "", Targets: targets } },
		{ ExpectedErrors: 1, Data: &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "push", Title: "", Targets: targets } },
		{ ExpectedErrors: 0, Data: &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "push", Title: "rt", Targets: targets } },
	}

	test.RunRequestTestCases(&testCases, t)
}
