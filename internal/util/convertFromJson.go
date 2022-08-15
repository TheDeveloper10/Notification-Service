package util

import (
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"

	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ConvertFromJson(res iface.IResponseWriter, req *http.Request, out iface.IRequest) bool {
	if req.Header.Get("Content-Type") != "application/json" {
		log.Error("Unsupported Content-Type")
		res.Status(http.StatusUnsupportedMediaType)
		return false
	}

	err := json.NewDecoder(req.Body).Decode(&out)
	if helper.IsError(err) {
		res.Status(http.StatusBadRequest).TextError("Invalid JSON")
		return false
	}

	errs := out.Validate()
	if errs.ErrorsCount() > 0 {
		res.Status(http.StatusBadRequest).Json(errs)
		return false
	}

	return true
}
