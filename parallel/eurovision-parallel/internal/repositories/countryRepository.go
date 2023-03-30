package repositories

import (
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
)

type CountryRepository struct{}

func (r *CountryRepository) ListAll() ([]models.Country, error) {
	countries := []models.Country{}
	if err := server.App.DB.Select(&countries, "SELECT * FROM countries"); err != nil {
		return nil, err
	}
	return countries, nil
}
