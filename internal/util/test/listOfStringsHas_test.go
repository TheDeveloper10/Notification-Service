package test

import (
	"notification-service/internal/util"
	"strconv"
	"testing"
)

type listOfStringsHasTest struct {
	data    util.ListOfStrings
	element string
	expectedResult bool
}

func (losht *listOfStringsHasTest) performTest(testId int, t *testing.T) {
	expected := "yes"
	if !losht.expectedResult {
		expected = "no"
	}

	result := losht.data.Has(losht.element)

	actual := "yes"
	if !result {
		actual = "no"
	}

	if losht.expectedResult != result {
		t.Error(strconv.Itoa(testId) + "\t\t Expected: " + expected + "\t\tActual: " + actual)
	}
}

func TestListOfStrings_Has(t *testing.T) {
	tests := []listOfStringsHasTest{
		{ data: []string { "a", "b", "c" }, element: "a", expectedResult: true },
		{ data: []string { "a", "b", "c" }, element: "b", expectedResult: true },
		{ data: []string { "a", "b", "c" }, element: "c", expectedResult: true },
		{ data: []string { "a", "b", "c" }, element: "f", expectedResult: false },
		{ data: []string { "aaaaa", "b", "c" }, element: "a", expectedResult: false },
		{ data: []string { "aaaaa", "b", "c" }, element: "aaa", expectedResult: false },
		{ data: []string { "aaaaa", "b", "c" }, element: "aaaaa", expectedResult: true },
		{ data: []string { "aaaaa", "b", "c" }, element: "b", expectedResult: true },
		{ data: []string { "aaaaa", "hello world!", "c" }, element: "hello world", expectedResult: false },
		{ data: []string { "aaaaa", "hello world!", "c" }, element: "hello world!", expectedResult: true },
		{ data: []string { "aaaaa", "hello world!", "c" }, element: "c", expectedResult: true },
		{ data: []string { "aaaaa", "hello world!", "c123" }, element: "c", expectedResult: false },
		{ data: []string { "aaaaa", "hello world!", "c123" }, element: "123c", expectedResult: false },
		{ data: []string { "aaaaa", "hello world!", "c123" }, element: "c123", expectedResult: true },
	}

	for testId, test := range tests {
		test.performTest(testId, t)
	}
}