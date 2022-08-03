package dto

type AbstractRequest interface {
	Validate() (bool, string)
}