package test

import (
	"notification-service/internal/dto"
	"testing"
)

func TestCreateClientRequest_Validate(t *testing.T) {
	testCases := []RequestTestCase{
		{ ExpectedErrors: 0, Data: &dto.ClientPermissionsRequest{} },
		{ ExpectedErrors: 0, Data: &dto.ClientPermissionsRequest{ Permissions: []string{ } } },
		{ ExpectedErrors: 0, Data: &dto.ClientPermissionsRequest{ Permissions: []string{"agnsodg" } } },
		{ ExpectedErrors: 0, Data: &dto.ClientPermissionsRequest{ Permissions: []string{"update_templates\", \"read_templates" } } },
	}

	RunRequestTestCases(&testCases, t)
}
