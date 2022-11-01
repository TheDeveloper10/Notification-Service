package test

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/controller/layer"
	"notification-service/internal/util/iface"
	"reflect"
	"strconv"
	"testing"
)

type testData struct {
	Id int `json:"id_2"`
	Text string `json:"TEXT"`
	Arr1 []string `json:"ar1r"`
	Arr2 []int `json:"arr2"`
}

func (td *testData) Validate() iface.IErrorList {
	return nil
}

func TestJSONConverterMiddleware(t *testing.T) {
	testCases := []LayerTestCase{
		{ ExpectedStatus: http.StatusUnsupportedMediaType, SetHeader: false, Body: nil },
		{ ExpectedStatus: http.StatusBadRequest, SetHeader: true, Body: nil },
		{ ExpectedStatus: http.StatusUnsupportedMediaType, SetHeader: false, Body: &testData{Text: "H123I"} },


		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{} },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Id: 1} },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Text: "H123I"} },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr1: []string{} } },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr1: []string{""} } },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr1: []string{"t1e2s3t"} } },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr1: []string{"t1e2s3t", "q2f3w4e5q"} } },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr2: []int{} } },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr2: []int{1} } },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr2: []int{45, 54} } },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr2: []int{45, 54, 66} } },
		{ ExpectedStatus: http.StatusOK, SetHeader: true, Body: &testData{Arr1: []string{"t1e2s3t", "q2f3w4e5q"}, Arr2: []int{45, 54, 66} } },
	}

	for testId, testCase := range testCases {
		performJSONConverterMiddlewareTest(t, testId, testCase)
	}
}

func performJSONConverterMiddlewareTest(t *testing.T, testId int, testCase LayerTestCase) {
	req, res := testCase.PrepareTest(t)

	before := ""
	if testCase.Body != nil {
		beforeBytes, _ := json.Marshal(testCase.Body)
		before = string(beforeBytes)
	}

	layer.JSONConverterMiddleware(res, req, testCase.Body)
	statusCode := reflect.ValueOf(res).Elem().FieldByName("statusCode").Int()

	if testCase.ExpectedStatus != int(statusCode) {
		t.Error(strconv.Itoa(testId) + ": Status Code of Response is " + strconv.Itoa(int(statusCode)) + " and not " + strconv.Itoa(testCase.ExpectedStatus))
		return
	}

	after := ""
	if testCase.Body != nil {
		afterBytes, _ := json.Marshal(testCase.Body)
		after = string(afterBytes)
	}

	if after != before {
		t.Error(strconv.Itoa(testId) + ":   Before: " + before + "   After: " + after)
	}
}