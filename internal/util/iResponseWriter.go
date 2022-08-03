package util

type IResponseWriter[T any] interface {
	Status(statusCode int) (*T)
	Text(text string) (*T)
	Json(data interface{}) (*T)
	Bytes(data []byte) (*T)
}