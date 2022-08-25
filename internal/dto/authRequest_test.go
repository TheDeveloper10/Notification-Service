package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

type AuthRequestTest struct {
	ClientId 	   string
	ClientSecret   string
	ExpectedErrors int
}

func TestAuthRequest_Validate(t *testing.T) {
	testCases := []AuthRequestTest {
		{ ClientId: "", ClientSecret: "", ExpectedErrors: 2},
		{ ClientId: "1", ClientSecret: "1", ExpectedErrors: 2},
		{ ClientId: "1234567890123456", ClientSecret: "1", ExpectedErrors: 1},
		{ ClientId: "1234567890123456", ClientSecret: "L0NYtEwFNmZS28eSeTLK37CLWPckRrCcsbTFUPI3dw2rdlwK4rhxj4epRCh969qFIao0W6OXrKngmTHPH0A5CqPhztijul05qMe22ErSGYcy6pcXzk8wN9JgKe8WwlwD", ExpectedErrors: 0},
	}

	RunAuthRequestTest(0, nil, nil, 2, t)

	ranTests := 0
	for _, testCase := range testCases {
		ranTests++
		RunAuthRequestTest(
			ranTests,
			&testCase.ClientId,
			&testCase.ClientSecret,
			testCase.ExpectedErrors,
			t,
		)
	}
}

func RunAuthRequestTest(id int, clientId *string, clientSecret *string, expectedErrors int, t *testing.T) {
	req := AuthRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}

	test.RunTest(&req, id, expectedErrors, t)
}