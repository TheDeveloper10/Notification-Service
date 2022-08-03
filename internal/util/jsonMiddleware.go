package util

import (
	"encoding/json"
	"net/http"

	"notification-service.com/packages/internal/dto"
)

func JsonMiddleware(res IResponseWriter, req *http.Request, out dto.AbstractRequest) bool {
	if req.Header.Get("Content-Type") != "application/json" {
		res.Status(http.StatusUnsupportedMediaType)
		return false
	}

	err := json.NewDecoder(req.Body).Decode(&out)
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
