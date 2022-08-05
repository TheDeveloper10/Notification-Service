package helper

import (
	log "github.com/sirupsen/logrus"
)

type Closer interface {
	Close() error
}

func HandledClose(toClose Closer) {
	err := toClose.Close()
	if err != nil {
		log.Error(err.Error())
	}
}
