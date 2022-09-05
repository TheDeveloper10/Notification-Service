package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestCreateTemplateRequest_Validate(t *testing.T) {
	testCases := []test.Case {
		{ 4, &CreateTemplateRequest{} },
		{ 4, &CreateTemplateRequest{ContactType: "", Template: "", Language: "", Type: "" } },
		{ 4, &CreateTemplateRequest{ContactType: "john", Template: "", Language: "", Type: "" } },
		{ 3, &CreateTemplateRequest{ContactType: "email", Template: "", Language: "", Type: "" } },
		{ 3, &CreateTemplateRequest{ContactType: "email", Template: "", Language: "Chinese", Type: "" } },
		{ 2, &CreateTemplateRequest{ContactType: "email", Template: "", Language: "EN", Type: "" } },
		{ 2, &CreateTemplateRequest{ContactType: "email", Template: "", Language: "EN", Type: "RandomType" } },
		{ 0, &CreateTemplateRequest{ContactType: "email", Template: "1", Language: "EN", Type: "type" } },
	}

	test.RunTestCases(&testCases, t)
}
