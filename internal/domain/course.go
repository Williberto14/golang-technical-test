package domain

import "golang-technical-test/utils"

type Course struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (c *Course) Validate() error {
	vali := utils.GetValidator()
	return vali.Struct(c)
}
