package util

import (
	"errors"
	"notification-service/internal/util/iface"
)

func ValidateRequestAndCombineErrors(req iface.IRequest) error {
	errs := req.Validate()
	if len(errs) > 0 {
		errMessage := ""
		for _, v := range errs {
			errMessage += v.Error() + "; "
		}

		return errors.New(errMessage)
	}
	return nil
}