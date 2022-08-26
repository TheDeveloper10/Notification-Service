package controller

import (
	"io"
	"net/http"
	"notification-service/internal/repository"
	"notification-service/internal/util/test"
	"strings"
	"testing"
)

func TestBasicNotificationV1Controller_HandleAll(t *testing.T) {
	notificationsTest(0, t, nil, nil, http.StatusUnauthorized, "GET")
	notificationsTest(1, t, nil, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized, "GET")

	notificationsTest(2, t, nil, map[string]string{ "Authentication": "Bearer 13124" }, http.StatusOK, "GET")

	notificationsTest(
		3,
		t,
		strings.NewReader("{ " +
			"\"appId\": \"test\", " +
			"\"templateId\": 4, " +
			"\"contactType\": \"email\"," +
			"\"Title\": \"Welcome\"," +
			"\"targets\": [ {" +
			" \"email\": \"test@example.com\"," +
			" \"placeholders\": [ { \"key\": \"firstName\", \"val\": \"John\" } ]" +
			"} ] }"),
		map[string]string{
			"Authentication": "Bearer 13124",
			"Content-Type": "application/json",
		},
		http.StatusCreated,
		"POST",
		)
}

func notificationsTest(testId int, t *testing.T, body io.Reader, headers map[string]string, statusCode int, method string) {
	templateRepository := repository.NewTemplateRepository(true)
	notificationRepository := repository.NewNotificationRepository(true)
	clientRepository := repository.NewClientRepository(true)
	tac := NewNotificationV1Controller(templateRepository, notificationRepository, clientRepository)

	test.ControllerTest(testId, t, body, headers, statusCode, method, tac.HandleAll, "", nil)
}