package entity

import (
	log "github.com/sirupsen/logrus"
	"regexp"
)

type TemplateEntity struct {
	Id 			int    `json:"id"`
	ContactType string `json:"contactType"`
	Template 	string `json:"template"`
	Language    string `json:"language"`
	Type        string `json:"type"`
}

func (te *TemplateEntity) GetPlaceholders() string {
	regex, err := regexp.Compile("@{[^$]*?}")
	if err != nil {
		log.Error("Failed to compile regex")
		return ""
	}

	matches := regex.FindAllString(te.Template, -1)
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