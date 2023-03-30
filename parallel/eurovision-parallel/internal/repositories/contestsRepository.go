package repositories

import (
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
)

type ContestsRepository struct{}

func (r *ContestsRepository) ListAll() ([]models.Contest, error) {
	contests := []models.Contest{}
	if err := server.App.DB.Select(&contests, "SELECT * FROM contests"); err != nil {
		return nil, err
	}
	return contests, nil
}
