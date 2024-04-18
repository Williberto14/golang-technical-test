package models

// Grade struct
type Grade struct {
	GradeID   int     `json:"grade_id"`
	StudentID int     `json:"student_id"`
	CourseID  int     `json:"course_id"`
	Grade     float64 `json:"grade"`
}
