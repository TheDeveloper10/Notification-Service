package util

import (
	"encoding/json"
	"net/http"
)

type ResponseWriterWrapper struct {
	IResponseWriter
	rw *http.ResponseWriter
}

func (brw *ResponseWriterWrapper) Status(statusCode int) (IResponseWriter) {
	(*brw.rw).WriteHeader(statusCode)
	return brw
} 

func (brw *ResponseWriterWrapper) Text(text string) (IResponseWriter) {
	(*brw.rw).Write([]byte(text))
	return brw
}

func (brw *ResponseWriterWrapper) Json(data interface{}) (IResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return brw.Bytes(bytes)
}

func (brw *ResponseWriterWrapper) Bytes(data []byte) (IResponseWriter) {
	(*brw.rw).Write(data)
	return brw
}

func ConvertResponseWriter(res *http.ResponseWriter) (IResponseWriter) {
	return &ResponseWriterWrapper {
		rw: res,
	}
}