package helper

import (
	"card_detector/internal/util/calc"
	"strings"
	"unicode"
)

func ContainsIgnoreCase(s, substr string) bool {
	if s == "" || substr == "" {
		return false
	}
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func CountDigits(input string) int {
	count := 0
	for _, char := range input {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// NotContainsLetters - проверяет, не содержит ли строка буквы
func NotContainsLetters(phone string) bool {
	for i, r := range phone {
		if unicode.IsDigit(r) {
			// Проверить соседние символы на наличие букв
			if (i > 0 && unicode.IsLetter(rune(phone[i-1]))) || (i < len(phone)-1 && unicode.IsLetter(rune(phone[i+1]))) {
				return false
			}
		}
	}
	return true
}

func ClearTrashSymbols(val string) string {
	val = strings.TrimFunc(val, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '+' && r != '@' && r != '(' && r != ')'
	})
	return val
}

// InsertSpaceIfNeeded - Вставляет пробелы если нужно. После подстроки должны идти пробелы или конец строки.
func InsertSpaceIfNeeded(s, substring string) string {
	lowerS := strings.ToLower(s)
	lowerSubstring := strings.ToLower(substring)
	if strings.HasPrefix(lowerS, lowerSubstring) && len(s) > len(substring) && s[len(substring)] != ' ' {
		return s[:len(substring)] + " " + s[len(substring):]
	}
	return s
}

// StringDifference - calculates how different two strings are.
func StringDifference(a, b string) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if a == b {
		return 0
	}

	aRunes := []rune(a)
	bRunes := []rune(b)

	length := len(aRunes)
	if len(bRunes) < length {
		length = len(bRunes)
	}

	diff := 0
	for i := 0; i < length; i++ {
		if aRunes[i] != bRunes[i] {
			diff++
		}
	}

	diff += calc.Abs(len(aRunes) - len(bRunes))
	return diff
}
