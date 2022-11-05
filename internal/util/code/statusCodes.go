package code

type StatusCode int8

const (
	StatusSuccess  = 0
	StatusNotFound = 1
	StatusError    = 2
	StatusExpired  = 3
)