package dto

import (
	"errors"
)

type TemplatePlaceholder struct {
	Key   string `json:"key"`
	Value string `json:"val"`
}

func (tp *TemplatePlaceholder) Validate() error {
	if tp.Key == "" {
		return errors.New("'key' must be given on each placeholder")
	}

	return nil
}
