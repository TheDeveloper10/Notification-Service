package util

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/helper"
)

type responseWriterWrapper struct {
	IResponseWriter
	rw *http.ResponseWriter
}

func (brw *responseWriterWrapper) Status(statusCode int) IResponseWriter {
	(*brw.rw).WriteHeader(statusCode)
	return brw
} 

func (brw *responseWriterWrapper) Text(text string) IResponseWriter {
	_, err := (*brw.rw).Write([]byte(text))
	if helper.IsError(err) {
		return brw
	}
	return brw
}

func (brw *responseWriterWrapper) Json(data interface{}) IResponseWriter {
	bytes, err := json.Marshal(data)
	if helper.IsError(err) {
		return brw
	}
	return brw.Bytes(bytes)
}

func (brw *responseWriterWrapper) Bytes(data []byte) IResponseWriter {
	_, err := (*brw.rw).Write(data)
	if helper.IsError(err) {
		return brw
	}
	return brw
}

func WrapResponseWriter(res *http.ResponseWriter) IResponseWriter {
	return &responseWriterWrapper {
		rw: res,
	}
}