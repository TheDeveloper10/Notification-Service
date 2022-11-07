package test

import (
	"notification-service/internal/controller/common"
	"notification-service/internal/data/dto"
	"testing"
)

func TestFillPlaceholders(t *testing.T) {
	s := func(str string) *string { return &str }
	tests := []struct {
		InText 		   string
		InPlaceholders []dto.TemplatePlaceholder
		OutString      *string
	} {
		{ InText: "a", InPlaceholders: []dto.TemplatePlaceholder{}, OutString: s("a") },
		{ InText: "123a", InPlaceholders: []dto.TemplatePlaceholder{}, OutString: s("123a") },
		{ InText: "123@{a}", InPlaceholders: []dto.TemplatePlaceholder{ { Key: "a", Value: "b" } }, OutString: s("123b") },
		{ InText: "123@{a}", InPlaceholders: []dto.TemplatePlaceholder{ { Key: "c", Value: "b" } }, OutString: s("123@{a}") },
		{ InText: "123@{a}", InPlaceholders: []dto.TemplatePlaceholder{ { Key: "c", Value: "b" } }, OutString: s("123@{a}") },
		{ InText: "12@{3}@{a}", InPlaceholders: []dto.TemplatePlaceholder{ { Key: "c", Value: "b" } }, OutString: s("12@{3}@{a}") },
	}

	for testId, test := range tests {
		res, _ := common.FillPlaceholders(test.InText, test.InPlaceholders)

		if (res == nil && test.OutString == nil) ||
			(res != nil && test.OutString != nil && *res == *test.OutString){
			return
		}
		exp := "<nil>"
		if test.OutString != nil {
			exp = *test.OutString
		}
		rec := "<nil>"
		if res != nil {
			rec = *res
		}

		t.Errorf(
			"Test: %d\tExpected: %s\tReceived: %s",
			testId, exp, rec,
		)
	}
}