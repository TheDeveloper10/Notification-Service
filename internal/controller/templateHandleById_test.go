package controller

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestBasicTemplateV1Controller_HandleById(t *testing.T) {
	templateByIdTest(0, t, nil, nil, http.StatusBadRequest, false, "GET")
	templateByIdTest(1, t, nil, nil, http.StatusUnauthorized, true, "GET")

	templateByIdTest(2, t, nil, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized, true, "GET")
	templateByIdTest(3, t, nil, map[string]string{ "Authentication": "Bearer 13124" }, http.StatusOK, true, "GET")

	templateByIdTest(4, t, nil, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized, true, "DELETE")
	templateByIdTest(5, t, nil, map[string]string{ "Authentication": "Bearer 13124" }, http.StatusOK, true, "DELETE")

	templateByIdTest(6, t, nil, map[string]string{ "Authentication": "Basic 13124" }, http.StatusUnauthorized, true, "PUT")
	templateByIdTest(
		7,
		t,
		strings.NewReader("{ \"id\": 1, \"contactType\": \"email\", \"template\": \"Hello, @{secondName}\", \"language\": \"EN\", \"type\": \"test2\" }"),
		map[string]string{
			"Authentication": "Bearer 13124",
			"Content-Type": "application/json",
		},
		http.StatusOK,
		true,
		"PUT",
		)
}

func templateByIdTest(testId int, t *testing.T, body io.Reader, headers map[string]string, statusCode int, setUrl bool, method string) {
	//templateRepository := repository.NewMockTemplateRepository()
	//clientRepository := repository.NewMockClientRepository()
	//tac := NewTemplateV1Controller(templateRepository, clientRepository)
	//
	//url := ""
	//urlVars := map[string]string{}
	//if setUrl {
	//	url = "/v1/templates/"
	//	urlVars["templateId"] = "1"
	//}
	//test.ControllerTest(testId, t, body, headers, statusCode, method, tac.HandleById, url, urlVars)
}