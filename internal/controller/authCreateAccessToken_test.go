package controller

import (
	"io"
	"net/http"
	"notification-service/internal/helper"
	"testing"
)

func TestBasicAuthV1Controller_HandleToken(t *testing.T) {
	helper.LoadConfig("../../" + helper.ServiceConfigPath)

	createAccessTokenTest(0, t, nil, nil, http.StatusUnauthorized)
	createAccessTokenTest(1, t, nil, map[string]string{ "Authentication": "Bearer 13124" }, http.StatusUnauthorized)
	createAccessTokenTest(2, t, nil, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized)
	//                                                                                 base64(id:secret)
	createAccessTokenTest(3, t, nil, map[string]string{ "Authentication": "Basic aWQ6c2VjcmV0" }, http.StatusOK)
}

func createAccessTokenTest(testId int, t *testing.T, body io.Reader, headers map[string]string, statusCode int) {
	//clientRepository := repository.NewMockClientRepository()
	//bac := NewAuthV1Controller(clientRepository)

	//test.ControllerTest(testId, t, body, headers, statusCode, "POST", bac.HandleToken, "", nil)
}