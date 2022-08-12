package iface

type IResponseWriter interface {
	Status(statusCode int) IResponseWriter
	Text(text string) IResponseWriter
	Error(err error) IResponseWriter
	Json(data interface{}) IResponseWriter
	Bytes(data []byte) IResponseWriter
}