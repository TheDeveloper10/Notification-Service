package test

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller"
	"notification-service/internal/repository"
	"testing"
)

func TestBasicTemplateV1Controller_Create(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := controller.NewTemplateV1Controller(templateRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqBody *string, reqHeaders map[string]string, expectedStatusCode int) ControllerTestCase {
		return ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodPost,
			ReqURL:          "/v1/templates",
			ReqHeaders:      reqHeaders,
			ReqBody:         reqBody,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	s := func(str string) *string { return &str }

	testCases := []ControllerTestCase{
		newTestCase(nil, nil, http.StatusUnauthorized),
		newTestCase(nil, map[string]string{ "Authorization": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase(
			s("{ \"body\": { \"email\": \"Hi, @{firstName}\", \"sms\": \"Hi, @{lastName}\", \"push\": \"Hi, @{username}\" }, \"language\": \"EN\", \"Type\": \"test\" }"),
			map[string]string{
				"Authorization": "Bearer 1234",
				"Content-Type": "application/json",
			},
			http.StatusCreated),
	}

	RunControllerTestCases(&testCases, t)
}
