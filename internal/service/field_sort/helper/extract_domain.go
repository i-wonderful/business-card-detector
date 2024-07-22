package helper

import "strings"

// ExtractDomainFromUrl - извлечение имени домена из URL
func ExtractDomainFromUrl(url string) string {
	url = ExtractDomainAndZoneUrlFromUrl(url)
	parts := strings.Split(url, ".")

	// Возвращаем первую часть (имя домена)
	if len(parts) > 0 {
		return parts[0]
	}

	return ""
}

func ExtractDomainAndZoneUrlFromUrl(url string) string {
	// Удаляем протокол, если он есть
	if strings.Contains(url, "://") {
		parts := strings.Split(url, "://")
		url = parts[len(parts)-1]
	}
	url = strings.TrimPrefix(url, "www.")

	return url
}

// ExtractMainNameDomainAndZone - извлечение основного имени, домена и зоны
func ExtractMainNameDomainAndZone(email string) (string, string, string) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", "", "" // Неверный формат адреса электронной почты
	}
	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) < 2 {
		return "", "", "" // Неверный формат домена
	}
	return parts[0], domainParts[0], domainParts[1]
}
