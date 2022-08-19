package util

import (
	"math/rand"
)

// TODO: ADD KEY REGENERATION!!

type KeyManager struct {
	keys 		 map[string]*string
	allowedRunes *string
}

func (km *KeyManager) Init() {
	km.keys = map[string]*string{}
	km.AllowedRunes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func (km *KeyManager) AllowedRunes(runes string) {
	km.allowedRunes = &runes
}

func (km *KeyManager) GetKey(id string) *string {
	if val, ok := km.keys[id]; ok  {
		return val
	}
	return nil
}

func (km *KeyManager) GenerateKey(id string, length int) *string {
	keyRunes := make([]rune, length)
	runesLen := len(*(km.allowedRunes))

	for length--; length >= 0; length-- {
		keyRunes[length] = rune((*km.allowedRunes)[rand.Intn(runesLen)])
	}

	key := string(keyRunes)
	km.keys[id] = &key
	return &key
}
