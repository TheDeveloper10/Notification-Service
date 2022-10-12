package layer

import (
	"github.com/TheDeveloper10/rem"
	"notification-service/internal/helper"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"

	"encoding/json"
	"net/http"
)

func JSONBytesConverterMiddleware(bytes []byte, out iface.IRequest) bool {
	err := json.Unmarshal(bytes, &out)
	return !helper.IsError(err)
}

func JSONConverterMiddleware(res rem.IResponse, req rem.IRequest, out iface.IRequest) bool {
	if req.GetHeaders().Get("Content-Type") != "application/json" {
		res.Status(http.StatusUnsupportedMediaType)
		return false
	}

	if req.GetBody() == nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Empty body"))
		return false
	}

	err := json.NewDecoder(req.GetBody()).Decode(&out)
	if helper.IsError(err) {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Invalid JSON"))
		return false
	}

	errs := out.Validate()
	if errs != nil && errs.ErrorsCount() > 0 {
		res.Status(http.StatusBadRequest).JSON(errs)
		return false
	}

	return true
}

func ToJSONString(in iface.IRequest) *string {
	bytes, err := json.Marshal(in)
	if err != nil {
		return nil
	}

	str := string(bytes)
	return &str
}