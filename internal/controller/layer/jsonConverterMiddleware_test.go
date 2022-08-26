package layer

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"notification-service/internal/helper"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"strconv"
	"strings"
	"testing"
)

type JsonTestStruct struct {
	iface.IRequest
	Id *int `json:"id"`
	Name *string `json:"string"`
	Arr []string `json:"arr"`
}

func Validate(testStruct JsonTestStruct) iface.IErrorList {
	return nil
}

func TestJSONConverterMiddleware(t *testing.T) {
	id := 0
	name := "test"
	arr := []string{"123", "234"}

	PerformJsonConverterMiddlewareTest(0, t, nil, nil)
	PerformJsonConverterMiddlewareTest(1, t, nil, JsonTestStruct{})
	PerformJsonConverterMiddlewareTest(2, t, nil, JsonTestStruct{Id: &id})
	PerformJsonConverterMiddlewareTest(3, t, nil, JsonTestStruct{Name: &name})
	PerformJsonConverterMiddlewareTest(4, t, nil, JsonTestStruct{Arr: arr})
	PerformJsonConverterMiddlewareTest(5, t, nil, JsonTestStruct{Id: &id, Name: &name, Arr: arr})
}

func PerformJsonConverterMiddlewareTest(testId int, t *testing.T, headers map[string]string, data iface.IRequest) {
	dataText, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "", strings.NewReader(string(dataText)))
	if helper.IsError(err) {
		t.Fatal(err.Error())
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	rec := httptest.NewRecorder()
	brw := util.WrapResponseWriter(rec)

	JSONConverterMiddleware(brw, req, data)
	dataText2, _ := json.Marshal(data)

	if string(dataText2) != string(dataText) {
		t.Error(strconv.Itoa(testId))
	}
}