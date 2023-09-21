package model

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name              string   `json:"name" bson:"name" xml:"name" form:"name" query:"name" validate:"required,min=5"`
	ID                string   `json:"id" bson:"_id" xml:"id" form:"id" query:"id" validate:"required,min=5"`
	Jobs              []string `json:"jobs,omitempty" bson:"jobs,omitempty" validate:"max=5"`
	CurrentlyEmployed bool     `json:"currentlyEmployed,omitempty" bson:"currently_employed,omitempty"`
	Experience        int      `json:"experience,omitempty" bson:"experience,omitempty"`
	Salary            float64  `json:"salary,omitempty" bson:"salary,omitempty"`
}

type UserValidator struct {
	validator *validator.Validate
}

func (u *UserValidator) Validate(i interface{}) error {
	return u.validator.Struct(i)
}
