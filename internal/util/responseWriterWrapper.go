package util

import (
	"encoding/json"
	"net/http"
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"
)

func WrapResponseWriter(res *http.ResponseWriter) iface.IResponseWriter {
	return &responseWriterWrapper{
		rw: res,
	}
}

type responseWriterWrapper struct {
	iface.IResponseWriter
	rw *http.ResponseWriter
}

func (rrw *responseWriterWrapper) Status(statusCode int) iface.IResponseWriter {
	(*rrw.rw).WriteHeader(statusCode)
	return rrw
}

func (rrw *responseWriterWrapper) Text(text string) iface.IResponseWriter {
	_, err := (*rrw.rw).Write([]byte(text))
	helper.IsError(err)
	return rrw
}

func (rrw *responseWriterWrapper) Error(err error) iface.IResponseWriter {
	return rrw.Json(NewErrorList().AddError(err))
}

func (rrw *responseWriterWrapper) TextError(err string) iface.IResponseWriter {
	return rrw.Json(NewErrorList().AddErrorFromString(err))
}

func (rrw *responseWriterWrapper) Json(data interface{}) iface.IResponseWriter {
	bytes, err := json.Marshal(data)
	if helper.IsError(err) {
		return rrw
	}
	return rrw.Bytes(bytes)
}

func (rrw *responseWriterWrapper) Bytes(data []byte) iface.IResponseWriter {
	_, err := (*rrw.rw).Write(data)
	helper.IsError(err)
	return rrw
}
