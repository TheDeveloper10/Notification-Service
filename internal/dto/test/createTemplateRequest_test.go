package test

import (
	"notification-service/internal/dto"
	"notification-service/internal/helper"
	"testing"
)

func TestCreateTemplateRequest_Validate(t *testing.T) {
	helper.LoadConfig("../../../" + helper.ServiceConfigPath)

	s := func(str string) *string { return &str }
	testCases := []RequestTestCase{
		{ ExpectedErrors: 3, Data: &dto.CreateTemplateRequest{} },
		{ ExpectedErrors: 3, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{SMS: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{SMS: s("test 2"), Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("test 2"), SMS: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("test 2"), Push: s("test template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("test 2"), Push: s("test template"), SMS: s("test 3")}, Language: "", Type: "" } },
		{ ExpectedErrors: 1, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("test 2")}, Language: "BG", Type: "" } },
		{ ExpectedErrors: 1, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("test 4")}, Language: "Bulgarian", Type: "qwe" } },
		{ ExpectedErrors: 0, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("test 2")}, Language: "BG", Type: "qwe" } },
	}

	RunRequestTestCases(&testCases, t)
}
