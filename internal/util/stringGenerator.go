package util

import "math/rand"

type StringGenerator struct {
	allowedRunes *string
}

func (sg *StringGenerator) Init() {
	sg.SetAllowedRunes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func (sg *StringGenerator) SetAllowedRunes(runes string) {
	sg.allowedRunes = &runes
}

func (sg *StringGenerator) GenerateString(length int) string {
	strRunes := make([]rune, length)
	runesLen := len(*(sg.allowedRunes))

	for length--; length >= 0; length-- {
		strRunes[length] = rune((*sg.allowedRunes)[rand.Intn(runesLen)])
	}

	return string(strRunes)
}
