package test

import (
	"github.com/TheDeveloper10/rem"
	"io"
	"net/http"
	"net/http/httptest"
	"notification-service/internal/util"
	"strconv"
	"strings"
	"testing"
)

// Controller Test Case

type ControllerTestCase struct {
	Router *rem.Router

	ReqMethod 		string
	ReqURL 			string
	ReqHeaders 		map[string]string
	ReqBody 		*string

	ExpectedStatus int
}

func (ctc *ControllerTestCase) RunTest(testId int, t *testing.T) {
	var body io.Reader = nil
	if ctc.ReqBody != nil {
		body = strings.NewReader(*ctc.ReqBody)
	}

	req, err := http.NewRequest(ctc.ReqMethod, ctc.ReqURL, body)
	if util.ManageError(err) {
		t.Fatal(err.Error())
	}

	if ctc.ReqHeaders != nil {
		for k, v := range ctc.ReqHeaders {
			req.Header.Add(k, v)
		}
	}

	rec := httptest.NewRecorder()

	ctc.Router.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != ctc.ExpectedStatus {
		t.Error(strconv.Itoa(testId) + ": Status Code of Response is " + strconv.Itoa(res.StatusCode) + " and not " + strconv.Itoa(ctc.ExpectedStatus))
	}
}

func RunControllerTestCases(cases *[]ControllerTestCase, t *testing.T) {
	for testId, testCase := range *cases {
		testCase.RunTest(testId, t)
	}
}
