package dto

import (
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"notification-service/internal/util/test"
	"testing"
)

type NotificationTargetTest struct {
	iface.IRequest
	ContactType string
	Email string
	PhoneNumber string
	FCMRegistrationToken string
}

func (ntt *NotificationTargetTest) Validate() iface.IErrorList {
	errs := util.NewErrorList()
	nt := NotificationTarget{
		Email: ntt.Email,
		PhoneNumber: ntt.PhoneNumber,
		FCMRegistrationToken: ntt.FCMRegistrationToken,
	}
	errs.AddError(nt.Validate(&ntt.ContactType))
	return errs
}

func TestNotificationTarget_Validate(t *testing.T) {
	testCases := []test.RequestTestCase{
		{ExpectedErrors: 1, Data: &NotificationTargetTest{}},
		{ExpectedErrors: 1, Data: &NotificationTargetTest{ContactType: "", Email: "", PhoneNumber: "", FCMRegistrationToken: ""}},
		{ExpectedErrors: 1, Data: &NotificationTargetTest{ContactType: "email", Email: "john", PhoneNumber: "", FCMRegistrationToken: ""}},
		{ExpectedErrors: 0, Data: &NotificationTargetTest{ContactType: "email", Email: "john@abv.bg", PhoneNumber: "", FCMRegistrationToken: ""}},
		{ExpectedErrors: 0, Data: &NotificationTargetTest{ContactType: "email", Email: "john@abv.bg", PhoneNumber: "087734125", FCMRegistrationToken: ""}},
		{ExpectedErrors: 1, Data: &NotificationTargetTest{ContactType: "email", Email: "", PhoneNumber: "087734125", FCMRegistrationToken: ""}},
		{ExpectedErrors: 1, Data: &NotificationTargetTest{ContactType: "sms", Email: "", PhoneNumber: "087734125", FCMRegistrationToken: ""}},
		{ExpectedErrors: 0, Data: &NotificationTargetTest{ContactType: "sms", Email: "", PhoneNumber: "+35987734125", FCMRegistrationToken: ""}},
		{ExpectedErrors: 1, Data: &NotificationTargetTest{ContactType: "push", Email: "", PhoneNumber: "+35987734125", FCMRegistrationToken: ""}},
		{ExpectedErrors: 0, Data: &NotificationTargetTest{ContactType: "push", Email: "", PhoneNumber: "", FCMRegistrationToken: "PIAHgfousdghouewht"}},
	}

	test.RunRequestTestCases(&testCases, t)
}