package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/repository"
	"notification-service/internal/util/test"
	"testing"
)

func TestBasicTemplateV1Controller_Create(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := NewTemplateV1Controller(templateRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqBody *string, reqHeaders map[string]string, expectedStatusCode int) test.ControllerTestCase {
		return test.ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodPost,
			ReqURL:          "/v1/templates",
			ReqHeaders:      reqHeaders,
			ReqBody:         reqBody,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	s := func(str string) *string { return &str }

	testCases := []test.ControllerTestCase{
		newTestCase(nil, nil, http.StatusUnauthorized),
		newTestCase(nil, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase(
			s("{ \"contactType\": \"email\", \"template\": \"Hi @{firstName}\", \"language\": \"EN\", \"Type\": \"test\" }"),
			map[string]string{
				"Authentication": "Bearer 1234",
				"Content-Type": "application/json",
			},
			http.StatusCreated),
	}

	test.RunControllerTestCases(&testCases, t)
}
