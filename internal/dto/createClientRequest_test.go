package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestCreateClientRequest_Validate(t *testing.T) {
	testCases := []test.RequestTestCase{
		{ 0, &CreateClientRequest{} },
		{ 0, &CreateClientRequest{ Permissions: []string{ } } },
		{ 0, &CreateClientRequest{ Permissions: []string{ "agnsodg" } } },
		{ 0, &CreateClientRequest{ Permissions: []string{ "update_templates\", \"read_templates" } } },
	}

	test.RunRequestTestCases(&testCases, t)
}
