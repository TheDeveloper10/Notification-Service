package test

import (
	"notification-service/internal/data/dto"
	"testing"
)

func TestTemplateIdRequest_Validate(t *testing.T) {
	testCases := []RequestTestCase{
		{ ExpectedErrors: 1, Data: &dto.TemplateIdRequest{} },
		{ ExpectedErrors: 1, Data: &dto.TemplateIdRequest{Id: 0} },
		{ ExpectedErrors: 1, Data: &dto.TemplateIdRequest{Id: -1} },
		{ ExpectedErrors: 1, Data: &dto.TemplateIdRequest{Id: -15250} },
		{ ExpectedErrors: 0, Data: &dto.TemplateIdRequest{Id: 15250} },
	}

	RunRequestTestCases(&testCases, t)
}
