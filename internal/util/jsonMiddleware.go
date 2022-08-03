package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"notification-service.com/packages/internal/dto"
)

func JsonMiddleware(res IResponseWriter, req *http.Request, out dto.AbstractRequest) bool {
	if req.Header.Get("Content-Type") != "application/json" {
		res.Status(http.StatusUnsupportedMediaType)
		return false
	}

	bodyBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.Status(http.StatusInternalServerError)
		return false
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		res.Status(http.StatusBadRequest)
		return false
	}

	status, message := out.Validate()
	if !status {
		res.Status(http.StatusBadRequest).Text(message)
		return false
	}

	return true
}
