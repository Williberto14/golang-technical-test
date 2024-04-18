package models

// Student struct
type Student struct {
	StudentID   int    `json:"student_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
}
