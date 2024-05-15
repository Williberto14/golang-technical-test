package domain

import "golang-technical-test/utils"

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (v *Course) Validate() error {
	vali := utils.GetValidator()
	return vali.Struct(v)
}
