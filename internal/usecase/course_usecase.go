package usecase

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/repository"
	"strconv"
	"sync"
)

type ICourseUsecase interface {
	GetAll() ([]*domain.Course, error)
	GetByID(id string) (*domain.Course, error)
	Create(course *domain.Course) error
	Update(course *domain.Course) error
	Delete(id string) error
}

type CourseUsecase struct {
	CourseRepo repository.ICourseRepository
}

var (
	courseUsecaseInstance *CourseUsecase
	courseUsecaseOnce     sync.Once
)

func NewCourseUsecase(repo repository.ICourseRepository) ICourseUsecase {
	courseUsecaseOnce.Do(func() {
		courseUsecaseInstance = &CourseUsecase{}
		courseUsecaseInstance.CourseRepo = repo
	})
	return courseUsecaseInstance
}

func (u *CourseUsecase) GetAll() ([]*domain.Course, error) {
	return u.CourseRepo.GetAll()
}

func (u *CourseUsecase) GetByID(id string) (*domain.Course, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return u.CourseRepo.GetByID(intID)
}

func (u *CourseUsecase) Create(course *domain.Course) error {
	err := course.Validate()
	if err != nil {
		return err
	}

	return u.CourseRepo.Create(course)
}

func (u *CourseUsecase) Update(course *domain.Course) error {
	return u.CourseRepo.Update(course)
}

func (u *CourseUsecase) Delete(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return u.CourseRepo.Delete(intID)
}
