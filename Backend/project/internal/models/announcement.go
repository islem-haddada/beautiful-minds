package models

import "time"

type Announcement struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	PublishedDate time.Time `json:"published_date"`
	IsPinned      bool      `json:"is_pinned"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateAnnouncementRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsPinned bool   `json:"is_pinned"`
}