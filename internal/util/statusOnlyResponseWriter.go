package util

import (
	"notification-service/internal/util/iface"
)

type StatusOnlyResponseWriter struct {
	iface.IResponseWriter
	StatusCode *int
}

func (sorw *StatusOnlyResponseWriter) Status(statusCode int) iface.IResponseWriter {
	sorw.StatusCode = &statusCode
	return sorw
}

func (sorw *StatusOnlyResponseWriter) Text(text string) iface.IResponseWriter      { return sorw }
func (sorw *StatusOnlyResponseWriter) Error(err error) iface.IResponseWriter       { return sorw }
func (sorw *StatusOnlyResponseWriter) TextError(err string) iface.IResponseWriter  { return sorw }
func (sorw *StatusOnlyResponseWriter) Json(data interface{}) iface.IResponseWriter { return sorw }
func (sorw *StatusOnlyResponseWriter) Bytes(data []byte) iface.IResponseWriter     { return sorw }