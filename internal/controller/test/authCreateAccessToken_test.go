package test

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"testing"
)

func TestBasicAuthV1Controller_CreateAccessToken(t *testing.T) {
	helper.LoadConfig("../../../" + helper.ServiceConfigPath)

	clientRepository := repository.NewMockClientRepository()
	bac := controller.NewAuthV1Controller(clientRepository)
	router := rem.NewRouter()
	bac.CreateRoutes(router)

	newTestCase := func(reqHeaders map[string]string, expectedStatusCode int) ControllerTestCase {
		return ControllerTestCase{
			Router: router,
			ReqMethod: http.MethodPost,
			ReqURL: "/v1/oauth/token",
			ReqHeaders: reqHeaders,
			ReqBody: nil,
			ExpectedStatus: expectedStatusCode,
		}
	}

	testCases := []ControllerTestCase{
		newTestCase(nil, http.StatusUnauthorized),
		newTestCase(map[string]string{ "Authorization": "Bearer 13124" }, http.StatusUnauthorized),
		newTestCase(map[string]string{ "Authorization": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase(map[string]string{ "Authorization": "Basic aWQ6c2VjcmV0" }, http.StatusOK),
	}

	RunControllerTestCases(&testCases, t)
}
