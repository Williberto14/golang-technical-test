package repository

import (
	"golang-technical-test/database"
	"golang-technical-test/internal/domain"
	"sync"
)

type IEnrollmentRepository interface {
	GetAll() ([]*domain.Enrollment, error)
	GetByID(id int) (*domain.Enrollment, error)
	Create(enrollment *domain.Enrollment) error
	Update(enrollment *domain.Enrollment) error
	Delete(id int) error
	GetByStudentID(studentID int) ([]*domain.Enrollment, error)
	GetByCourseID(courseID int) ([]*domain.Enrollment, error)
}

type EnrollmentRepository struct {
	db *database.Database
}

var (
	enrollmentRepoOnce     sync.Once
	enrollmentRepoInstance *EnrollmentRepository
)

func NewEnrollmentRepository(db *database.Database) IEnrollmentRepository {
	enrollmentRepoOnce.Do(func() {
		enrollmentRepoInstance = &EnrollmentRepository{}
		enrollmentRepoInstance.db = db
	})
	return enrollmentRepoInstance
}

func (r *EnrollmentRepository) GetAll() ([]*domain.Enrollment, error) {
	rows, err := r.db.Query("SELECT * FROM Enrollment")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []*domain.Enrollment
	for rows.Next() {
		enrollment := &domain.Enrollment{}
		err := rows.Scan(&enrollment.ID, &enrollment.StudentID, &enrollment.CourseID)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}

func (r *EnrollmentRepository) GetByID(id int) (*domain.Enrollment, error) {
	row := r.db.QueryRow("SELECT * FROM Enrollment WHERE ID = ?", id)

	enrollment := &domain.Enrollment{}
	err := row.Scan(&enrollment.ID, &enrollment.StudentID, &enrollment.CourseID)
	if err != nil {
		return nil, err
	}

	return enrollment, nil
}

func (r *EnrollmentRepository) Create(enrollment *domain.Enrollment) error {
	stmt, err := r.db.Prepare("INSERT INTO Enrollment (StudentID, CourseID) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(enrollment.StudentID, enrollment.CourseID)
	if err != nil {
		return err
	}

	enrollmentId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	enrollment.ID = int(enrollmentId)

	return nil
}

func (r *EnrollmentRepository) Update(enrollment *domain.Enrollment) error {
	stmt, err := r.db.Prepare("UPDATE Enrollment SET StudentID = ?, CourseID = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(enrollment.StudentID, enrollment.CourseID, enrollment.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *EnrollmentRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM Enrollment WHERE ID = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EnrollmentRepository) GetByStudentID(studentID int) ([]*domain.Enrollment, error) {
	rows, err := r.db.Query("SELECT * FROM Enrollment WHERE StudentID = ?", studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []*domain.Enrollment
	for rows.Next() {
		enrollment := &domain.Enrollment{}
		err := rows.Scan(&enrollment.ID, &enrollment.StudentID, &enrollment.CourseID)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}

func (r *EnrollmentRepository) GetByCourseID(courseID int) ([]*domain.Enrollment, error) {
	rows, err := r.db.Query("SELECT * FROM Enrollment WHERE CourseID = ?", courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []*domain.Enrollment
	for rows.Next() {
		enrollment := &domain.Enrollment{}
		err := rows.Scan(&enrollment.ID, &enrollment.StudentID, &enrollment.CourseID)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}
