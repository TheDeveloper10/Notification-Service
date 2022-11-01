package test

import (
	"notification-service/internal/helper"
	"strconv"
	"testing"
)

type listOfStringsJoinTest struct {
	data     helper.ListOfStrings
	operator string
	expectedResult string
}

func (losjt *listOfStringsJoinTest) performTest(testId int, t *testing.T) {
	result := losjt.data.Join(losjt.operator)

	if losjt.expectedResult != result {
		t.Error(strconv.Itoa(testId) + "\t\t Expected: " + losjt.expectedResult + "\t\tActual: " + result)
	}
}

func TestListOfStrings_Join(t *testing.T) {
	tests := []listOfStringsJoinTest{
		{ data: []string{ "a", "b", "c" }, operator: "", expectedResult: "abc" },
		{ data: []string{ "a", "b", "c" }, operator: " ", expectedResult: "a b c" },
		{ data: []string{ "a", "b", "c" }, operator: ", ", expectedResult: "a, b, c" },
		{ data: []string{ "a", "b", "c" }, operator: " , ", expectedResult: "a , b , c" },
		{ data: []string{ "a", "bbb", "c" }, operator: " , ", expectedResult: "a , bbb , c" },
		{ data: []string{ "a", "bbb", "cqwe" }, operator: " , ", expectedResult: "a , bbb , cqwe" },
		{ data: []string{ "h", "e", "w" }, operator: "j", expectedResult: "hjejw" },
	}

	for testId, test := range tests {
		test.performTest(testId, t)
	}
}