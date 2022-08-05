package helper

import (
	log "github.com/sirupsen/logrus"
)

type Closer interface {
	Close() error
}

func HandledClose(toClose Closer) {
	err := toClose.Close()
	IsError(err)
}

func IsError(err error) bool {
	if err != nil {
		log.Error(err.Error())
		return true
	}
	return false
}
