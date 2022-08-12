package iface

type IRequest interface {
	Validate() []error
}

type IRequestEntity[T any] interface {
	IRequest
	ToEntity() *T
}