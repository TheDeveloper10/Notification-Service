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
		{ 5, &SendNotificationRequest{}},
		{ 4, &SendNotificationRequest{Targets: targets}},
		{ 4, &SendNotificationRequest{ AppID: "", TemplateID: 0, ContactType: "", Title: "", Targets: targets } },
		{ 3, &SendNotificationRequest{ AppID: "q", TemplateID: 0, ContactType: "", Title: "", Targets: targets } },
		{ 3, &SendNotificationRequest{ AppID: "w", TemplateID: -5, ContactType: "", Title: "", Targets: targets } },
		{ 2, &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "", Title: "", Targets: targets } },
		{ 2, &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "r", Title: "", Targets: targets } },
		{ 1, &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "email", Title: "", Targets: targets } },
		{ 1, &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "sms", Title: "", Targets: targets } },
		{ 1, &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "push", Title: "", Targets: targets } },
		{ 0, &SendNotificationRequest{ AppID: "w", TemplateID: 5, ContactType: "push", Title: "rt", Targets: targets } },
	}

	test.RunRequestTestCases(&testCases, t)
}
