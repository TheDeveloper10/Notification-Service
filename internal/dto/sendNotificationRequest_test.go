package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

type SendNotificationRequestTest struct {
	AppID 	   	   		  string
	TemplateID     		  int
	ContactType    		  string
	Title 				  string
	ExpectedErrors 		  int
}

func TestSendNotificationRequest_Validate(t *testing.T) {
	testCases := []SendNotificationRequestTest {
		{ "", 0, "", "", 4 },
		{ "q", 0, "", "", 3 },
		{ "w", -5, "", "", 3 },
		{ "w", 5, "", "", 2 },
		{ "w", 5, "r", "", 2 },
		{ "w", 5, "email", "", 1 },
		{ "w", 5, "sms", "", 1 },
		{ "w", 5, "push", "", 1 },
		{ "w", 5, "push", "rt", 0 },
	}

	RunSendNotificationRequestTest(0, nil, nil, nil, nil, 4, t)

	ranTests := 0
	for _, testCase := range testCases {
		ranTests++
		RunSendNotificationRequestTest(
			ranTests,
			&testCase.AppID,
			&testCase.TemplateID,
			&testCase.ContactType,
			&testCase.Title,
			testCase.ExpectedErrors,
			t,
		)
	}
}

func RunSendNotificationRequestTest(id int,
									appID *string,
									templateID *int,
									contactType *string,
									title *string,
									expectedErrors int, t *testing.T) {
	m := "test@example.com"
	req := SendNotificationRequest{
		AppID: appID,
		TemplateID: templateID,
		ContactType: contactType,
		Title: title,
		Targets: []NotificationTarget{
			{Email: &m},
		},
		UniversalPlaceholders: []TemplatePlaceholder{},
	}

	test.RunTest(&req, id, expectedErrors, t)
}