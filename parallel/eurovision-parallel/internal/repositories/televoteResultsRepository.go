package repositories

import (
	"database/sql"
	"errors"

	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
)

type TelevoteResultsRepository struct{}

func (r *TelevoteResultsRepository) ListAll() ([]models.TelevoteResult, error) {
	televoteResults := []models.TelevoteResult{}
	if err := server.App.DB.Select(&televoteResults, "SELECT * FROM televote_results"); err != nil {
		return nil, err
	}
	return televoteResults, nil
}

func (r *TelevoteResultsRepository) ListAllInYear(year string) ([]models.TelevoteResult, error) {
	televoteResults := []models.TelevoteResult{}
	if err := server.App.DB.Select(&televoteResults, "SELECT * FROM televote_results WHERE year = ?", year); err != nil {
		return nil, err
	}
	return televoteResults, nil
}

func (r *TelevoteResultsRepository) AvgCountryScoreYear(year, country string) (float64, error) {
	var avg sql.NullFloat64
	if err := server.App.DB.Get(&avg, "SELECT AVG(score) FROM televote_results WHERE year = ? AND contestant = ?;", year, country); err != nil {
		return -1, err
	}
	if avg.Valid {
		return avg.Float64, nil
	} else {
		return -1, errors.New("no such year")
	}
}
