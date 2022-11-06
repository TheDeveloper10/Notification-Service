package test

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/httpctrl"
	"notification-service/internal/repository"
	"testing"
)

func TestBasicNotificationV1Controller_Get(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	notificationRepository := repository.NewMockNotificationRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := httpctrl.NewNotificationV1Controller(templateRepository, notificationRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqHeaders map[string]string, expectedStatusCode int) ControllerTestCase {
		return ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodGet,
			ReqURL:          "/v1/notifications",
			ReqHeaders:      reqHeaders,
			ReqBody:         nil,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	testCases := []ControllerTestCase{
		newTestCase(nil, http.StatusUnauthorized),
		newTestCase(map[string]string{ "Authorization": "Basic 13124" }, http.StatusUnauthorized),
		newTestCase(map[string]string{ "Authorization": "Bearer aaa" }, http.StatusUnauthorized),
		newTestCase(map[string]string{ "Authorization": "Bearer bbb" }, http.StatusUnauthorized),
		newTestCase(map[string]string{ "Authorization": "Bearer ccc" }, http.StatusForbidden),
		newTestCase(map[string]string{ "Authorization": "Bearer 13124" }, http.StatusOK),
	}

	RunControllerTestCases(&testCases, t)
}