package iface

import "github.com/TheDeveloper10/rem"

type IHTTPController interface {
	CreateRoutes(router *rem.Router)
}
