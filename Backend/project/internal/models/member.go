package models

import "time"

type Member struct {
	ID              int       `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	StudentID       string    `json:"student_id"`
	FieldOfStudy    string    `json:"field_of_study"`
	RegistrationDate time.Time `json:"registration_date"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateMemberRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	StudentID    string `json:"student_id"`
	FieldOfStudy string `json:"field_of_study"`
}