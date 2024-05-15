package usecase

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/repository"
	"strconv"
	"sync"
)

type IGradeUsecase interface {
	GetAll() ([]*domain.Grade, error)
	GetByID(id string) (*domain.Grade, error)
	Create(grade *domain.Grade) error
	Update(grade *domain.Grade) error
	Delete(id string) error
	GetByStudentID(studentID string) ([]*domain.Grade, error)
	GetByCourseID(courseID string) ([]*domain.Grade, error)
	GetByProfessorID(professorID string) ([]*domain.Grade, error)
}

type GradeUsecase struct {
	GradeRepo repository.IGradeRepository
}

var (
	gradeUsecaseInstance *GradeUsecase
	gradeUsecaseOnce     sync.Once
)

func NewGradeUsecase(repo repository.IGradeRepository) IGradeUsecase {
	gradeUsecaseOnce.Do(func() {
		gradeUsecaseInstance = &GradeUsecase{
			GradeRepo: repo,
		}
	})
	return gradeUsecaseInstance
}

func (uc *GradeUsecase) GetAll() ([]*domain.Grade, error) {
	return uc.GradeRepo.GetAll()
}

func (uc *GradeUsecase) GetByID(id string) (*domain.Grade, error) {
	gradeID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return uc.GradeRepo.GetByID(gradeID)
}

func (uc *GradeUsecase) Create(grade *domain.Grade) error {
	err := grade.Validate()
	if err != nil {
		return err
	}
	return uc.GradeRepo.Create(grade)
}

func (uc *GradeUsecase) Update(grade *domain.Grade) error {
	err := grade.Validate()
	if err != nil {
		return err
	}
	return uc.GradeRepo.Update(grade)
}

func (uc *GradeUsecase) Delete(id string) error {
	gradeID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return uc.GradeRepo.Delete(gradeID)
}

func (uc *GradeUsecase) GetByStudentID(studentID string) ([]*domain.Grade, error) {
	intStudentID, err := strconv.Atoi(studentID)
	if err != nil {
		return nil, err
	}
	return uc.GradeRepo.GetByStudentID(intStudentID)
}

func (uc *GradeUsecase) GetByCourseID(courseID string) ([]*domain.Grade, error) {
	intCourseID, err := strconv.Atoi(courseID)
	if err != nil {
		return nil, err
	}
	return uc.GradeRepo.GetByCourseID(intCourseID)
}

func (uc *GradeUsecase) GetByProfessorID(professorID string) ([]*domain.Grade, error) {
	intProfessorID, err := strconv.Atoi(professorID)
	if err != nil {
		return nil, err
	}
	return uc.GradeRepo.GetByProfessorID(intProfessorID)
}
