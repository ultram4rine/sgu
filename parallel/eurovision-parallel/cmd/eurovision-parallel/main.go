package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/controllers"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/helpers"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
)

func main() {
	var err error
	server.App.DB, err = sqlx.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatalf("cannot open db connection: %v", err)
	}
	defer server.App.DB.Close()
	server.App.Store = sessions.NewCookieStore([]byte("HZ2gX8btMMUg6m0IvD0xjQ5kIb8ZZkBw"), []byte("gp5zEGQUm1w94vM3MvQUBCaUiEnjmxQ1"))

	router := mux.NewRouter()
	router.HandleFunc("/register", controllers.RegisterController)
	router.HandleFunc("/login", controllers.LoginController)
	router.HandleFunc("/logout", controllers.LogoutHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":3456",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("web server started")
	log.Fatal(srv.ListenAndServe())
}

// AuthCheck is a middleware for handlers.
func AuthCheck(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !helpers.AlreadyLogin(r) {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}

		handler(w, r)
	}
}
