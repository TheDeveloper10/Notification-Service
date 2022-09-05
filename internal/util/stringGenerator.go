package util

import (
	"math/rand"
	"time"
)

func NewStringGenerator() *StringGenerator {
	sg := StringGenerator{}
	sg.Init()
	return &sg
}

type StringGenerator struct {
	allowedRunes *string
}

func (sg *StringGenerator) Init() {
	sg.SetAllowedRunes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixMicro())
}

func (sg *StringGenerator) SetAllowedRunes(runes string) {
	sg.allowedRunes = &runes
}

func (sg *StringGenerator) GenerateString(length int) string {
	strRunes := make([]rune, length)
	runesLen := len(*(sg.allowedRunes))

	for i := 0; i < length; i++ {
		strRunes[i] = rune((*sg.allowedRunes)[rand.Intn(runesLen)])
	}

	return string(strRunes)
}
