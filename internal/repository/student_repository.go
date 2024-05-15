package repository

import (
	"database/sql"
	"golang-technical-test/database"
	"golang-technical-test/internal/domain"
	"sync"
)

type IStudentRepository interface {
	GetAll() ([]*domain.Student, error)
	GetByID(id int) (*domain.Student, error)
	Create(student *domain.Student) error
	Update(student *domain.Student) error
	Delete(id int) error
}

type StudentRepository struct {
	db *database.Database
}

var (
	studentsRepoOnce     sync.Once
	studentsRepoInstance *StudentRepository
)

func NewStudentRepository(db *database.Database) IStudentRepository {
	studentsRepoOnce.Do(func() {
		studentsRepoInstance = &StudentRepository{}
		studentsRepoInstance.db = db
	})
	return studentsRepoInstance
}

func (r *StudentRepository) GetAll() ([]*domain.Student, error) {
	rows, err := r.db.Query("SELECT ID, Name, Lastname, DateOfBirth, Address, Email FROM Students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*domain.Student
	for rows.Next() {
		var s domain.Student
		err = rows.Scan(&s.ID, &s.Name, &s.LastName, &s.DateOfBirth, &s.Address, &s.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, &s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepository) GetByID(id int) (*domain.Student, error) {
	row := r.db.QueryRow("SELECT ID, Name, Lastname, DateOfBirth, Address, Email FROM Students WHERE ID = ?", id)

	var s domain.Student
	err := row.Scan(&s.ID, &s.Name, &s.LastName, &s.DateOfBirth, &s.Address, &s.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// There were no rows, but otherwise no error occurred
			return nil, nil
		}
		return nil, err
	}

	return &s, nil
}

func (r *StudentRepository) Create(student *domain.Student) error {
	stmt, err := r.db.Prepare("INSERT INTO Students (Name, Lastname, DateOfBirth, Address, Email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(student.Name, student.LastName, student.DateOfBirth, student.Address, student.Email)
	if err != nil {
		return err
	}

	courseID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	student.ID = int(courseID)

	return nil
}

func (r *StudentRepository) Update(student *domain.Student) error {
	stmt, err := r.db.Prepare("UPDATE Students SET Name = ?, Lastname = ?, DateOfBirth = ?, Address = ?, Email = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.Name, student.LastName, student.DateOfBirth, student.Address, student.Email, student.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *StudentRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM Students WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
