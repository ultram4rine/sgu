package services

import (
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) Create(user models.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetByName(username string) (*models.User, error) {
	return s.Repo.GetByName(username)
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	return s.Repo.GetByEmail(email)
}

func (s *UserService) AuthCheck(email, password string) (bool, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Email == email {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
