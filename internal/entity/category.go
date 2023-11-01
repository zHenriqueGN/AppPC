package entity

import "github.com/google/uuid"

type Category struct {
	ID          string
	Name        string
	Description string
}

func NewCategory(name, description string) *Category {
	return &Category{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
	}
}
