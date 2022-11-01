package test

import (
	"notification-service/internal/dto"
	"notification-service/internal/util/testutils"
	"testing"
)

func TestNotificationTarget_Validate(t *testing.T) {
	s := func(str string) *string { return &str }
	testCases := []testutils.RequestTestCase{
		{ExpectedErrors: 1, Data: &dto.NotificationTarget{}},
		{ExpectedErrors: 1, Data: &dto.NotificationTarget{Email: nil, PhoneNumber: nil, FCMRegistrationToken: nil}},
		{ExpectedErrors: 1, Data: &dto.NotificationTarget{Email: s("john"), PhoneNumber: nil, FCMRegistrationToken: nil}},
		{ExpectedErrors: 0, Data: &dto.NotificationTarget{Email: s("john@example.com"), PhoneNumber: nil, FCMRegistrationToken: nil}},
		{ExpectedErrors: 1, Data: &dto.NotificationTarget{Email: s("john@example.com"), PhoneNumber: s("087734125"), FCMRegistrationToken: nil}},
		{ExpectedErrors: 1, Data: &dto.NotificationTarget{Email: nil, PhoneNumber: s("087734125"), FCMRegistrationToken: nil}},
		{ExpectedErrors: 1, Data: &dto.NotificationTarget{Email: nil, PhoneNumber: s("087734125"), FCMRegistrationToken: nil}},
		{ExpectedErrors: 0, Data: &dto.NotificationTarget{Email: nil, PhoneNumber: s("+35987734125"), FCMRegistrationToken: nil}},
		{ExpectedErrors: 0, Data: &dto.NotificationTarget{Email: s("john@example.com"), PhoneNumber: s("+35987734125"), FCMRegistrationToken: s("pahsfopiHOIho")}},
		{ExpectedErrors: 0, Data: &dto.NotificationTarget{Email: nil, PhoneNumber: nil, FCMRegistrationToken: s("PIAHgfousdghouewht")}},
	}

	testutils.RunRequestTestCases(&testCases, t)
}