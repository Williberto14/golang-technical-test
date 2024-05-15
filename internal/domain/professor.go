package domain

import "golang-technical-test/utils"

type Professor struct {
	ID             int    `json:"id"`
	Name           string `json:"name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Specialization string `json:"specialization,omitempty"`
}

func (v *Professor) Validate() error {
	vali := utils.GetValidator()
	return vali.Struct(v)
}
