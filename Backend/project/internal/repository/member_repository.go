package repository

import (
	"database/sql"
	"beautiful-minds/backend/project/internal/models"
)

type MemberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

func (r *MemberRepository) GetAll() ([]models.Member, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, student_id, 
		       field_of_study, registration_date, is_active, created_at
		FROM members
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.Member
	for rows.Next() {
		var m models.Member
		err := rows.Scan(
			&m.ID, &m.FirstName, &m.LastName, &m.Email, &m.Phone,
			&m.StudentID, &m.FieldOfStudy, &m.RegistrationDate,
			&m.IsActive, &m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		members = append(members, m)
	}

	return members, nil
}

func (r *MemberRepository) GetByID(id int) (*models.Member, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, student_id,
		       field_of_study, registration_date, is_active, created_at
		FROM members WHERE id = $1
	`

	var m models.Member
	err := r.db.QueryRow(query, id).Scan(
		&m.ID, &m.FirstName, &m.LastName, &m.Email, &m.Phone,
		&m.StudentID, &m.FieldOfStudy, &m.RegistrationDate,
		&m.IsActive, &m.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *MemberRepository) Create(req *models.CreateMemberRequest) (*models.Member, error) {
	query := `
		INSERT INTO members (first_name, last_name, email, phone, student_id, field_of_study)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, first_name, last_name, email, phone, student_id, 
		          field_of_study, registration_date, is_active, created_at
	`

	var m models.Member
	err := r.db.QueryRow(
		query, req.FirstName, req.LastName, req.Email, 
		req.Phone, req.StudentID, req.FieldOfStudy,
	).Scan(
		&m.ID, &m.FirstName, &m.LastName, &m.Email, &m.Phone,
		&m.StudentID, &m.FieldOfStudy, &m.RegistrationDate,
		&m.IsActive, &m.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}