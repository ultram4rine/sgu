package dtos

import "github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"

type UserDTO struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *UserDTO) ToUser() *models.User {
	return &models.User{
		Username: u.FirstName + " " + u.LastName,
		Email:    u.Email,
		Password: u.Password,
	}
}
