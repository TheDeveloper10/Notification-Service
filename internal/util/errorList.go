package util

import "notification-service/internal/util/iface"

func NewErrorList() iface.IErrorList {
	return &ErrorList{
		Errors: []string{},
	}
}

func ErrorListFromTextError(err string) iface.IErrorList {
	return &ErrorList{
		Errors: []string{err},
	}
}

type ErrorList struct {
	iface.IErrorList `json:"-"`
	Errors           []string `json:"errors"`
}

func (el *ErrorList) Merge(errs iface.IErrorList) iface.IErrorList {
	el.Errors = append(el.Errors, (*errs.GetErrors())...)
	return el
}

func (el *ErrorList) AddError(err error) iface.IErrorList {
	if err == nil {
		return el
	}
	el.AddErrorFromString(err.Error())
	return el
}

func (el *ErrorList) AddErrorFromString(err string) iface.IErrorList {
	el.Errors = append(el.Errors, err)
	return el
}

func (er *ErrorList) GetErrors() *[]string {
	return &er.Errors
}

func (el *ErrorList) ErrorsCount() int {
	return len(el.Errors)
}
