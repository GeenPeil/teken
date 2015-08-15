package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // import postgres driver (registers itself to database/sql)

	"github.com/go-recaptcha/recaptcha"
)

type Server struct {
	captcha *recaptcha.Recaptcha
	db      *sql.DB
	options *Options
}

func New(o *Options) *Server {
	return &Server{
		options: o,
	}
}

// Run is a forever blocking call that starts the pechtold server.
// When setupDoneCh is not nil, it is closed by Run when most of the setup is done and the http server is started (blocking).
func (s *Server) Run(setupDoneCh chan struct{}) {
	s.setupDB()

	s.setupCaptcha()

	http.HandleFunc("/pechtold/submit", s.newSubmitHandlerFunc())

	if setupDoneCh != nil {
		close(setupDoneCh)
	}

	err := http.ListenAndServe(s.options.HTTPAddress, nil)
	if err != nil {
		log.Fatalf("error during ListenAndServe: %v", err)
	}
}

func (s *Server) setupDB() {
	var err error
	s.db, err = sql.Open("postgres", fmt.Sprintf("host=%s sslmode=disable user=pechtold password=pechtold dbname=geenpeil", s.options.PostgresSocketLocation))
	if err != nil {
		log.Fatalf("error setting up db conn (open): %v", err)
	}

	// force connection
	err = s.db.Ping()
	if err != nil {
		log.Fatalf("error setting up db conn (ping): %v", err)
	}
}

func (s *Server) setupCaptcha() {
	s.captcha = recaptcha.New(s.options.CaptchaSecret)
}

func (s *Server) verbosef(format string, args ...interface{}) {
	if s.options.Verbose {
		log.Printf(format, args...)
	}
}
