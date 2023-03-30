package services

import (
	"math"
	"sort"

	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/repositories"
)

type CountryService struct {
	CountryRepo  *repositories.CountryRepository
	ContestRepo  *repositories.ContestsRepository
	JuryRepo     *repositories.JuryResultsRepository
	TelevoteRepo *repositories.TelevoteResultsRepository
}

func (s *CountryService) MostFirstPlaces() (string, error) {
	contests, err := s.ContestRepo.ListAll()
	if err != nil {
		return "", err
	}

	var countryPrices = make(map[string]int)
	for _, c := range contests {
		var countryVotes = make(map[string]int)

		juryResults, err := s.JuryRepo.ListAllInYear(c.Year)
		if err != nil {
			return "", err
		}
		televoteResults, err := s.TelevoteRepo.ListAllInYear(c.Year)
		if err != nil {
			return "", err
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

			countryPrices[keys[0]] += 1
		}
	}

	keys := make([]string, 0, len(countryPrices))
	for key := range countryPrices {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return countryPrices[keys[i]] > countryPrices[keys[j]]
	})

	return keys[0], nil
}

func (s *CountryService) MostLastPlaces() (string, error) {
	contests, err := s.ContestRepo.ListAll()
	if err != nil {
		return "", err
	}

	var countryPrices = make(map[string]int)
	for _, c := range contests {
		var countryVotes = make(map[string]int)

		juryResults, err := s.JuryRepo.ListAllInYear(c.Year)
		if err != nil {
			return "", err
		}
		televoteResults, err := s.TelevoteRepo.ListAllInYear(c.Year)
		if err != nil {
			return "", err
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
				return countryVotes[keys[i]] < countryVotes[keys[j]]
			})

			countryPrices[keys[0]] += 1
		}
	}

	keys := make([]string, 0, len(countryPrices))
	for key := range countryPrices {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return countryPrices[keys[i]] > countryPrices[keys[j]]
	})

	return keys[0], nil
}

type CountryYear struct {
	Year    string
	Country string
}

type AvgVotes struct {
	Jury     float64
	Televote float64
}

func (s *CountryService) AvgJuryTelevote() (map[CountryYear]AvgVotes, error) {
	countries, err := s.CountryRepo.ListAll()
	if err != nil {
		return nil, err
	}
	contests, err := s.ContestRepo.ListAll()
	if err != nil {
		return nil, err
	}

	var avgVotes = make(map[CountryYear]AvgVotes)
	for _, contest := range contests {
		for _, country := range countries {
			juryAvg, err := s.JuryRepo.AvgCountryScoreYear(contest.Year, country.Country)
			if err != nil {
				if err.Error() == "no such year" {
					continue
				}
				return nil, err
			}
			televoteAvg, err := s.TelevoteRepo.AvgCountryScoreYear(contest.Year, country.Country)
			if err != nil {
				if err.Error() == "no such year" {
					continue
				}
				return nil, err
			}

			avgVotes[CountryYear{Year: contest.Year, Country: country.Country}] = AvgVotes{
				Jury:     juryAvg,
				Televote: televoteAvg,
			}
		}
	}

	return avgVotes, nil
}

func (s *CountryService) MinMaxDiffVotes() (*CountryYear, float64, *CountryYear, float64, error) {
	avgVotes, err := s.AvgJuryTelevote()
	if err != nil {
		return nil, -1, nil, -1, err
	}

	keys := make([]CountryYear, 0, len(avgVotes))
	for key := range avgVotes {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return math.Abs(avgVotes[keys[i]].Jury-avgVotes[keys[i]].Televote) < math.Abs(avgVotes[keys[j]].Jury-avgVotes[keys[j]].Televote)
	})

	return &keys[0], math.Abs(avgVotes[keys[0]].Jury - avgVotes[keys[0]].Televote),
		&keys[len(keys)-1], math.Abs(avgVotes[keys[len(keys)-1]].Jury - avgVotes[keys[len(keys)-1]].Televote),
		nil
}
