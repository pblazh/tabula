package evaluator

import (
	"regexp"
)

const (
	intPlacehoder = iota
	floatPlacehoder
	stringPlacehoder
	boolPlacehoder
	datePlacehoder
)

// detectPlaceholderType detects the type of scanf placeholder in the format string
func detectPlaceholderType(format string) int {
	if matched, _ := regexp.MatchString(`%[#+ -]?(?:\*|\d+)?[diouxX]`, format); matched {
		return intPlacehoder
	}
	if matched, _ := regexp.MatchString(`%[#+ -]?(?:\*|\d+)?(?:\.(?:\*|\d+))?[eEfFgGaA]`, format); matched {
		return floatPlacehoder
	}
	if matched, _ := regexp.MatchString(`%[#+ -]?(?:\*|\d+)?[sc]`, format); matched {
		return stringPlacehoder
	}
	if matched, _ := regexp.MatchString(`%[#+ -]?(?:\*|\d+)?[t]`, format); matched {
		return boolPlacehoder
	}
	return datePlacehoder
}
