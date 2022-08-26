package layer

import (
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"

	"encoding/json"
	"net/http"
)

func JSONBytesConverterMiddleware(bytes []byte, out iface.IRequest) bool {
	err := json.Unmarshal(bytes, &out)
	return !helper.IsError(err)
}

func JSONConverterMiddleware(res iface.IResponseWriter, req *http.Request, out iface.IRequest) bool {
	if req.Header.Get("Content-Type") != "application/json" {
		res.Status(http.StatusUnsupportedMediaType)
		return false
	}

	if req.Body == nil {
		res.Status(http.StatusBadRequest).TextError("Empty body")
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
