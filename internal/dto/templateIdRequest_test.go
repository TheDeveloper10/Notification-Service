package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

type TemplateIdRequestTest struct {
	Id 	           int
	ExpectedErrors int
}

func TestTemplateIdRequest_Validate(t *testing.T) {
	testCases := []TemplateIdRequestTest {
		{ Id: 0, ExpectedErrors: 1},
		{ Id: -1, ExpectedErrors: 1},
		{ Id: -15250, ExpectedErrors: 1},
		{ Id: 15250, ExpectedErrors: 0},
	}

	RunTemplateIdRequestTest(0, nil, 1, t)

	ranTests := 0
	for _, testCase := range testCases {
		ranTests++
		RunTemplateIdRequestTest(
			ranTests,
			&testCase.Id,
			testCase.ExpectedErrors,
			t,
		)
	}
}

func RunTemplateIdRequestTest(testId int, id *int, expectedErrors int, t *testing.T) {
	req := TemplateIdRequest{
		Id: id,
	}

	test.RunTest(&req, testId, expectedErrors, t)
}