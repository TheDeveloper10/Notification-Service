package iface

import "github.com/TheDeveloper10/rem"

type IController interface {
	CreateRoutes(router *rem.Router)
}