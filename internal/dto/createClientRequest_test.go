package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestCreateClientRequest_Validate(t *testing.T) {
	testCases := []test.Case {
		{ 0, &CreateClientRequest{} },
		{ 0, &CreateClientRequest{ Permissions: []string{ } } },
		{ 0, &CreateClientRequest{ Permissions: []string{ "agnsodg" } } },
		{ 0, &CreateClientRequest{ Permissions: []string{ "update_templates\", \"read_templates" } } },
	}

	test.RunTestCases(&testCases, t)
}
