package util

import (
	"encoding/json"
	"net/http"
)

type responseWriterWrapper struct {
	IResponseWriter
	rw *http.ResponseWriter
}

func (brw *responseWriterWrapper) Status(statusCode int) (IResponseWriter) {
	(*brw.rw).WriteHeader(statusCode)
	return brw
} 

func (brw *responseWriterWrapper) Text(text string) (IResponseWriter) {
	(*brw.rw).Write([]byte(text))
	return brw
}

func (brw *responseWriterWrapper) Json(data interface{}) (IResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return brw.Bytes(bytes)
}

func (brw *responseWriterWrapper) Bytes(data []byte) (IResponseWriter) {
	(*brw.rw).Write(data)
	return brw
}

func WrapResponseWriter(res *http.ResponseWriter) (IResponseWriter) {
	return &responseWriterWrapper {
		rw: res,
	}
}