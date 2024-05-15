package usecase

import (
	"fmt"
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/repository"
	"strconv"
	"sync"
)

type IStudentUsecase interface {
	GetAll() ([]*domain.Student, error)
	GetByID(id string) (*domain.Student, error)
	Create(student *domain.Student) error
	Update(student *domain.Student) error
	Delete(id string) error
}

type StudentUsecase struct {
	StudentRepo repository.IStudentRepository
}

var (
	studentUsecaseInstance *StudentUsecase
	once                   sync.Once
)

func NewStudentUsecase(repo repository.IStudentRepository) IStudentUsecase {
	once.Do(func() {
		studentUsecaseInstance = &StudentUsecase{
			StudentRepo: repo,
		}
	})
	return studentUsecaseInstance
}

func (uc *StudentUsecase) GetAll() ([]*domain.Student, error) {
	return uc.StudentRepo.GetAll()
}

func (uc *StudentUsecase) GetByID(id string) (*domain.Student, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return uc.StudentRepo.GetByID(intID)
}

func (uc *StudentUsecase) Create(student *domain.Student) error {
	err := student.Validate()
	if err != nil {
		return fmt.Errorf("error validating student data: %v", err)
	}

	return uc.StudentRepo.Create(student)
}

func (uc *StudentUsecase) Update(student *domain.Student) error {
	return uc.StudentRepo.Update(student)
}

func (uc *StudentUsecase) Delete(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return uc.StudentRepo.Delete(intID)
}
