package controller

import (
	"io"
	"net/http"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"notification-service/internal/util/test"
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
	clientRepository := repository.NewClientRepository(true)
	bac := NewAuthV1Controller(clientRepository)

	test.ControllerTest(testId, t, body, headers, statusCode, "POST", bac.HandleClient, "")
}