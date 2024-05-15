package repository

import (
	"golang-technical-test/database"
	"golang-technical-test/internal/domain"
	"sync"
)

type ICourseRepository interface {
	GetAll() ([]*domain.Course, error)
	GetByID(id int) (*domain.Course, error)
	Create(course *domain.Course) error
	Update(course *domain.Course) error
	Delete(id int) error
}

type CourseRepository struct {
	db *database.Database
}

var (
	courseRepoOnce     sync.Once
	courseRepoInstance *CourseRepository
)

func NewCourseRepository(db *database.Database) ICourseRepository {
	courseRepoOnce.Do(func() {
		courseRepoInstance = &CourseRepository{}
		courseRepoInstance.db = db
	})
	return courseRepoInstance
}

func (r *CourseRepository) GetAll() ([]*domain.Course, error) {
	rows, err := r.db.Query("SELECT ID, Name, Description FROM Courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := make([]*domain.Course, 0)
	for rows.Next() {
		course := new(domain.Course)
		err := rows.Scan(&course.ID, &course.Name, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (r *CourseRepository) GetByID(id int) (*domain.Course, error) {
	course := new(domain.Course)
	err := r.db.QueryRow("SELECT ID, Name, Description FROM Courses WHERE ID = ?", id).Scan(&course.ID, &course.Name, &course.Description)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (r *CourseRepository) Create(course *domain.Course) error {
	result, err := r.db.Exec("INSERT INTO Courses (Name, Description) VALUES (?, ?)", course.Name, course.Description)
	if err != nil {
		return err
	}
	courseID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	course.ID = int(courseID)

	return nil
}

func (r *CourseRepository) Update(course *domain.Course) error {
	_, err := r.db.Exec("UPDATE Courses SET Name = ?, Description = ? WHERE ID = ?", course.Name, course.Description, course.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CourseRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Courses WHERE ID = ?", id)
	if err != nil {
		return err
	}

	return nil
}
