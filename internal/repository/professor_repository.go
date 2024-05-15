package repository

import (
	"golang-technical-test/database"
	"golang-technical-test/internal/domain"
	"sync"
)

type IProfessorRepository interface {
	GetAll() ([]*domain.Professor, error)
	GetByID(id int) (*domain.Professor, error)
	Create(professor *domain.Professor) error
	Update(professor *domain.Professor) error
	Delete(id int) error
}

type ProfessorRepository struct {
	db *database.Database
}

var (
	professorRepoOnce     sync.Once
	professorRepoInstance *ProfessorRepository
)

func NewProfessorRepository(db *database.Database) IProfessorRepository {
	professorRepoOnce.Do(func() {
		professorRepoInstance = &ProfessorRepository{}
		professorRepoInstance.db = db
	})
	return professorRepoInstance
}

func (r *ProfessorRepository) GetAll() ([]*domain.Professor, error) {
	var professors []*domain.Professor
	rows, err := r.db.Query("SELECT ID, Name, Lastname, Email, Specialization FROM Professors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var professor domain.Professor
		err := rows.Scan(&professor.ID, &professor.Name, &professor.LastName, &professor.Email, &professor.Specialization)
		if err != nil {
			return nil, err
		}
		professors = append(professors, &professor)
	}
	return professors, nil
}

func (r *ProfessorRepository) GetByID(id int) (*domain.Professor, error) {
	var professor domain.Professor
	err := r.db.QueryRow("SELECT ID, Name, Lastname, Email, Specialization FROM Professors WHERE ID = ?", id).Scan(&professor.ID, &professor.Name, &professor.LastName, &professor.Email, &professor.Specialization)
	if err != nil {
		return nil, err
	}
	return &professor, nil
}

func (r *ProfessorRepository) Create(professor *domain.Professor) error {
	_, err := r.db.Exec("INSERT INTO Professors (Name, Lastname, Email, Specialization) VALUES (?, ?, ?, ?)", professor.Name, professor.LastName, professor.Email, professor.Specialization)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProfessorRepository) Update(professor *domain.Professor) error {
	_, err := r.db.Exec("UPDATE Professors SET Name = ?, Lastname = ?, Email = ?, Specialization = ? WHERE ID = ?", professor.Name, professor.LastName, professor.Email, professor.Specialization, professor.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProfessorRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Professors WHERE ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}
