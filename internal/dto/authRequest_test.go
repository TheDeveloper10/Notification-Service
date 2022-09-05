package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestAuthRequest_Validate(t *testing.T) {
	testCases := []test.RequestTestCase{
		{ 2, &AuthRequest{}},
		{ 2, &AuthRequest{ ClientId: "", ClientSecret: "" }},
		{ 2, &AuthRequest{ ClientId: "1", ClientSecret: "" }},
		{ 2, &AuthRequest{ ClientId: "", ClientSecret: "1" }},
		{ 2, &AuthRequest{ ClientId: "1", ClientSecret: "1" }},
		{ 1, &AuthRequest{ ClientId: "1234567890123456", ClientSecret: "1" }},
		{ 0, &AuthRequest{ ClientId: "1234567890123456", ClientSecret: "L0NYtEwFNmZS28eSeTLK37CLWPckRrCcsbTFUPI3dw2rdlwK4rhxj4epRCh969qFIao0W6OXrKngmTHPH0A5CqPhztijul05qMe22ErSGYcy6pcXzk8wN9JgKe8WwlwD" }},
	}

	test.RunRequestTestCases(&testCases, t)
}
