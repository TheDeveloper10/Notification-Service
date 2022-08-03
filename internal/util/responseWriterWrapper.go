package util

import (
	"encoding/json"
	"net/http"
)

type ResponseWriterWrapper struct {
	IResponseWriter
	RW *http.ResponseWriter
}

func (brw *ResponseWriterWrapper) Status(statusCode int) (IResponseWriter) {
	(*brw.RW).WriteHeader(statusCode)
	return brw
} 

func (brw *ResponseWriterWrapper) Text(text string) (IResponseWriter) {
	(*brw.RW).Write([]byte(text))
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
	(*brw.RW).Write(data)
	return brw
}