package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/repository"
	"notification-service/internal/util/test"
	"testing"
)

func TestBasicTemplateV1Controller_HandleById(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := NewTemplateV1Controller(templateRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqMethod string, reqURLVariable string, reqBody *string, reqHeaders map[string]string, expectedStatusCode int) test.ControllerTestCase {
		return test.ControllerTestCase{
			Router:          router,
			ReqMethod:       reqMethod,
			ReqURL:          "/v1/templates/" + reqURLVariable,
			ReqHeaders:      reqHeaders,
			ReqBody:         reqBody,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	s := func(str string) *string { return &str }

	testCases := []test.ControllerTestCase{
		newTestCase(http.MethodGet, "", nil, nil, http.StatusUnauthorized),

		newTestCase(http.MethodGet, "1", nil, nil, http.StatusUnauthorized),
		newTestCase(http.MethodGet, "1", nil, map[string]string{ "Authorization": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase(http.MethodGet, "1", nil, map[string]string{ "Authorization": "Bearer 13124" }, http.StatusOK),
		newTestCase(http.MethodGet, "a", nil, map[string]string{ "Authorization": "Bearer 13124" }, http.StatusBadRequest),
		newTestCase(http.MethodGet, "1a", nil, map[string]string{ "Authorization": "Bearer 13124" }, http.StatusBadRequest),

		newTestCase(http.MethodDelete, "1", nil, map[string]string{ "Authorization": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase(http.MethodDelete, "1", nil, map[string]string{ "Authorization": "Bearer 13124" }, http.StatusOK),

		newTestCase(http.MethodPut, "1", nil, map[string]string{ "Authorization": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase(
			http.MethodPut,
			"1",
			s("{ \"id\": 1, \"body\": { \"email\": \"Hello, @{secondName}\", \"sms\": \"Hello, @{firstName}\", \"push\": \"Hi, @{username}\" }, \"language\": \"EN\", \"type\": \"test2\" }"),
			map[string]string{
				"Authorization": "Bearer 13124",
				"Content-Type": "application/json",
			},
			http.StatusOK,
		),
	}

	test.RunControllerTestCases(&testCases, t)
}
