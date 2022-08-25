package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

type NotificationTargetTest struct {
	ContactType          string
	Email 	    		 string
	PhoneNumber 		 string
	FCMRegistrationToken string
	ExpectedErrors       int
}

func TestNotificationTarget_Validate(t *testing.T) {
	testCases := []NotificationTargetTest {
		{ "email", "john", "", "", 1 },
		{ "email", "john@abv.bg", "", "", 0 },
		{ "email", "john@abv.bg", "087734125", "", 0 },
		{ "email", "", "087734125", "", 1 },
		{ "sms", "", "087734125", "", 1 },
		{ "sms", "", "+35987734125", "", 0 },
		{ "push", "", "+35987734125", "", 1 },
		{ "push", "", "", "PIAHgfousdghouewht", 0 },
	}

	RunNotificationTargetTest(0, nil, nil, nil, nil, 1, t)

	ranTests := 0
	for _, testCase := range testCases {
		ranTests++

		var contactType *string = nil
		var email *string = nil
		var phoneNumber *string = nil
		var fcmRegistrationToken *string = nil

		if testCase.ContactType != "" {
			contactType = &testCase.ContactType
		}
		if testCase.Email != "" {
			email = &testCase.Email
		}
		if testCase.PhoneNumber != "" {
			phoneNumber = &testCase.PhoneNumber
		}
		if testCase.FCMRegistrationToken != "" {
			fcmRegistrationToken = &testCase.FCMRegistrationToken
		}

		RunNotificationTargetTest(
			ranTests,
			contactType,
			email,
			phoneNumber,
			fcmRegistrationToken,
			testCase.ExpectedErrors,
			t,
		)
	}
}

func RunNotificationTargetTest(id int, contactType *string, email *string, phoneNumber *string, fcmRegistrationToken *string, expectedErrors int, t *testing.T) {
	req := NotificationTarget{
		Email: email,
		PhoneNumber: phoneNumber,
		FCMRegistrationToken: fcmRegistrationToken,
	}

	err := req.Validate(contactType)
	actualErrors := 0
	if err != nil {
		actualErrors++
	}
	if expectedErrors != actualErrors {
		test.LogError(id, expectedErrors, actualErrors, t)
	}
}