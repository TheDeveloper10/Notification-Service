package dto

type AbstractRequest interface {
	Validate() []error
}