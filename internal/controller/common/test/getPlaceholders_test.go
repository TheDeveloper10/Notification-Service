package test

import (
	"notification-service/internal/controller/common"
	"testing"
)

func TestGetPlaceholders(t *testing.T) {
	tests := []struct {
		text                string
		expectedReturnValue string
	}{
		{ text: "Hello!", expectedReturnValue: "" },
		{ text: "Hello, @{firstName}!", expectedReturnValue: "firstName" },
		{ text: "Hello, @{firstName}! We are @{company-name}", expectedReturnValue: "firstName, company-name" },
		{ text: "Hello, @{firstName}! We are @{company-name}. @{company-name} tries to improve user experience...", expectedReturnValue: "firstName, company-name" },
		{ text: "Hello, @{firstName}! We are @{company-name}. @{FirstName} we are the best!", expectedReturnValue: "firstName, company-name, FirstName" },
		{ text: "@{a}@{b}{c}", expectedReturnValue: "a, b" },
		{ text: "@{a}{b}@{c}{@c}@{d}{@e}", expectedReturnValue: "a, c, d" },
		{ text: "@{a}{b}@{c}{@c}@{d}{@e}@{d}", expectedReturnValue: "a, c, d" },
		{ text: "@{a}{b}@{c}{@c}@{d}{@e}@{a}", expectedReturnValue: "a, c, d" },
		{ text: "@{a}@{B}@{ C}@{D }@{ E }", expectedReturnValue: "a, B,  C, D ,  E "},
	}

	for testId, test := range tests {
		returnValue := common.GetPlaceholders(&test.text)
		if returnValue != test.expectedReturnValue {
			t.Errorf("Test: %d\tExpected: %s\tReceived: %s", testId, test.expectedReturnValue, returnValue)
		}
	}
}