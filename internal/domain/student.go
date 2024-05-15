package domain

import "golang-technical-test/utils"

type Student struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Email       string `json:"email" validate:"required"`
}

func (v *Student) Validate() error {
	vali := utils.GetValidator()
	return vali.Struct(v)
}
