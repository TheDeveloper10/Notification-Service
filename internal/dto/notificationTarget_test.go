package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestNotificationTarget_Validate(t *testing.T) {
	s := func(str string) *string { return &str }
	testCases := []test.RequestTestCase{
		{ExpectedErrors: 1, Data: &NotificationTarget{}},
		{ExpectedErrors: 1, Data: &NotificationTarget{Email: nil, PhoneNumber: nil, FCMRegistrationToken: nil}},
		{ExpectedErrors: 1, Data: &NotificationTarget{Email: s("john"), PhoneNumber: nil, FCMRegistrationToken: nil}},
		{ExpectedErrors: 0, Data: &NotificationTarget{Email: s("john@example.com"), PhoneNumber: nil, FCMRegistrationToken: nil}},
		{ExpectedErrors: 1, Data: &NotificationTarget{Email: s("john@example.com"), PhoneNumber: s("087734125"), FCMRegistrationToken: nil}},
		{ExpectedErrors: 1, Data: &NotificationTarget{Email: nil, PhoneNumber: s("087734125"), FCMRegistrationToken: nil}},
		{ExpectedErrors: 1, Data: &NotificationTarget{Email: nil, PhoneNumber: s("087734125"), FCMRegistrationToken: nil}},
		{ExpectedErrors: 0, Data: &NotificationTarget{Email: nil, PhoneNumber: s("+35987734125"), FCMRegistrationToken: nil}},
		{ExpectedErrors: 0, Data: &NotificationTarget{Email: s("john@example.com"), PhoneNumber: s("+35987734125"), FCMRegistrationToken: s("pahsfopiHOIho")}},
		{ExpectedErrors: 0, Data: &NotificationTarget{Email: nil, PhoneNumber: nil, FCMRegistrationToken: s("PIAHgfousdghouewht")}},
	}

	test.RunRequestTestCases(&testCases, t)
}