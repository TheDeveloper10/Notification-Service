package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestCreateClientRequest_Validate(t *testing.T) {
	testCases := []test.RequestTestCase{
		{ ExpectedErrors: 0, Data: &ClientPermissionsRequest{} },
		{ ExpectedErrors: 0, Data: &ClientPermissionsRequest{ Permissions: []string{ } } },
		{ ExpectedErrors: 0, Data: &ClientPermissionsRequest{ Permissions: []string{"agnsodg" } } },
		{ ExpectedErrors: 0, Data: &ClientPermissionsRequest{ Permissions: []string{"update_templates\", \"read_templates" } } },
	}

	test.RunRequestTestCases(&testCases, t)
}
