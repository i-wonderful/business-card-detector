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

// InsertSpaceIfNeeded - Вставляет пробелы если нужно. После подстроки должны идти пробелы или конец строки.
func InsertSpaceIfNeeded(s, substring string) string {
	lowerS := strings.ToLower(s)
	lowerSubstring := strings.ToLower(substring)
	if strings.HasPrefix(lowerS, lowerSubstring) && len(s) > len(substring) &&
		unicode.IsLetter(rune(s[len(substring)])) {
		//	(s[len(substring)] != ' ' && s[len(substring)] != ',' && s[len(substring)] != '.')
		return s[:len(substring)] + " " + s[len(substring):]
	}
	return s
}

// Функция для нахождения минимального значения среди трех чисел
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// LevenshteinDistance - Функция для вычисления расстояния Левенштейна
func LevenshteinDistance(a, b string) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	// Создание двумерного среза для хранения значений
	dp := make([][]int, len(a)+1)
	for i := range dp {
		dp[i] = make([]int, len(b)+1)
	}

	// Инициализация базовых значений
	for i := 0; i <= len(a); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		dp[0][j] = j
	}

	// Заполнение таблицы
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j]+1,   // Удаление
					dp[i][j-1]+1,   // Вставка
					dp[i-1][j-1]+1, // Замена
				)
			}
		}
	}

	return dp[len(a)][len(b)]
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
