package repositories

import (
	"database/sql"
	"errors"

	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
)

type JuryResultsRepository struct{}

func (r *JuryResultsRepository) ListAll() ([]models.JuryResult, error) {
	juryResults := []models.JuryResult{}
	if err := server.App.DB.Select(&juryResults, "SELECT * FROM jury_results"); err != nil {
		return nil, err
	}
	return juryResults, nil
}

func (r *JuryResultsRepository) ListAllInYear(year string) ([]models.JuryResult, error) {
	juryResults := []models.JuryResult{}
	if err := server.App.DB.Select(&juryResults, "SELECT * FROM jury_results WHERE year = ?", year); err != nil {
		return nil, err
	}
	return juryResults, nil
}

func (r *JuryResultsRepository) AvgCountryScoreYear(year, country string) (float64, error) {
	var avg sql.NullFloat64
	if err := server.App.DB.Get(&avg, "SELECT AVG(score) FROM jury_results WHERE year = ? AND contestant = ?;", year, country); err != nil {
		return -1, err
	}
	if avg.Valid {
		return avg.Float64, nil
	} else {
		return -1, errors.New("no such year")
	}
}
