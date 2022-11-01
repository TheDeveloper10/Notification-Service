package test

import (
	"notification-service/internal/data/dto"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"testing"
)

type TemplatePlaceholderTest struct {
	Key string
	Val string
}

func (tpt *TemplatePlaceholderTest) Validate() iface.IErrorList {
	errs := util.NewErrorList()
	tp := dto.TemplatePlaceholder{
		Key: tpt.Key,
		Value: tpt.Val,
	}
	errs.AddError(tp.Validate())
	return errs
}

func TestTemplatePlaceholder_Validate(t *testing.T) {
	testCases := []RequestTestCase{
		{ ExpectedErrors: 1, Data: &TemplatePlaceholderTest{} },
		{ ExpectedErrors: 1, Data: &TemplatePlaceholderTest{"", ""} },
		{ ExpectedErrors: 0, Data: &TemplatePlaceholderTest{"j", ""} },
		{ ExpectedErrors: 0, Data: &TemplatePlaceholderTest{"j", "a"} },
	}

	RunRequestTestCases(&testCases, t)
}
