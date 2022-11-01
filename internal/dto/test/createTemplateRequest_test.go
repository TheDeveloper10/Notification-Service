package test

import (
	"notification-service/internal/dto"
	"notification-service/internal/helper"
	"notification-service/internal/util/testutils"
	"testing"
)

func TestCreateTemplateRequest_Validate(t *testing.T) {
	helper.LoadConfig("../../../" + helper.ServiceConfigPath)

	s := func(str string) *string { return &str }
	testCases := []testutils.RequestTestCase{
		{ ExpectedErrors: 3, Data: &dto.CreateTemplateRequest{} },
		{ ExpectedErrors: 3, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("testutils template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{SMS: s("testutils template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Push: s("testutils template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{SMS: s("testutils 2"), Push: s("testutils template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("testutils 2"), SMS: s("testutils template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("testutils 2"), Push: s("testutils template")}, Language: "", Type: "" } },
		{ ExpectedErrors: 2, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("testutils 2"), Push: s("testutils template"), SMS: s("testutils 3")}, Language: "", Type: "" } },
		{ ExpectedErrors: 1, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("testutils 2")}, Language: "BG", Type: "" } },
		{ ExpectedErrors: 1, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("testutils 4")}, Language: "Bulgarian", Type: "qwe" } },
		{ ExpectedErrors: 0, Data: &dto.CreateTemplateRequest{Body: dto.TemplateBodyRequest{Email: s("testutils 2")}, Language: "BG", Type: "qwe" } },
	}

	testutils.RunRequestTestCases(&testCases, t)
}
