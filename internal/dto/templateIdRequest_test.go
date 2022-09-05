package dto

import (
	"notification-service/internal/util/test"
	"testing"
)

func TestTemplateIdRequest_Validate(t *testing.T) {
	testCases := []test.Case {
		{ 1, &TemplateIdRequest{} },
		{ 1, &TemplateIdRequest{Id: 0} },
		{ 1, &TemplateIdRequest{Id: -1} },
		{ 1, &TemplateIdRequest{Id: -15250} },
		{ 0, &TemplateIdRequest{Id: 15250} },
	}

	test.RunTestCases(&testCases, t)
}
