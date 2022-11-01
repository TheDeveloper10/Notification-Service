package test

import (
	"notification-service/internal/dto"
	"testing"
)

func TestAuthRequest_Validate(t *testing.T) {
	testCases := []RequestTestCase{
		{ ExpectedErrors: 2, Data: &dto.AuthRequest{}},
		{ ExpectedErrors: 2, Data: &dto.AuthRequest{ ClientId: "", ClientSecret: "" }},
		{ ExpectedErrors: 2, Data: &dto.AuthRequest{ ClientId: "1", ClientSecret: "" }},
		{ ExpectedErrors: 2, Data: &dto.AuthRequest{ ClientId: "", ClientSecret: "1" }},
		{ ExpectedErrors: 2, Data: &dto.AuthRequest{ ClientId: "1", ClientSecret: "1" }},
		{ ExpectedErrors: 1, Data: &dto.AuthRequest{ ClientId: "1234567890123456", ClientSecret: "1" }},
		{ ExpectedErrors: 0, Data: &dto.AuthRequest{ ClientId: "1234567890123456", ClientSecret: "L0NYtEwFNmZS28eSeTLK37CLWPckRrCcsbTFUPI3dw2rdlwK4rhxj4epRCh969qFIao0W6OXrKngmTHPH0A5CqPhztijul05qMe22ErSGYcy6pcXzk8wN9JgKe8WwlwD" }},
	}

	RunRequestTestCases(&testCases, t)
}
