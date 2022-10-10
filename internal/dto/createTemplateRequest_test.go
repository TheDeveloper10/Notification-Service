package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestCreateTemplateRequest_Validate(t *testing.T) {
	testCases := []test.RequestTestCase{
		{ ExpectedErrors: 4, Data: &CreateTemplateRequest{} },
		{ ExpectedErrors: 4, Data: &CreateTemplateRequest{ContactType: "", Template: "", Language: "", Type: "" } },
		{ ExpectedErrors: 4, Data: &CreateTemplateRequest{ContactType: "john", Template: "", Language: "", Type: "" } },
		{ ExpectedErrors: 3, Data: &CreateTemplateRequest{ContactType: "email", Template: "", Language: "", Type: "" } },
		{ ExpectedErrors: 3, Data: &CreateTemplateRequest{ContactType: "email", Template: "", Language: "Chinese", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{ContactType: "email", Template: "", Language: "EN", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{ContactType: "email", Template: "", Language: "EN", Type: "RandomType" } },
		{ ExpectedErrors: 0, Data: &CreateTemplateRequest{ContactType: "email", Template: "1", Language: "EN", Type: "type" } },
	}

	test.RunRequestTestCases(&testCases, t)
}
