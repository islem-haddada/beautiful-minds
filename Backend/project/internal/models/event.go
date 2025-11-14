package models

import "time"

type Event struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Date            time.Time `json:"date"`
	Location        string    `json:"location"`
	ImageURL        *string   `json:"image_url"`
	MaxParticipants int       `json:"max_participants"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateEventRequest struct {
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Date            string  `json:"date"`
	Location        string  `json:"location"`
	ImageURL        *string `json:"image_url"`
	MaxParticipants int     `json:"max_participants"`
}

type RegisterEventRequest struct {
	MemberID int `json:"member_id"`
}
