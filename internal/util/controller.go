package util

import "net/http"

type Controller interface {
	Handle(res http.ResponseWriter, req *http.Request)
}