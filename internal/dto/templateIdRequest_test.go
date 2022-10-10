package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestTemplateIdRequest_Validate(t *testing.T) {
	testCases := []test.RequestTestCase{
		{ ExpectedErrors: 1, Data: &TemplateIdRequest{} },
		{ ExpectedErrors: 1, Data: &TemplateIdRequest{Id: 0} },
		{ ExpectedErrors: 1, Data: &TemplateIdRequest{Id: -1} },
		{ ExpectedErrors: 1, Data: &TemplateIdRequest{Id: -15250} },
		{ ExpectedErrors: 0, Data: &TemplateIdRequest{Id: 15250} },
	}

	test.RunRequestTestCases(&testCases, t)
}
