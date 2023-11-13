package entity

import "github.com/google/uuid"

type Course struct {
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(name, description, categoryID string) *Course {
	return &Course{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}
}
