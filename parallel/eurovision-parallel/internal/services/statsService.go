package services

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type ViewData struct {
	Elapsed      time.Duration
	BestCountry  string
	WorstCountry string
	SameArtists  map[string]int
	AvgVotes     map[CountryYear]AvgVotes
	MinMax       struct {
		Min     CountryYear
		MinDiff float64
		Max     CountryYear
		MaxDiff float64
	}
	Styles  []string
	Langs   []string
	Females float64
	Males   float64
}

type StatsService struct {
	CountryServ *CountryService
	SongServ    *SongService
}

func (s *StatsService) GetStatsPar() (*ViewData, error) {
	var wg sync.WaitGroup
	var vd = &ViewData{}

	start := time.Now()

	wg.Add(1)
	go func() {
		defer wg.Done()
		vd.BestCountry, _ = s.CountryServ.MostFirstPlaces()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		vd.WorstCountry, _ = s.CountryServ.MostLastPlaces()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		vd.SameArtists, _ = s.SongServ.FindSameArtists()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		vd.AvgVotes, _ = s.CountryServ.AvgJuryTelevote()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		min, minDiff, max, maxDiff, _ := s.CountryServ.MinMaxDiffVotes()
		vd.MinMax = struct {
			Min     CountryYear
			MinDiff float64
			Max     CountryYear
			MaxDiff float64
		}{
			Min:     *min,
			MinDiff: minDiff,
			Max:     *max,
			MaxDiff: maxDiff,
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		vd.Styles, vd.Langs, _ = s.SongServ.StyleLangFinal()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		vd.Females, vd.Males, _ = s.SongServ.GenderPercents()
	}()

	wg.Wait()

	vd.Elapsed = time.Since(start)

	return vd, nil
}

func (s *StatsService) GetStatsSeq() (*ViewData, error) {
	start := time.Now()

	mcountry, err := s.CountryServ.MostFirstPlaces()
	if err != nil {
		log.Warnf("Error getting most priced country: %v", err)
		return nil, err
	}
	lcountry, err := s.CountryServ.MostLastPlaces()
	if err != nil {
		log.Warnf("Error getting country with most last places: %v", err)
		return nil, err
	}
	artists, err := s.SongServ.FindSameArtists()
	if err != nil {
		log.Warnf("Error getting same artists: %v", err)
		return nil, err
	}
	avgVotes, err := s.CountryServ.AvgJuryTelevote()
	if err != nil {
		log.Warnf("Error getting average votes: %v", err)
		return nil, err
	}
	min, minDiff, max, maxDiff, err := s.CountryServ.MinMaxDiffVotes()
	if err != nil {
		log.Warnf("Error getting min and max votes: %v", err)
		return nil, err
	}
	styles, langs, err := s.SongServ.StyleLangFinal()
	if err != nil {
		log.Warnf("Error getting best styles and languages: %v", err)
		return nil, err
	}
	females, males, err := s.SongServ.GenderPercents()
	if err != nil {
		log.Warnf("Error getting gender winners percents: %v", err)
		return nil, err
	}

	elapsed := time.Since(start)
	log.Info(mcountry)
	log.Info(lcountry)
	for k, v := range artists {
		log.Info(k, ":", v)
	}
	for k, v := range avgVotes {
		log.Info(k.Year, ",", k.Country, ":", v.Jury, ",", v.Televote)
	}
	log.Info(min.Country, minDiff, max.Country, maxDiff)

	for _, style := range styles {
		log.Info(style)
	}
	for _, lang := range langs {
		log.Info(lang)
	}
	log.Info(females)
	log.Info(males)

	return &ViewData{
		Elapsed:      elapsed,
		BestCountry:  mcountry,
		WorstCountry: lcountry,
		SameArtists:  artists,
		AvgVotes:     avgVotes,
		MinMax: struct {
			Min     CountryYear
			MinDiff float64
			Max     CountryYear
			MaxDiff float64
		}{
			Min:     *min,
			MinDiff: minDiff,
			Max:     *max,
			MaxDiff: maxDiff,
		},
		Styles:  styles,
		Langs:   langs,
		Females: females,
		Males:   males,
	}, nil
}
