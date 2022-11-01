package test

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller"
	"notification-service/internal/repository"
	"testing"
)

func TestBasicTemplateV1Controller_GetBulk(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := controller.NewTemplateV1Controller(templateRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqURL string, reqHeaders map[string]string, expectedStatusCode int) ControllerTestCase {
		reqURL = "/v1/templates" + reqURL
		return ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodGet,
			ReqURL:          reqURL,
			ReqHeaders:      reqHeaders,
			ReqBody:         nil,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	testCases := []ControllerTestCase{
		newTestCase("", nil, http.StatusUnauthorized),
		newTestCase("", map[string]string{ "Authorization": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase("", map[string]string{ "Authorization": "Bearer 13124" }, http.StatusOK),
		newTestCase("?size=10&page=2", map[string]string{ "Authorization": "Bearer 13124" }, http.StatusOK),
	}

	RunControllerTestCases(&testCases, t)
}
