package util

import (
	"github.com/sirupsen/logrus"
)

type Closer interface {
	Close() error
}

func HandledClose(toClose Closer) {
	err := toClose.Close()
	ManageError(err)
}

func ManageError(err error) bool {
	if err != nil {
		logrus.Error(err.Error())
		return true
	}
	return false
}