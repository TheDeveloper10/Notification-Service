package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"strconv"
	"strings"
	"testing"
)

func TestBasicAuthV1Controller_CreateClient(t *testing.T) {
	// TODO: fix this path*
	helper.LoadConfig("../../" + helper.ServiceConfigPath)

	createClientTest(0, t, nil, nil, http.StatusUnauthorized)
	createClientTest(1, t, nil, map[string]string{ "Authentication": "Basic test:13124" }, http.StatusUnauthorized)
	createClientTest(2, t, nil, map[string]string{ "Authentication": "Bearer 1234" }, http.StatusForbidden)

	createClientTest(
		3,
		t,
		strings.NewReader("{ }"),
		map[string]string{
			"Authentication": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
			"Content-Type": "application/json",
		},
		http.StatusBadRequest,
	)

	createClientTest(
		4,
		t,
		strings.NewReader("{ \"permissions\": [ \"read_templates\" ] }"),
		map[string]string{
			"Authentication": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
			"Content-Type": "application/json",
		},
		http.StatusCreated,
	)
}

func createClientTest(testId int, t *testing.T, body io.Reader, headers map[string]string, statusCode int) {
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

	bac.HandleClient(rec, req)

	res := rec.Result()

	if res.StatusCode != statusCode {
		t.Error(strconv.Itoa(testId) + ": Status Code of Response is " + strconv.Itoa(res.StatusCode) + " and not " + strconv.Itoa(statusCode))
	}
}