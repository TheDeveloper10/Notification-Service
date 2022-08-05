package util

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	if err != nil {
		log.Error(err.Error())
		return brw
	}
	return brw
}

func (brw *responseWriterWrapper) Json(data interface{}) IResponseWriter {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
		return brw
	}
	return brw.Bytes(bytes)
}

func (brw *responseWriterWrapper) Bytes(data []byte) IResponseWriter {
	_, err := (*brw.rw).Write(data)
	if err != nil {
		log.Error(err.Error())
		return brw
	}
	return brw
}

func WrapResponseWriter(res *http.ResponseWriter) IResponseWriter {
	return &responseWriterWrapper {
		rw: res,
	}
}