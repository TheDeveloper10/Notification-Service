package util

import "testing"

func TestStringGenerator_GenerateString(t *testing.T) {
	sg := StringGenerator{}
	sg.Init()

	const lengthOfStringToGenerate = 64

	for i := 0; i < 1000; i++ {
		str := sg.GenerateString(lengthOfStringToGenerate)
		if len(str) != lengthOfStringToGenerate {
			t.Errorf("Generated a string of %d characters instead of %d characters", len(str), lengthOfStringToGenerate)
			return
		}
	}
}