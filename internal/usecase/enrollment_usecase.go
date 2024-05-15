package usecase

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/repository"
	"strconv"
	"sync"
)

type IEnrollmentUsecase interface {
	GetAll() ([]*domain.Enrollment, error)
	GetByID(id string) (*domain.Enrollment, error)
	Create(enrollment *domain.Enrollment) error
	Update(enrollment *domain.Enrollment) error
	Delete(id string) error
	GetByStudentID(studentID string) ([]*domain.Enrollment, error)
	GetByCourseID(courseID string) ([]*domain.Enrollment, error)
}

type EnrollmentUsecase struct {
	EnrollmentRepo repository.IEnrollmentRepository
}

var (
	enrollmentUsecaseInstance *EnrollmentUsecase
	enrollmentUsecaseOnce     sync.Once
)

func NewEnrollmentUsecase(repo repository.IEnrollmentRepository) IEnrollmentUsecase {
	enrollmentUsecaseOnce.Do(func() {
		enrollmentUsecaseInstance = &EnrollmentUsecase{
			EnrollmentRepo: repo,
		}
	})
	return enrollmentUsecaseInstance
}

func (u *EnrollmentUsecase) GetAll() ([]*domain.Enrollment, error) {
	return u.EnrollmentRepo.GetAll()
}

func (u *EnrollmentUsecase) GetByID(id string) (*domain.Enrollment, error) {
	enrollmentID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return u.EnrollmentRepo.GetByID(enrollmentID)
}

func (u *EnrollmentUsecase) Create(enrollment *domain.Enrollment) error {
	err := u.EnrollmentRepo.Create(enrollment)
	if err != nil {
		return err
	}

	return nil
}

func (u *EnrollmentUsecase) Update(enrollment *domain.Enrollment) error {
	err := u.EnrollmentRepo.Update(enrollment)
	if err != nil {
		return err
	}

	return nil
}

func (u *EnrollmentUsecase) Delete(id string) error {
	enrollmentID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = u.EnrollmentRepo.Delete(enrollmentID)
	if err != nil {
		return err
	}

	return nil
}

func (u *EnrollmentUsecase) GetByStudentID(studentID string) ([]*domain.Enrollment, error) {
	studentIDInt, err := strconv.Atoi(studentID)
	if err != nil {
		return nil, err
	}

	return u.EnrollmentRepo.GetByStudentID(studentIDInt)
}

func (u *EnrollmentUsecase) GetByCourseID(courseID string) ([]*domain.Enrollment, error) {
	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		return nil, err
	}

	return u.EnrollmentRepo.GetByCourseID(courseIDInt)
}
