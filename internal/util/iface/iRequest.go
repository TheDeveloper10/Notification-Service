package iface

type IRequest interface {
	Validate() IErrorList
}

type IRequestEntity[T any] interface {
	IRequest
	ToEntity() *T
}
