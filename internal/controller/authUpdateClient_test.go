package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"notification-service/internal/util/test"
	"testing"
)

func TestBasicAuthV1Controller_UpdateClient(t *testing.T) {
	// TODO: fix this path*
	helper.LoadConfig("../../" + helper.ServiceConfigPath)

	clientRepository := repository.NewMockClientRepository()
	bac := NewAuthV1Controller(clientRepository)
	router := rem.NewRouter()
	bac.CreateRoutes(router)

	newTestCase := func(id string, reqBody *string, reqHeaders map[string]string, expectedStatusCode int) test.ControllerTestCase {
		return test.ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodPut,
			ReqURL:          "/v1/oauth/client/" + id,
			ReqHeaders:      reqHeaders,
			ReqBody:         reqBody,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	s := func(str string) *string { return &str }

	testCases := []test.ControllerTestCase{
		newTestCase("aa", nil, nil, http.StatusUnauthorized),
		newTestCase("aa", nil, map[string]string{ "Authorization": "Basic test:13124" }, http.StatusUnauthorized),
		newTestCase("aa", nil, map[string]string{ "Authorization": "Bearer 1234" }, http.StatusForbidden),
		newTestCase(
			"aa",
			s("{}"),
			map[string]string{
				"Authorization": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusOK,
		),
		newTestCase(
			"aa",
			s("{ \"permissions\": [ \"read_templates\" ] }"),
			map[string]string{
				"Authorization": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusOK,
		),
		newTestCase(
			"bb",
			s("{ \"permissions\": [ \"read_templates\" ] }"),
			map[string]string{
				"Authorization": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusNotFound,
		),
		newTestCase(
			"cc",
			s("{ \"permissions\": [ \"read_templates\" ] }"),
			map[string]string{
				"Authorization": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusBadRequest,
		),
	}

	test.RunControllerTestCases(&testCases, t)
}
