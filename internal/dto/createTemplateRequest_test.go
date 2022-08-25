package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

type CreateTemplateRequestTest struct {
	ContactType    string `json:"contactType"`
	Template       string `json:"template"`
	Language       string `json:"language"`
	Type           string `json:"type"`
	ExpectedErrors int
}

func TestCreateTemplateRequest_Validate(t *testing.T) {
	testCases := []CreateTemplateRequestTest {
		{ "", "", "", "", 4 },
		{ "john", "", "", "", 4 },
		{ "email", "", "", "", 3 },
		{ "email", "", "Chinese", "", 3 },
		{ "email", "", "EN", "", 2 },
		{ "email", "", "EN", "RandomType", 2 },
		{ "email", "1", "EN", "type", 0 },
	}

	RunCreateTemplateRequestTest(0, nil, nil, nil, nil, 4, t)

	ranTests := 0
	for _, testCase := range testCases {
		ranTests++
		RunCreateTemplateRequestTest(
			ranTests,
			&testCase.ContactType,
			&testCase.Template,
			&testCase.Language,
			&testCase.Type,
			testCase.ExpectedErrors,
			t,
		)
	}
}

func RunCreateTemplateRequestTest(id int, contactType *string, template *string, language *string, _type *string, expectedErrors int, t *testing.T) {
	req := CreateTemplateRequest{
		ContactType: contactType,
		Template:    template,
		Language:    language,
		Type:        _type,
	}

	test.RunTest(&req, id, expectedErrors, t)
}