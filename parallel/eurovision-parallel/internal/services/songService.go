package services

import (
	"sort"

	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/repositories"
)

type SongService struct {
	SongRepo     *repositories.SongsRepository
	ContestRepo  *repositories.ContestsRepository
	JuryRepo     *repositories.JuryResultsRepository
	TelevoteRepo *repositories.TelevoteResultsRepository
}

func (s *SongService) FindSameArtists() (map[string]int, error) {
	songs, err := s.SongRepo.ListAll()
	if err != nil {
		return nil, err
	}

	var artistsCount = make(map[string]int)
	for _, song := range songs {
		artistsCount[song.Artist_name] += 1
	}

	for k, v := range artistsCount {
		if v == 1 {
			delete(artistsCount, k)
		}
	}

	return artistsCount, nil
}

func (s *SongService) StyleLangFinal() ([]string, []string, error) {
	var (
		styleCount = make(map[string]int)
		langCount  = make(map[string]int)
	)

	songs, err := s.SongRepo.ListSongsInFinal()
	if err != nil {
		return nil, nil, err
	}

	for _, song := range songs {
		styleCount[song.Style] += 1
		langCount[song.Language] += 1
	}

	sKeys := make([]string, 0, len(styleCount))
	for key := range styleCount {
		sKeys = append(sKeys, key)
	}

	sort.SliceStable(sKeys, func(i, j int) bool {
		return styleCount[sKeys[i]] > styleCount[sKeys[j]]
	})

	lKeys := make([]string, 0, len(langCount))
	for key := range langCount {
		lKeys = append(lKeys, key)
	}

	sort.SliceStable(lKeys, func(i, j int) bool {
		return langCount[lKeys[i]] > langCount[lKeys[j]]
	})

	return sKeys, lKeys, nil
}

func (s *SongService) GenderPercents() (float64, float64, error) {
	contests, err := s.ContestRepo.ListAll()
	if err != nil {
		return -1, -1, err
	}

	var (
		males   float64
		females float64
	)

	for _, c := range contests {
		var countryVotes = make(map[string]int)

		juryResults, err := s.JuryRepo.ListAllInYear(c.Year)
		if err != nil {
			return -1, -1, err
		}
		televoteResults, err := s.TelevoteRepo.ListAllInYear(c.Year)
		if err != nil {
			return -1, -1, err
		}

		for _, jr := range juryResults {
			countryVotes[jr.Contestant] += jr.Score
		}
		for _, tr := range televoteResults {
			countryVotes[tr.Contestant] += tr.Score
		}

		keys := make([]string, 0, len(countryVotes))
		for key := range countryVotes {
			keys = append(keys, key)
		}

		if len(keys) > 0 {
			sort.SliceStable(keys, func(i, j int) bool {
				return countryVotes[keys[i]] > countryVotes[keys[j]]
			})

			song, err := s.SongRepo.GetSongByYearAndCountry(c.Year, keys[0])
			if err != nil {
				return -1, -1, err
			}

			if song.Gender == "Female" {
				females += 1
			} else {
				males += 1
			}
		}
	}

	return (females / (males + females)) * 100, (males / (males + females)) * 100, nil
}
