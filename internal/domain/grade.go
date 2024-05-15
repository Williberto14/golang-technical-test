package domain

import "golang-technical-test/utils"

type Grade struct {
	ID          int     `json:"id" validate:"required"`
	StudentID   int     `json:"student_id" validate:"required"`
	CourseID    int     `json:"course_id" validate:"required"`
	ProfessorID int     `json:"professor_id" validate:"required"`
	Grade       float64 `json:"grade" validate:"required"`
}

func (v *Grade) Validate() error {
	vali := utils.GetValidator()
	return vali.Struct(v)
}
