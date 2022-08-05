package dto

type AbstractRequest interface {
	Validate() []error
}

type AbstractRequestEntity[T any] interface {
	AbstractRequest
	ToEntity() *T
}