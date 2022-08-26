package test

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"notification-service/internal/helper"
	"strconv"
	"testing"
)

func ControllerTest(
		testId int,
		t *testing.T,
		body io.Reader,
		headers map[string]string,
		statusCode int,
		method string,
		ctrl func(http.ResponseWriter, *http.Request),
		url string,
		urlVars map[string]string,
	) {
	req, err := http.NewRequest(method, url, body)
	if helper.IsError(err) {
		t.Fatal(err.Error())
	}

	if urlVars != nil {
		req = mux.SetURLVars(req, urlVars)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	rec := httptest.NewRecorder()

	ctrl(rec, req)

	res := rec.Result()

	if res.StatusCode != statusCode {
		t.Error(strconv.Itoa(testId) + ": Status Code of Response is " + strconv.Itoa(res.StatusCode) + " and not " + strconv.Itoa(statusCode))
	}
}