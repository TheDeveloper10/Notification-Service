package test

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller"
	"notification-service/internal/helper"
	"notification-service/internal/repository"
	"notification-service/internal/util/testutils"
	"testing"
)

func TestBasicAuthV1Controller_DeleteClient(t *testing.T) {
	// TODO: fix this path*
	helper.LoadConfig("../../../" + helper.ServiceConfigPath)

	clientRepository := repository.NewMockClientRepository()
	bac := controller.NewAuthV1Controller(clientRepository)
	router := rem.NewRouter()
	bac.CreateRoutes(router)

	newTestCase := func(id string, reqHeaders map[string]string, expectedStatusCode int) testutils.ControllerTestCase {
		return testutils.ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodDelete,
			ReqURL:          "/v1/oauth/client/" + id,
			ReqHeaders:      reqHeaders,
			ReqBody:         nil,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	testCases := []testutils.ControllerTestCase{
		newTestCase("a", nil, http.StatusUnauthorized),
		newTestCase("a", map[string]string{ "Authorization": "Basic testutils:13124" }, http.StatusUnauthorized),
		newTestCase("a", map[string]string{ "Authorization": "Bearer 1234" }, http.StatusForbidden),
		newTestCase(
			"a",
			map[string]string{
				"Authorization": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
			},
			http.StatusBadRequest,
		),
		newTestCase(
			"aa",
			map[string]string{
				"Authorization": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusOK,
		),
		newTestCase(
			"bb",
			map[string]string{
				"Authorization": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusNotFound,
		),
		newTestCase(
			"cc",
			map[string]string{
				"Authorization": "Bearer " + helper.Config.HTTPServer.MasterAccessToken,
				"Content-Type": "application/json",
			},
			http.StatusBadRequest,
		),
	}

	testutils.RunControllerTestCases(&testCases, t)
}
