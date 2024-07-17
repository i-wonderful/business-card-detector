package manage_str

import (
	"strings"
	"unicode"
)

// SplitByWorldsAndUppercase - разбиение строки на слова по пробелам и большим буквам
func SplitByWorldsAndUppercase(val string) []string {
	words := strings.Split(val, " ")
	var result []string

	for _, word := range words {
		splitItems := splitByUpperCase(word)
		result = append(result, splitItems...)
	}
	return result
}

func splitByUpperCase(s string) []string {
	var result []string
	var word strings.Builder

	for _, char := range s {
		if unicode.IsUpper(char) && word.Len() > 0 {
			result = append(result, word.String())
			word.Reset()
		}
		word.WriteRune(char)
	}

	if word.Len() > 0 {
		result = append(result, word.String())
	}

	return result
}
