package database

import (
	"database/sql"

	"github.com/zHenriqueGN/AppPC/internal/entity"
)

type CourseDB struct {
	db *sql.DB
}

func NewCourse(db *sql.DB) *CourseDB {
	return &CourseDB{db: db}
}

func (c *CourseDB) Create(course *entity.Course) error {
	stmt, err := c.db.Prepare("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.ID, course.Name, course.Description, course.CategoryID)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseDB) FindAll() ([]entity.Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []entity.Course
	for rows.Next() {
		var course entity.Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}
	return courses, nil
}

func (c *CourseDB) FindByCategoryID(categoryID string) ([]entity.Course, error) {
	stmt, err := c.db.Prepare("SELECT id, name, description, category_id FROM courses WHERE category_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []entity.Course
	for rows.Next() {
		var course entity.Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
