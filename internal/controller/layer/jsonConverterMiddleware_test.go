package layer

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/util/iface"
	"notification-service/internal/util/test"
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
	testCases := []test.LayerTestCase{
		{ http.StatusUnsupportedMediaType, false, nil },
		{ http.StatusBadRequest, true, nil },
		{ http.StatusUnsupportedMediaType, false, &testData{Text: "H123I"} },


		{ http.StatusOK, true, &testData{} },
		{ http.StatusOK, true, &testData{Id: 1} },
		{ http.StatusOK, true, &testData{Text: "H123I"} },
		{ http.StatusOK, true, &testData{Arr1: []string{} } },
		{ http.StatusOK, true, &testData{Arr1: []string{""} } },
		{ http.StatusOK, true, &testData{Arr1: []string{"t1e2s3t"} } },
		{ http.StatusOK, true, &testData{Arr1: []string{"t1e2s3t", "q2f3w4e5q"} } },
		{ http.StatusOK, true, &testData{Arr2: []int{} } },
		{ http.StatusOK, true, &testData{Arr2: []int{1} } },
		{ http.StatusOK, true, &testData{Arr2: []int{45, 54} } },
		{ http.StatusOK, true, &testData{Arr2: []int{45, 54, 66} } },
		{ http.StatusOK, true, &testData{Arr1: []string{"t1e2s3t", "q2f3w4e5q"}, Arr2: []int{45, 54, 66} } },
	}

	for testId, testCase := range testCases {
		performJSONConverterMiddlewareTest(t, testId, testCase)
	}

	//id := 0
	//name := "test"
	//arr := []string{"123", "234"}
	//
	//PerformJsonConverterMiddlewareTest(0, t, nil, nil)
	//PerformJsonConverterMiddlewareTest(1, t, nil, JsonTestStruct{})
	//PerformJsonConverterMiddlewareTest(2, t, nil, JsonTestStruct{Id: &id})
	//PerformJsonConverterMiddlewareTest(3, t, nil, JsonTestStruct{Name: &name})
	//PerformJsonConverterMiddlewareTest(4, t, nil, JsonTestStruct{Arr: arr})
	//PerformJsonConverterMiddlewareTest(5, t, nil, JsonTestStruct{Id: &id, Name: &name, Arr: arr})
}

func performJSONConverterMiddlewareTest(t *testing.T, testId int, testCase test.LayerTestCase) {
	req, res := testCase.PrepareTest(t)

	before := ""
	if testCase.Body != nil {
		beforeBytes, _ := json.Marshal(testCase.Body)
		before = string(beforeBytes)
	}

	JSONConverterMiddleware(res, req, testCase.Body)
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