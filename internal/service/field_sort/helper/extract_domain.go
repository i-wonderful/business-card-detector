package helper

import "strings"

// ExtractDomainNameFromUrl - извлечение имени домена из URL
func ExtractDomainNameFromUrl(url string) string {
	// Удаляем протокол, если он есть
	if strings.Contains(url, "://") {
		parts := strings.Split(url, "://")
		url = parts[len(parts)-1]
	}
	url = strings.TrimPrefix(url, "www.")
	parts := strings.Split(url, ".")

	// Возвращаем первую часть (имя домена)
	if len(parts) > 0 {
		return parts[0]
	}

	return ""
}
