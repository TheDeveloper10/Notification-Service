package util

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/helper"

	log "github.com/sirupsen/logrus"

	"notification-service/internal/dto"
)

func ConvertFromJson(res IResponseWriter, req *http.Request, out dto.AbstractRequest) bool {
	if req.Header.Get("Content-Type") != "application/json" {
		log.Error("Unsupported Content-Type")
		res.Status(http.StatusUnsupportedMediaType)
		return false
	}

	err := json.NewDecoder(req.Body).Decode(&out)
	if helper.IsError(err) {
		res.Status(http.StatusBadRequest).Text("Invalid JSON")
		return false
	}

	err = ValidateRequestAndCombineErrors(out)
	if helper.IsError(err) {
		res.Status(http.StatusBadRequest).Text(err.Error())
		return false
	}

	return true
}
