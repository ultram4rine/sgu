package repositories

import (
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
)

type UserRepository struct{}

func (r *UserRepository) Create(user models.User) error {
	if _, err := server.App.DB.NamedExec("INSERT INTO users (username, email, password) VALUES (:username, :email, :password)", &user); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByName(username string) (*models.User, error) {
	var user = &models.User{}
	if err := server.App.DB.Get(user, "SELECT * FROM users WHERE username = ?", username); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user = &models.User{}
	if err := server.App.DB.Get(user, "SELECT * FROM users WHERE email = ?", email); err != nil {
		return nil, err
	}
	return user, nil
}
