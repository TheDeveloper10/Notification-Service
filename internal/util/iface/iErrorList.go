package iface

type IErrorList interface {
	Merge(IErrorList) IErrorList

	AddError(error) IErrorList
	AddErrorFromString(string) IErrorList

	GetErrors() *[]string
	ErrorsCount() int
}
