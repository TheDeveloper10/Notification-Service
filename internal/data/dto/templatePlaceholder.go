package dto

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
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

func GetPlaceholders(text *string) string {
	regex, err := regexp.Compile("@{[^$]*?}")
	if err != nil {
		log.Error("Failed to compile regex")
		return ""
	}

	matches := regex.FindAllString(*text, -1)
	if len(matches) <= 0 {
		return ""
	}

	type void struct{}
	var empty void

	first := true
	res := ""
	set := make(map[string]void)
	for _, placeholder := range matches {
		placeholder = placeholder[2:len(placeholder)-1]

		if _, ok := set[placeholder]; !ok {
			if first {
				res += placeholder
				first = false
			} else {
				res += ", " + placeholder
			}
		} else {
			set[placeholder] = empty
		}
	}

	return res
}

func FillPlaceholders(text string, placeholders *[]TemplatePlaceholder) (*string, error) {
	for i := 0; i < len(*placeholders); i++ {
		placeholder := &((*placeholders)[i])
		if err := placeholder.Validate(); err != nil {
			return nil, err
		}

		key := "@{" + placeholder.Key + "}"
		text = strings.ReplaceAll(text, key, placeholder.Value)
	}

	return &text, nil
}