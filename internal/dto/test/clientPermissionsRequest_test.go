package test

import (
	"notification-service/internal/dto"
	"notification-service/internal/util/testutils"
	"testing"
)

func TestCreateClientRequest_Validate(t *testing.T) {
	testCases := []testutils.RequestTestCase{
		{ ExpectedErrors: 0, Data: &dto.ClientPermissionsRequest{} },
		{ ExpectedErrors: 0, Data: &dto.ClientPermissionsRequest{ Permissions: []string{ } } },
		{ ExpectedErrors: 0, Data: &dto.ClientPermissionsRequest{ Permissions: []string{"agnsodg" } } },
		{ ExpectedErrors: 0, Data: &dto.ClientPermissionsRequest{ Permissions: []string{"update_templates\", \"read_templates" } } },
	}

	testutils.RunRequestTestCases(&testCases, t)
}
