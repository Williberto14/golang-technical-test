package models

// Course struct
type Course struct {
	CourseID    int    `json:"course_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
