package repository

import (
	"database/sql"
	"backend/internal/models"
)

type AnnouncementRepository struct {
	db *sql.DB
}

func NewAnnouncementRepository(db *sql.DB) *AnnouncementRepository {
	return &AnnouncementRepository{db: db}
}

func (r *AnnouncementRepository) GetAll() ([]models.Announcement, error) {
	query := `
		SELECT id, title, content, published_date, is_pinned, created_at
		FROM announcements
		ORDER BY is_pinned DESC, published_date DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var announcements []models.Announcement
	for rows.Next() {
		var a models.Announcement
		err := rows.Scan(
			&a.ID, &a.Title, &a.Content, &a.PublishedDate,
			&a.IsPinned, &a.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		announcements = append(announcements, a)
	}

	return announcements, nil
}

func (r *AnnouncementRepository) GetByID(id int) (*models.Announcement, error) {
	query := `
		SELECT id, title, content, published_date, is_pinned, created_at
		FROM announcements WHERE id = $1
	`

	var a models.Announcement
	err := r.db.QueryRow(query, id).Scan(
		&a.ID, &a.Title, &a.Content, &a.PublishedDate,
		&a.IsPinned, &a.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *AnnouncementRepository) Create(req *models.CreateAnnouncementRequest) (*models.Announcement, error) {
	query := `
		INSERT INTO announcements (title, content, is_pinned)
		VALUES ($1, $2, $3)
		RETURNING id, title, content, published_date, is_pinned, created_at
	`

	var a models.Announcement
	err := r.db.QueryRow(query, req.Title, req.Content, req.IsPinned).Scan(
		&a.ID, &a.Title, &a.Content, &a.PublishedDate,
		&a.IsPinned, &a.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}