package field_sort

import "regexp"

var emailRegex = regexp.MustCompile(`(?i)(?:Mail:\s*)?([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)
var phoneRegex = regexp.MustCompile(`\+?[\d\s\(\)-]{6,16}\d`)
var phoneRegexExtract = regexp.MustCompile(`[\+\(]?[0-9 .\(\)-]{7,}`)
var nameRegex = regexp.MustCompile(`^[A-ZА-Я][A-ZА-Яa-zа-я-]+ [A-ZА-Я][A-ZА-Яa-zа-я-]+([ \-][A-ZА-Я][A-ZА-Яa-zа-я-]+)?$`)
var telegramRegex = regexp.MustCompile(`(?:https?://)?(t\.me/|@)[A-Za-z][A-Za-z0-9_]{4,31}(?:\s[A-Za-z0-9_]+)*`)
var emailCheckRegex = regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)
