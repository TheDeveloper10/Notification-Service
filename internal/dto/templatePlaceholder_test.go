package dto

import (
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"notification-service/internal/util/test"
	"testing"
)

type TemplatePlaceholderTest struct {
	Key string
	Val string
}

func (tpt *TemplatePlaceholderTest) Validate() iface.IErrorList {
	errs := util.NewErrorList()
	tp := TemplatePlaceholder{
		Key: tpt.Key,
		Value: tpt.Val,
	}
	errs.AddError(tp.Validate())
	return errs
}

func TestTemplatePlaceholder_Validate(t *testing.T) {
	testCases := []test.RequestTestCase{
		{ ExpectedErrors: 1, Data: &TemplatePlaceholderTest{} },
		{ ExpectedErrors: 1, Data: &TemplatePlaceholderTest{"", ""} },
		{ ExpectedErrors: 0, Data: &TemplatePlaceholderTest{"j", ""} },
		{ ExpectedErrors: 0, Data: &TemplatePlaceholderTest{"j", "a"} },
	}

	test.RunRequestTestCases(&testCases, t)
}
