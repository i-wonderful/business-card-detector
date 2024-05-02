package model

import "time"

type Card struct {
	Id         int64
	Owner      string // todo
	UploadedAt time.Time
	PhotoUrl   string

	// --- Fields --- //
	Email    []string
	Site     []string
	Phone    []string
	Skype    []string
	Telegram []string
	Name     string
	Company  string
	JobTitle string
	Other    string
}
