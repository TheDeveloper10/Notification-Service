package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"strconv"
	"testing"
)

func TestBasicAuthV1Controller_HandleToken(t *testing.T) {
	helper.LoadConfig("../../" + helper.ServiceConfigPath)

	createAccessTokenTest(0, t, nil, nil, http.StatusUnauthorized)
	createAccessTokenTest(1, t, nil, map[string]string{ "Authentication": "Bearer 13124" }, http.StatusUnauthorized)
	createAccessTokenTest(2, t, nil, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized)
	// base64(id:secret)
	createAccessTokenTest(3, t, nil, map[string]string{ "Authentication": "Basic aWQ6c2VjcmV0" }, http.StatusOK)
}

func createAccessTokenTest(testId int, t *testing.T, body io.Reader, headers map[string]string, statusCode int) {
	req, err := http.NewRequest("POST", "localhost/v1/oauth/client", body)
	if helper.IsError(err) {
		t.Fatal(err.Error())
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	rec := httptest.NewRecorder()

	clientRepository := repository.NewClientRepository(true)
	bac := NewAuthV1Controller(clientRepository)

	bac.HandleToken(rec, req)

	res := rec.Result()

	if res.StatusCode != statusCode {
		t.Error(strconv.Itoa(testId) + ": Status Code of Response is " + strconv.Itoa(res.StatusCode) + " and not " + strconv.Itoa(statusCode))
	}
}