package controllers

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/repositories"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/services"
)

type Data struct {
	ParData *services.ViewData
	SeqData *services.ViewData
}

func StatsController(w http.ResponseWriter, r *http.Request) {
	var statsService = &services.StatsService{
		CountryServ: &services.CountryService{
			ContestRepo:  &repositories.ContestsRepository{},
			JuryRepo:     &repositories.JuryResultsRepository{},
			TelevoteRepo: &repositories.TelevoteResultsRepository{},
		},
		SongServ: &services.SongService{
			SongRepo:     &repositories.SongsRepository{},
			ContestRepo:  &repositories.ContestsRepository{},
			JuryRepo:     &repositories.JuryResultsRepository{},
			TelevoteRepo: &repositories.TelevoteResultsRepository{},
		},
	}

	if r.Method == "GET" {
		seqData, err := statsService.GetStatsSeq()
		if err != nil {
			log.Warnf("Error getting seq data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		parData, err := statsService.GetStatsPar()
		if err != nil {
			log.Warnf("Error getting seq data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/stats.html")
		if err != nil {
			log.Warnf("Error parsing template file(stats): %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := Data{
			ParData: parData,
			SeqData: seqData,
		}

		if err = tmpl.Execute(w, data); err != nil {
			log.Warnf("Error executing template(stats): %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
