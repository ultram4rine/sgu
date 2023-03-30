package dtos

import (
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserDTO struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *UserDTO) ToUser() (*models.User, error) {
	encPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Username: u.FirstName + " " + u.LastName,
		Email:    u.Email,
		Password: string(encPass),
	}, nil
}
