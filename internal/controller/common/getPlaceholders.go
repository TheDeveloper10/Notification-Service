package common

import (
	"github.com/sirupsen/logrus"
	"regexp"
)

func GetPlaceholders(text *string) string {
	regex, err := regexp.Compile("@{[^$]*?}")
	if err != nil {
		logrus.Error("Failed to compile regex")
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
			set[placeholder] = empty
		}
	}

	return res
}
