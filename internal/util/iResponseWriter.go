package util

type IResponseWriter interface {
	Status(statusCode int) (IResponseWriter)
	Text(text string) (IResponseWriter)
	Json(data interface{}) (IResponseWriter)
	Bytes(data []byte) (IResponseWriter)
}