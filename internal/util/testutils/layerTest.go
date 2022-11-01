package testutils

import (
	"encoding/json"
	"github.com/TheDeveloper10/rem"
	"io"
	"net/http"
	"net/http/httptest"
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"
	"strings"
	"testing"
)

type LayerTestCase struct {
	ExpectedStatus int
	SetHeader      bool
	Body 		   iface.IRequest
}

func (ltc *LayerTestCase) PrepareTest(t *testing.T) (rem.IRequest, rem.IResponse) {
	var body io.Reader = nil

	if ltc.Body != nil {
		bodyBytes, _ := json.Marshal(ltc.Body)
		body = strings.NewReader(string(bodyBytes))
	}

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "", body)
	if helper.IsError(err) {
		t.Fatal(err.Error())
	}

	if ltc.SetHeader {
		req.Header.Add("Content-Type", "application/json")
	}

	response := rem.NewHTTPResponseWriter(rec)
	request := rem.NewBasicRequest(req)

	response.Status(http.StatusOK)

	return request, response
}