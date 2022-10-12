package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestUpdateClientRequest_Validate(t *testing.T) {
	testCases := []test.RequestTestCase{
		{ ExpectedErrors: 1, Data: &UpdateClientRequest{} },
		{ ExpectedErrors: 1, Data: &UpdateClientRequest{ Permissions: []string{ } } },
		{ ExpectedErrors: 1, Data: &UpdateClientRequest{ Permissions: []string{ "agnsodg" } } },
		{ ExpectedErrors: 1, Data: &UpdateClientRequest{ Permissions: []string{ "update_templates\", \"read_templates" } } },
		{ ExpectedErrors: 0, Data: &UpdateClientRequest{ ClientID: "a" } },
		{ ExpectedErrors: 0, Data: &UpdateClientRequest{ ClientID: "a", Permissions: []string{ } } },
		{ ExpectedErrors: 0, Data: &UpdateClientRequest{ ClientID: "a", Permissions: []string{ "agnsodg" } } },
		{ ExpectedErrors: 0, Data: &UpdateClientRequest{ ClientID: "a", Permissions: []string{ "update_templates\", \"read_templates" } } },
	}

	test.RunRequestTestCases(&testCases, t)
}
