package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/repository"
	"notification-service/internal/util/test"
	"testing"
)

func TestBasicTemplateV1Controller_GetBulk(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := NewTemplateV1Controller(templateRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqURL string, reqHeaders map[string]string, expectedStatusCode int) test.ControllerTestCase {
		reqURL = "/v1/templates" + reqURL
		return test.ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodGet,
			ReqURL:          reqURL,
			ReqHeaders:      reqHeaders,
			ReqBody:         nil,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	testCases := []test.ControllerTestCase{
		newTestCase("", nil, http.StatusUnauthorized),
		newTestCase("", map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase("", map[string]string{ "Authentication": "Bearer 13124" }, http.StatusOK),
		newTestCase("?size=10&page=2", map[string]string{ "Authentication": "Bearer 13124" }, http.StatusOK),
	}

	test.RunControllerTestCases(&testCases, t)
}
