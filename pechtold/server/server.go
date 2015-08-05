package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-recaptcha/recaptcha"
)

var (
	captcha *recaptcha.Recaptcha
	db      *sql.DB
)

// Run is a forever blocking call that starts the pechtold server
func Run() {
	parseFlags()

	setupDB()

	setupCaptcha()

	http.HandleFunc("/pechtold/upload", handleUpload)
	err := http.ListenAndServe(flags.HTTPAddress, nil)
	if err != nil {
		log.Fatalf("error during ListenAndServe: %v", err)
	}
}

func setupDB() {
	var err error
	db, err = sql.Open("postgres", "user=geenpeil dbname=geenpeil") // TODO: use flags
	if err != nil {
		log.Fatalf("error setting up db conn (open): %v", err)
	}

	// force connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("error setting up db conn (ping): %v", err)
	}
}

func setupCaptcha() {
	captcha = recaptcha.New(flags.CaptchaSecret)
}
