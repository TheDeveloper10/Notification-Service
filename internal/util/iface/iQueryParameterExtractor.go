package iface

type IQueryParameterExtractor interface {
	GetPositiveInteger(key string, defaultValue int) (*int, error)
	GetInteger(key string) (*int, error)
	GetString(key string) *string
}