package util

import (
	"encoding/json"
	"net/http"
)

type BetterResponseWriter struct {
	RW *http.ResponseWriter
}

func (brw *BetterResponseWriter) Status(statusCode int) (*BetterResponseWriter) {
	(*brw.RW).WriteHeader(statusCode)
	return brw
} 

func (brw *BetterResponseWriter) Text(text string) (*BetterResponseWriter) {
	(*brw.RW).Write([]byte(text))
	return brw
}

func (brw *BetterResponseWriter) Json(data interface{}) (*BetterResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return brw.Bytes(bytes)
}

func (brw *BetterResponseWriter) Bytes(data []byte) (*BetterResponseWriter) {
	(*brw.RW).Write(data)
	return brw
}