package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/repository"
	"notification-service/internal/util/test"
	"testing"
)

func TestBasicNotificationV1Controller_Create(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	notificationRepository := repository.NewMockNotificationRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := NewNotificationV1Controller(templateRepository, notificationRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqBody *string, reqHeaders map[string]string, expectedStatusCode int) test.ControllerTestCase {
		return test.ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodPost,
			ReqURL:          "/v1/notifications",
			ReqHeaders:      reqHeaders,
			ReqBody:         reqBody,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	s := func(str string) *string { return &str }

	testCases := []test.ControllerTestCase{
		newTestCase(s("{ " +
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
			http.StatusCreated),
	}

	test.RunControllerTestCases(&testCases, t)
}