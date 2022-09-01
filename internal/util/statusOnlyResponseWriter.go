package util

import (
	"github.com/TheDeveloper10/rem"
)

type StatusOnlyResponseWriter struct {
	rem.IResponse
	StatusCode *int
}

func (sorw *StatusOnlyResponseWriter) Status(statusCode int) rem.IResponse {
	sorw.StatusCode = &statusCode
	return sorw
}

func (sorw *StatusOnlyResponseWriter) Header(key string, value string) rem.IResponse {
	return sorw
}

func (sorw *StatusOnlyResponseWriter) Bytes(data []byte) rem.IResponse {
	return sorw
}

func (sorw *StatusOnlyResponseWriter) Text(text string) rem.IResponse {
	return sorw
}

func (sorw *StatusOnlyResponseWriter) JSON(data interface{}) rem.IResponse {
	return sorw
}
