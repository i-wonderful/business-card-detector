package manage_str

import (
	"strings"
	"unicode"
)

func ClearTrashSymbols(val string) string {
	val = strings.TrimFunc(val, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '+' && r != '@' && r != '(' && r != ')'
	})
	return val
}
