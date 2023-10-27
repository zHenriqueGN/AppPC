package database

import (
	"database/sql"
	"errors"

	"github.com/zHenriqueGN/AppPC/internal/entity"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (c *CategoryDB) Create(name, description string) (*entity.Category, error) {
	category := entity.NewCategory(name, description)
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", category.ID, category.Name, category.Description)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category entity.Category
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		return &category, nil
	}
	return nil, ErrCategoryNotFound
}

func (c *CategoryDB) FindAll() ([]*entity.Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (c *CategoryDB) FindByCourseID(courseID string) (*entity.Category, error) {
	row := c.db.QueryRow(
		`SELECT A.id, A.name, A.description FROM categories A
		 INNER JOIN courses B 
		 ON A.id = B.category_id
		 WHERE B.id = $1
		`, courseID,
	)
	var category entity.Category
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
