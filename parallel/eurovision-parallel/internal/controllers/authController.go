package controllers

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/dtos"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/helpers"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/repositories"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/services"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	var s *services.UserService = &services.UserService{
		Repo: &repositories.UserRepository{},
	}

	if helpers.AlreadyLogin(r) {
		http.Redirect(w, r, "/stats", http.StatusFound)
		return
	}
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			log.Warnf("Error parsing template file(register): %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = tmpl.Execute(w, nil); err != nil {
			log.Warnf("Error executing template(register): %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		var userDTO = &dtos.UserDTO{
			FirstName: r.FormValue("firstName"),
			LastName:  r.FormValue("lastName"),
			Email:     r.FormValue("email"),
			Password:  r.FormValue("password"),
		}
		user, err := userDTO.ToUser()
		if err != nil {
			log.Warnf("Error hashing password: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := s.Create(*user); err != nil {
			log.Warnf("Error registering user: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var s *services.UserService = &services.UserService{
		Repo: &repositories.UserRepository{},
	}

	if helpers.AlreadyLogin(r) {
		http.Redirect(w, r, "/stats", http.StatusFound)
		return
	}

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			log.Warnf("Error parsing template file(login): %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = tmpl.Execute(w, nil); err != nil {
			log.Warnf("Error executing template(login): %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		username := r.FormValue("username")

		ok, err := s.AuthCheck(username, r.FormValue("password"))
		if !ok || err != nil {
			log.Warnf("Error authenticating %s user: %v", username, err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		session, err := server.App.Store.Get(r, "my_session")
		if err != nil {
			log.Warnf("Error getting session for login page: %v", err)
			return
		}

		session.Values["user"] = username

		err = session.Save(r, w)
		if err != nil {
			log.Warnf("Error saving cookies on login: %v", err)
		}

		http.Redirect(w, r, "/stats", http.StatusFound)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := server.App.Store.Get(r, "my_session")
	if err != nil {
		log.Warnf("Error getting session for login page: %v", err)
		return
	}

	session.Values["user"] = nil
	if err = session.Save(r, w); err != nil {
		log.Warnf("Error saving cookies on logout: %v", err)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
