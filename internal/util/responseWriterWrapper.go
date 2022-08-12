package util

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"
)

func WrapResponseWriter(res *http.ResponseWriter) iface.IResponseWriter {
	return &responseWriterWrapper {
		rw: res,
	}
}

type responseWriterWrapper struct {
	iface.IResponseWriter
	rw *http.ResponseWriter
}

func (brw *responseWriterWrapper) Status(statusCode int) iface.IResponseWriter {
	(*brw.rw).WriteHeader(statusCode)
	return brw
} 

func (brw *responseWriterWrapper) Text(text string) iface.IResponseWriter {
	_, err := (*brw.rw).Write([]byte(text))
	if helper.IsError(err) {
		return brw
	}
	return brw
}

func (brw *responseWriterWrapper) Error(err error) iface.IResponseWriter {
	_, err2 := (*brw.rw).Write([]byte(err.Error()))
	if helper.IsError(err2) {
		return brw
	}
	return brw
}

func (brw *responseWriterWrapper) Json(data interface{}) iface.IResponseWriter {
	bytes, err := json.Marshal(data)
	if helper.IsError(err) {
		return brw
	}
	return brw.Bytes(bytes)
}

func (brw *responseWriterWrapper) Bytes(data []byte) iface.IResponseWriter {
	_, err := (*brw.rw).Write(data)
	if helper.IsError(err) {
		return brw
	}
	return brw
}