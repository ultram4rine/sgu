package repositories

import (
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/models"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
)

type SongsRepository struct{}

func (r *SongsRepository) ListAll() ([]models.Song, error) {
	songs := []models.Song{}
	if err := server.App.DB.Select(&songs, "SELECT * FROM songs"); err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *SongsRepository) ListSongsInFinal() ([]models.Song, error) {
	songs := []models.Song{}
	if err := server.App.DB.Select(&songs, "SELECT * FROM songs WHERE final_draw_position IS NOT NULL"); err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *SongsRepository) GetSongByYearAndCountry(year, country string) (*models.Song, error) {
	song := &models.Song{}
	if err := server.App.DB.Get(song, "SELECT * FROM songs WHERE year = ? AND country = ?;", year, country); err != nil {
		return nil, err
	}
	return song, nil
}
