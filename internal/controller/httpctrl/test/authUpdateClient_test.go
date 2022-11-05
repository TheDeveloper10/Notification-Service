package test

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/httpctrl"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"testing"
)

func TestBasicAuthV1Controller_UpdateClient(t *testing.T) {
	// TODO: fix this path*
	util.LoadConfig("../../../" + util.ServiceConfigPath)

	clientRepository := repository.NewMockClientRepository()
	bac := httpctrl.NewAuthV1Controller(clientRepository)
	router := rem.NewRouter()
	bac.CreateRoutes(router)

	newTestCase := func(id string, reqBody *string, reqHeaders map[string]string, expectedStatusCode int) ControllerTestCase {
		return ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodPut,
			ReqURL:          "/v1/oauth/client/" + id,
			ReqHeaders:      reqHeaders,
			ReqBody:         reqBody,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	s := func(str string) *string { return &str }

	testCases := []ControllerTestCase{
		newTestCase("aa", nil, nil, http.StatusUnauthorized),
		newTestCase("aa", nil, map[string]string{ "Authorization": "Basic test:13124" }, http.StatusUnauthorized),
		newTestCase("aa", nil, map[string]string{ "Authorization": "Bearer 1234" }, http.StatusForbidden),
		newTestCase(
			"aa",
			s("{}"),
			map[string]string{
				"Authorization": "Bearer " + util.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusOK,
		),
		newTestCase(
			"aa",
			s("{ \"permissions\": [ \"read_templates\" ] }"),
			map[string]string{
				"Authorization": "Bearer " + util.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusOK,
		),
		newTestCase(
			"bb",
			s("{ \"permissions\": [ \"read_templates\" ] }"),
			map[string]string{
				"Authorization": "Bearer " + util.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusNotFound,
		),
		newTestCase(
			"cc",
			s("{ \"permissions\": [ \"read_templates\" ] }"),
			map[string]string{
				"Authorization": "Bearer " + util.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusBadRequest,
		),
	}

	RunControllerTestCases(&testCases, t)
}
