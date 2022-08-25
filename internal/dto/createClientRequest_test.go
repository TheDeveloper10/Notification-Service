package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

type CreateClientRequestTest struct {
	Permissions []string
	ExpectedErrors int
}

func TestCreateClientRequest_Validate(t *testing.T) {
	testCases := []CreateClientRequestTest {
		{ Permissions: []string{ "agnsodg" }, ExpectedErrors: 0},
		{ Permissions: []string{ "update_templates", "read_templates" }, ExpectedErrors: 0},
	}

	RunCreateClientRequestTest(0, nil, 1, t)

	ranTests := 0
	for _, testCase := range testCases {
		ranTests++
		RunCreateClientRequestTest(
			ranTests,
			&testCase.Permissions,
			testCase.ExpectedErrors,
			t,
		)
	}
}

func RunCreateClientRequestTest(id int, permissions *[]string, expectedErrors int, t *testing.T) {
	req := CreateClientRequest{
		Permissions: permissions,
	}

	test.RunTest(&req, id, expectedErrors, t)
}