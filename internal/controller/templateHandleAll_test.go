package controller

import (
	"io"
	"net/http"
	"notification-service/internal/repository"
	"notification-service/internal/util/test"
	"strings"
	"testing"
)

func TestBasicTemplateV1Controller_HandleAll(t *testing.T) {
	createTemplateTest(0, t, nil, nil, http.StatusUnauthorized)
	createTemplateTest(1, t, nil, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized)
	createTemplateTest(
		2,
		t,
		strings.NewReader("{ \"contactType\": \"email\", \"template\": \"Hi @{firstName}\", \"language\": \"EN\", \"Type\": \"test\" }"),
		map[string]string{
			"Authentication": "Bearer 1234",
			"Content-Type": "application/json",
		},
		http.StatusCreated,
	)

	getBulkTemplatesTest(3, t, nil, http.StatusUnauthorized, "")
	getBulkTemplatesTest(4, t, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized, "")
	getBulkTemplatesTest(5, t, map[string]string{ "Authentication": "Bearer 13124" }, http.StatusOK, "")
	getBulkTemplatesTest(6, t, map[string]string{ "Authentication": "Bearer 13124" }, http.StatusOK, "/v1/templates?size=10&page=2")
}

func createTemplateTest(testId int, t *testing.T, body io.Reader, headers map[string]string, statusCode int) {
	templateRepository := repository.NewTemplateRepository(true)
	clientRepository := repository.NewClientRepository(true)
	tac := NewTemplateV1Controller(templateRepository, clientRepository)

	test.ControllerTest(testId, t, body, headers, statusCode, "POST", tac.HandleAll, "")
}

func getBulkTemplatesTest(testId int, t *testing.T, headers map[string]string, statusCode int, url string) {
	templateRepository := repository.NewTemplateRepository(true)
	clientRepository := repository.NewClientRepository(true)
	tac := NewTemplateV1Controller(templateRepository, clientRepository)

	test.ControllerTest(testId, t, nil, headers, statusCode, "GET", tac.HandleAll, url)
}