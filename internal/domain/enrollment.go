package domain

import "golang-technical-test/utils"

type Enrollment struct {
	ID        int `json:"id" validate:"required"`
	StudentID int `json:"student_id" validate:"required"`
	CourseID  int `json:"course_id" validate:"required"`
}

func (v *Enrollment) Validate() error {
	vali := utils.GetValidator()
	return vali.Struct(v)
}
