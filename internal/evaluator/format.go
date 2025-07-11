package evaluator

import (
	"regexp"
)

const (
	intPlacehoder = iota
	floatPlacehoder
	stringPlacehoder
	boolPlacehoder
	unsupported
)

// validateFormatString checks if format contains exactly one scanf placeholder
func validateFormatString(format string) error {
	placeholderRegex := regexp.MustCompile(`%(?:%|[#+ -]?(?:\*|\d+)?(?:\.(?:\*|\d+))?[diouxXeEfFgGaAcspvt])`)
	matches := placeholderRegex.FindAllString(format, -1)

	// Filter out escaped %% which are literal % characters
	actualPlaceholders := 0
	for _, match := range matches {
		if match != "%%" {
			actualPlaceholders++
		}
	}

	if actualPlaceholders == 0 {
		return ErrMissedPlaceholder()
	}
	if actualPlaceholders > 1 {
		return ErrManyPlaceholders(actualPlaceholders)
	}
	return nil
}

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
	return unsupported
}
