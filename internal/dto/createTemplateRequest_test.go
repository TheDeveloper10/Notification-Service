package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestCreateTemplateRequest_Validate(t *testing.T) {
	s := func(str string) *string { return &str }
	testCases := []test.RequestTestCase{
		{ ExpectedErrors: 3, Data: &CreateTemplateRequest{} },
		{ ExpectedErrors: 3, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{Email: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{SMS: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{SMS: s("test 2"), Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{Email: s("test 2"), SMS: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{Email: s("test 2"), Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{Email: s("test 2"), Push: s("test template"), SMS: s("test 3")}, Language: "", Type: "" } },
		{ ExpectedErrors: 1, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{Email: s("test 2")}, Language: "BG", Type: "" } },
		{ ExpectedErrors: 1, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{Email: s("test 2")}, Language: "Bulgarian", Type: "qwe" } },
		{ ExpectedErrors: 0, Data: &CreateTemplateRequest{Body: TemplateBodyRequest{Email: s("test 2")}, Language: "BG", Type: "qwe" } },
	}

	test.RunRequestTestCases(&testCases, t)
}
