package dtos

type AbstractRequest interface {
	Validate() (bool, string)
}