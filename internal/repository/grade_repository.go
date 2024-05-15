package repository

import (
	"golang-technical-test/database"
	"golang-technical-test/internal/domain"
	"sync"
)

type IGradeRepository interface {
	GetAll() ([]*domain.Grade, error)
	GetByID(id int) (*domain.Grade, error)
	Create(grade *domain.Grade) error
	Update(grade *domain.Grade) error
	Delete(id int) error
	GetByStudentID(studentID int) ([]*domain.Grade, error)
	GetByCourseID(courseID int) ([]*domain.Grade, error)
	GetByProfessorID(professorID int) ([]*domain.Grade, error)
}

type GradeRepository struct {
	db *database.Database
}

var (
	gradeRepoOnce     sync.Once
	gradeRepoInstance *GradeRepository
)

func NewGradeRepository(db *database.Database) IGradeRepository {
	gradeRepoOnce.Do(func() {
		gradeRepoInstance = &GradeRepository{}
		gradeRepoInstance.db = db
	})
	return gradeRepoInstance
}

func (r *GradeRepository) GetAll() ([]*domain.Grade, error) {
	rows, err := r.db.Query("SELECT * FROM Grades")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []*domain.Grade
	for rows.Next() {
		grade := &domain.Grade{}
		err := rows.Scan(&grade.ID, &grade.StudentID, &grade.CourseID, &grade.ProfessorID, &grade.Grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

func (r *GradeRepository) GetByID(id int) (*domain.Grade, error) {
	row := r.db.QueryRow("SELECT * FROM Grades WHERE ID = ?", id)

	grade := &domain.Grade{}
	err := row.Scan(&grade.ID, &grade.StudentID, &grade.CourseID, &grade.ProfessorID, &grade.Grade)
	if err != nil {
		return nil, err
	}

	return grade, nil
}

func (r *GradeRepository) Create(grade *domain.Grade) error {
	_, err := r.db.Exec("INSERT INTO Grades (StudentID, CourseID, ProfessorID, Grade) VALUES (?, ?, ?, ?)", grade.StudentID, grade.CourseID, grade.ProfessorID, grade.Grade)
	if err != nil {
		return err
	}

	return nil
}

func (r *GradeRepository) Update(grade *domain.Grade) error {
	_, err := r.db.Exec("UPDATE Grades SET StudentID = ?, CourseID = ?, ProfessorID = ?, Grade = ? WHERE ID = ?", grade.StudentID, grade.CourseID, grade.ProfessorID, grade.Grade, grade.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *GradeRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Grades WHERE ID = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *GradeRepository) GetByStudentID(studentID int) ([]*domain.Grade, error) {
	rows, err := r.db.Query("SELECT * FROM Grades WHERE StudentID = ?", studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []*domain.Grade
	for rows.Next() {
		grade := &domain.Grade{}
		err := rows.Scan(&grade.ID, &grade.StudentID, &grade.CourseID, &grade.ProfessorID, &grade.Grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

func (r *GradeRepository) GetByCourseID(courseID int) ([]*domain.Grade, error) {
	rows, err := r.db.Query("SELECT * FROM Grades WHERE CourseID = ?", courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []*domain.Grade
	for rows.Next() {
		grade := &domain.Grade{}
		err := rows.Scan(&grade.ID, &grade.StudentID, &grade.CourseID, &grade.ProfessorID, &grade.Grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

func (r *GradeRepository) GetByProfessorID(professorID int) ([]*domain.Grade, error) {
	rows, err := r.db.Query("SELECT * FROM Grades WHERE ProfessorID = ?", professorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []*domain.Grade
	for rows.Next() {
		grade := &domain.Grade{}
		err := rows.Scan(&grade.ID, &grade.StudentID, &grade.CourseID, &grade.ProfessorID, &grade.Grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}
