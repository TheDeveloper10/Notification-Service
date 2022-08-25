package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

type TemplatePlaceholderTest struct {
	Key 	   	   string
	Val            string
	ExpectedErrors int
}

func TestTemplatePlaceholder_Validate(t *testing.T) {
	testCases := []TemplatePlaceholderTest {
		{ "", "", 1},
		{ "j", "", 0 },
		{ "j", "a", 0 },
	}

	RunTemplatePlaceholderTest(0, nil, nil, 1, t)

	ranTests := 0
	for _, testCase := range testCases {
		ranTests++
		RunTemplatePlaceholderTest(
			ranTests,
			&testCase.Key,
			&testCase.Val,
			testCase.ExpectedErrors,
			t,
		)
	}
}

func RunTemplatePlaceholderTest(id int, key *string, val *string, expectedErrors int, t *testing.T) {
	req := TemplatePlaceholder{
		Key: key,
		Value: val,
	}

	err := req.Validate()
	actualErrors := 0
	if err != nil {
		actualErrors++
	}
	if expectedErrors != actualErrors {
		test.LogError(id, expectedErrors, actualErrors, t)
	}
}