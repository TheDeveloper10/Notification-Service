package test

import (
	"notification-service/internal/util/iface"
	"testing"
)

type RequestTestCase struct {
	ExpectedErrors int
	Data iface.IRequest
}

func (rtc *RequestTestCase) RunTest(testId int, t *testing.T) {
	errs := rtc.Data.Validate()
	errCount := 0
	if errs != nil {
		errCount = errs.ErrorsCount()
	}
	if errCount != rtc.ExpectedErrors {
		rtc.LogError(testId, errCount, t)
	}
}

func (rtc *RequestTestCase) LogError(testId int, actual int, t *testing.T) {
	t.Errorf(
		"Test Id: %d\tExpected Errors: %d\tReceived Errors: %d",
		testId, rtc.ExpectedErrors, actual,
	)
}

func RunRequestTestCases(cases *[]RequestTestCase, t *testing.T) {
	for testId, testCase := range *cases {
		testCase.RunTest(testId, t)
	}
}