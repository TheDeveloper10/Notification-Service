package test

import (
	dto2 "notification-service/internal/data/dto"
	"notification-service/internal/helper"
	"testing"
)

func TestCreateTemplateRequest_Validate(t *testing.T) {
	helper.LoadConfig("../../../../" + helper.ServiceConfigPath)

	s := func(str string) *string { return &str }
	testCases := []RequestTestCase{
		{ ExpectedErrors: 3, Data: &dto2.CreateTemplateRequest{} },
		{ ExpectedErrors: 3, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{Email: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{SMS: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{SMS: s("test 2"), Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{Email: s("test 2"), SMS: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{Email: s("test 2"), Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{Email: s("test 2"), Push: s("test template"), SMS: s("test 3")}, Language: "", Type: "" } },
		{ ExpectedErrors: 1, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{Email: s("test 2")}, Language: "BG", Type: "" } },
		{ ExpectedErrors: 1, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{Email: s("test 4")}, Language: "Bulgarian", Type: "qwe" } },
		{ ExpectedErrors: 0, Data: &dto2.CreateTemplateRequest{Body: dto2.TemplateBodyRequest{Email: s("test 2")}, Language: "BG", Type: "qwe" } },
	}

	RunRequestTestCases(&testCases, t)
}
