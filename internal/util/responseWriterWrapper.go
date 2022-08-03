package util

import (
	"encoding/json"
	"net/http"
)

type ResponseWriterWrapper struct {
	IResponseWriter[ResponseWriterWrapper]
	RW *http.ResponseWriter
}

func (brw *ResponseWriterWrapper) Status(statusCode int) (*ResponseWriterWrapper) {
	(*brw.RW).WriteHeader(statusCode)
	return brw
} 

func (brw *ResponseWriterWrapper) Text(text string) (*ResponseWriterWrapper) {
	(*brw.RW).Write([]byte(text))
	return brw
}

func (brw *ResponseWriterWrapper) Json(data interface{}) (*ResponseWriterWrapper) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return brw.Bytes(bytes)
}

func (brw *ResponseWriterWrapper) Bytes(data []byte) (*ResponseWriterWrapper) {
	(*brw.RW).Write(data)
	return brw
}