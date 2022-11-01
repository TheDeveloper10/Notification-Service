package test

import (
	"notification-service/internal/dto"
	"notification-service/internal/util/testutils"
	"testing"
)

func TestTemplateIdRequest_Validate(t *testing.T) {
	testCases := []testutils.RequestTestCase{
		{ ExpectedErrors: 1, Data: &dto.TemplateIdRequest{} },
		{ ExpectedErrors: 1, Data: &dto.TemplateIdRequest{Id: 0} },
		{ ExpectedErrors: 1, Data: &dto.TemplateIdRequest{Id: -1} },
		{ ExpectedErrors: 1, Data: &dto.TemplateIdRequest{Id: -15250} },
		{ ExpectedErrors: 0, Data: &dto.TemplateIdRequest{Id: 15250} },
	}

	testutils.RunRequestTestCases(&testCases, t)
}
