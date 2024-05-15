package usecase

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/repository"
	"strconv"
	"sync"
)

type IProfessorUsecase interface {
	GetAll() ([]*domain.Professor, error)
	GetByID(id string) (*domain.Professor, error)
	Create(professor *domain.Professor) error
	Update(professor *domain.Professor) error
	Delete(id string) error
}

type ProfessorUsecase struct {
	ProfessorRepo repository.IProfessorRepository
}

var (
	professorUsecaseInstance *ProfessorUsecase
	professorUsecaseOnce     sync.Once
)

func NewProfessorUsecase(repo repository.IProfessorRepository) IProfessorUsecase {
	professorUsecaseOnce.Do(func() {
		professorUsecaseInstance = &ProfessorUsecase{}
		professorUsecaseInstance.ProfessorRepo = repo
	})
	return professorUsecaseInstance
}

func (u *ProfessorUsecase) GetAll() ([]*domain.Professor, error) {
	return u.ProfessorRepo.GetAll()
}

func (u *ProfessorUsecase) GetByID(id string) (*domain.Professor, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return u.ProfessorRepo.GetByID(intID)
}

func (u *ProfessorUsecase) Create(professor *domain.Professor) error {
	return u.ProfessorRepo.Create(professor)
}

func (u *ProfessorUsecase) Update(professor *domain.Professor) error {
	return u.ProfessorRepo.Update(professor)
}

func (u *ProfessorUsecase) Delete(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return u.ProfessorRepo.Delete(intID)
}
