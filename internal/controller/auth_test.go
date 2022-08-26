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

func TestBasicAuthV1Controller_CreateClient(t *testing.T) {
	createClientTest(t, nil, http.StatusUnauthorized)
}

func createClientTest(t *testing.T, body io.Reader, statusCode int) {
	req, err := http.NewRequest("POST", "localhost/v1/oauth/client", body)
	if helper.IsError(err) {
		t.Fatal(err.Error())
	}

	rec := httptest.NewRecorder()

	clientRepository := repository.NewClientRepository(true)
	bac := NewAuthV1Controller(clientRepository)

	bac.HandleClient(rec, req)

	res := rec.Result()

	if res.StatusCode != statusCode {
		t.Error("Status Code of Response is " + strconv.Itoa(res.StatusCode) + " and not " + strconv.Itoa(statusCode))
	}
}