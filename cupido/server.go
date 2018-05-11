package cupido

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/http"

	_ "github.com/lib/pq" // import postgres driver (registers itself to database/sql)

	"github.com/go-recaptcha/recaptcha"
)

const reqHashingSaltBits = 256

type Server struct {
	captcha     *recaptcha.Recaptcha
	db          *sql.DB
	options     *Options
	hashingSalt []byte
}

func New(o *Options) *Server {
	s := &Server{
		options: o,
	}

	var err error
	s.hashingSalt, err = base64.StdEncoding.DecodeString(o.HashingSalt)
	if err != nil {
		log.Fatalf("error setting up server, invalid hashing salt: %v", err)
	}
	if len(s.hashingSalt) != reqHashingSaltBits/8 {
		salt := make([]byte, reqHashingSaltBits/8)
		_, err := rand.Read(salt)
		var saltStr string
		if err != nil {
			saltStr = fmt.Sprintf("error generating, %v", err)
		} else {
			saltStr = base64.StdEncoding.EncodeToString(salt)
		}
		log.Fatalf("invalid hashing salt, expecting %d bits, generating one now: `%s`", reqHashingSaltBits, saltStr)
	}

	return s
}

// Run is a forever blocking call that starts the cupido server.
// When setupDoneCh is not nil, it is closed by Run when most of the setup is done and the http server is started (blocking).
func (s *Server) Run(setupDoneCh chan struct{}) {
	s.setupDB()

	s.setupCaptcha()

	http.HandleFunc("/cupido/submit", s.newSubmitHandlerFunc())
	// http.HandleFunc("/cupido/verify", s.newVerifyHandlerFunc())
	http.HandleFunc("/cupido/health-check", s.newHealthCheckHandlerFunc())
	http.HandleFunc("/cupido/api/stats", s.newAPIStatsHandlerFunc())

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
	s.db, err = sql.Open("postgres", fmt.Sprintf("host=%s sslmode=disable dbname=teken", s.options.PostgresSocketLocation))
	if err != nil {
		log.Fatalf("error setting up db conn (open): %v", err)
	}

	s.db.SetMaxOpenConns(10)

	// force connection
	err = s.db.Ping()
	if err != nil {
		log.Fatalf("error setting up db conn (ping): %v", err)
	}
}

func (s *Server) setupCaptcha() {
	s.captcha = recaptcha.New(s.options.CaptchaSecret)
}

func (s *Server) resolveRemoteIP(r *http.Request) string {
	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	if xRealIP := r.Header.Get("X-Real-IP"); xRealIP != "" {
		remoteIP = xRealIP
	}
	return remoteIP
}

func (s *Server) verbosef(format string, args ...interface{}) {
	if s.options.Verbose {
		log.Printf(format, args...)
	}
}
