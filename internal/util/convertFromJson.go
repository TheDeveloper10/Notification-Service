package util

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"notification-service.com/packages/internal/dto"
)

func ConvertFromJson(res IResponseWriter, req *http.Request, out dto.AbstractRequest) bool {
	if req.Header.Get("Content-Type") != "application/json" {
		log.Error("Unsupported Content-Type")
		res.Status(http.StatusUnsupportedMediaType)
		return false
	}

	err := json.NewDecoder(req.Body).Decode(&out)
	if err != nil {
		log.Error(err.Error())
		res.Status(http.StatusBadRequest)
		return false
	}

	errors := out.Validate()
	if len(errors) > 0 {
		errorMessage := ""
		for _, v := range errors {
			errorMessage += v.Error() + "\n"
		}
		log.Error(errorMessage)
		res.Status(http.StatusBadRequest).Text(errorMessage)
		return false
	}

	return true
}
