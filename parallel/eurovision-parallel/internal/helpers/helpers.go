package helpers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/ultram4rine/ssu/parallel/eurovision-parallel/internal/server"
	"golang.org/x/crypto/bcrypt"
)

// AlreadyLogin checks is user already logged in.
func AlreadyLogin(r *http.Request) bool {
	session, err := server.App.Store.Get(r, "my_session")
	if err != nil {
		log.Printf("Error getting session: %s", err)
		return false
	}

	return session.Values["user"] != nil
}

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
