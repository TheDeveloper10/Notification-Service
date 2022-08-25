package test

import (
	"notification-service/internal/util/iface"
	"testing"
)


func RunTest(req iface.IRequest, testId int, expectedErrors int, t *testing.T) {
	errors := req.Validate()
	actualErrors := errors.ErrorsCount()
	if expectedErrors != actualErrors {
		LogError(testId, expectedErrors, actualErrors, t)
	}
}

func LogError(testId int, expected int, actual int, t *testing.T) {
	t.Errorf(
		"Error: expected %d errors but got %d on test %d",
		expected, actual, testId,
	)
}