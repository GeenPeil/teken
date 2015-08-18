package server

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"net"
	"net/http"
	"net/mail"
	"text/template"

	"github.com/GeenPeil/teken/data"
	"github.com/GeenPeil/teken/storage"

	"github.com/davecgh/go-spew/spew"
	"github.com/jpoehls/gophermail"
	"github.com/lib/pq"
)

var (
	fieldErr   = "form values missing or invalid"
	captchaErr = "captcha invalid"
	imgErr     = "image invalid"
	mailErr    = "mail has been used"
)
var (
	mailFrom = mail.Address{
		Name:    "GeenPeil no-reply",
		Address: "no-reply@teken.geenpeil.nl",
	}
	mailSubject = "GeenPeil verificatie mail"
)

var (
	tmplVerificationMailPlainText = template.Must(template.New("plain").Parse(`Beste {{.Handtekening.Voornaam}} {{.Handtekening.Tussenvoegsel}} {{.Handtekening.Achternaam}},

{{.VerificatieLink}}

Bedankt,
Leger des Peils`))

	tmplVerificationMailHTML = template.Must(template.New("html").Parse(`Beste {{.Handtekening.Voornaam}} {{.Handtekening.Tussenvoegsel}} {{.Handtekening.Achternaam}},

<a href="{{.VerificatieLink}}" >Klik hier</a> of gebruik onderstaande URL.

{{.VerificatieLink}}

Bedankt,
Leger des Peils`))
)

func (s *Server) newSubmitHandlerFunc() http.HandlerFunc {

	stmtInsertNawHash, err := s.db.Prepare(`INSERT INTO nawhashes (hash) VALUES ($1)`)
	if err != nil {
		log.Fatalf("error preparing stmtInsertNAWHash: %v", err)
	}

	stmtInsertHandtekening, err := s.db.Prepare(`INSERT INTO handtekeningen (insert_time, iphash) VALUES (NOW(), $1) RETURNING ID`)
	if err != nil {
		log.Fatalf("error preparing stmtInsertHandtekening: %v", err)
	}

	stmtSelectMail, err := s.db.Prepare(`SELECT COUNT(*) FROM handtekeningen WHERE mailhash = $1 AND mailcheckdone = true`)
	if err != nil {
		log.Fatalf("error preparing stmtSelectMail: %v", err)
	}

	saver, err := storage.NewSaver(s.options.StoragePubkeyFile, s.options.StorageLocation)
	if err != nil {
		log.Fatalf("error creating storage.Saver: %v", err)
	}

	type mailData struct {
		Handtekening    *data.Handtekening
		VerificatieLink string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
		if xRealIP := r.Header.Get("X-Real-IP"); xRealIP != "" {
			remoteIP = xRealIP
		}

		s.verbosef("have request from remoteIP=%s origin=%s method=%s", remoteIP, r.Header.Get("Origin"), r.Method)

		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}
		if r.Method != "POST" {
			http.Error(w, "invalid http method", http.StatusBadRequest)
			return
		}

		h := &data.Handtekening{}
		err := json.NewDecoder(r.Body).Decode(h)
		r.Body.Close()
		if err != nil {
			http.Error(w, "input json error", http.StatusInternalServerError)
			log.Printf("error decoding json in request from %s: %v", remoteIP, err)
			return
		}

		if s.options.Verbose {
			spew.Dump(h)
		}

		out := &SubmitOutput{}

		{
			// check captcha
			if !s.options.CaptchaDisable {
				valid, err := s.captcha.Verify(h.CaptchaResponse, remoteIP)
				if err != nil {
					http.Error(w, "server error", http.StatusInternalServerError)
					log.Printf("error verifying captcha: %v", err)
					return
				}
				if !valid {
					out.Error = captchaErr
					log.Printf("invalid captcha in request from %s", remoteIP)
					goto Response
				}
			}

			// checken of alle data is ingevuld
			if len(h.Voornaam) == 0 {
				out.Error = fieldErr
				log.Printf("missing field voornaam in request from %s", remoteIP)
				goto Response
			}

			if len(h.Achternaam) == 0 {
				out.Error = fieldErr
				log.Printf("missing field achternaam in request from %s", remoteIP)
				goto Response
			}

			if len(h.Geboortedatum) == 0 {
				out.Error = fieldErr
				log.Printf("missing field geboortedatum in request from %s", remoteIP)
				goto Response
			}

			if len(h.Geboorteplaats) == 0 {
				out.Error = fieldErr
				log.Printf("missing field geboorteplaats in request from %s", remoteIP)
				goto Response
			}

			if len(h.Straat) == 0 {
				out.Error = fieldErr
				log.Printf("missing field straat in request from %s", remoteIP)
				goto Response
			}

			if len(h.Huisnummer) == 0 {
				out.Error = fieldErr
				log.Printf("missing field huisnummer in request from %s", remoteIP)
				goto Response
			}

			if len(h.Postcode) == 0 {
				out.Error = fieldErr
				log.Printf("missing field postcode in request from %s", remoteIP)
				goto Response
			}

			if len(h.Woonplaats) == 0 {
				out.Error = fieldErr
				log.Printf("missing field woonplaats in request from %s", remoteIP)
				goto Response
			}

			if len(h.Email) == 0 {
				out.Error = fieldErr
				log.Printf("missing field email in request from %s", remoteIP)
				goto Response
			}

			if len(h.Handtekening) == 0 {
				out.Error = fieldErr
				log.Printf("missing field handtekening in request from %s", remoteIP)
				goto Response
			}

			_, err = png.Decode(bytes.NewBuffer(h.Handtekening))
			if err != nil {
				out.Error = imgErr
				log.Printf("invalid image from %s: %v", remoteIP, err)
				goto Response
			}

			mailHash := sha256.New()
			mailHash.Write(s.hashingSalt)
			mailHashBytes := mailHash.Sum(bytes.ToLower(bytes.TrimSpace([]byte(h.Email))))
			var exists int
			err := stmtSelectMail.QueryRow(mailHashBytes).Scan(&exists)
			if err != nil {
				http.Error(w, "server error", http.StatusInternalServerError)
				log.Printf("error checking for duplicate email hash: %v", err)
				return
			}
			if exists > 0 {
				out.Error = mailErr
				goto Response
			}

			// all ok
			out.Success = true

			// naw hash check (false positive)
			nawHash := sha256.New()
			nawHash.Write(s.hashingSalt)
			nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Voornaam))))
			nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Achternaam))))
			nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Geboortedatum))))
			nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Geboorteplaats))))
			nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Straat))))
			nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Huisnummer))))
			nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Postcode))))
			nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Woonplaats))))
			nawHashBytes := nawHash.Sum(nil)
			_, err = stmtInsertNawHash.Exec(nawHashBytes)
			if err != nil {
				if perr, ok := err.(*pq.Error); ok {
					if perr.Code == "23505" {
						log.Printf("duplicate n.a.w. hash from %s: %x", remoteIP, nawHashBytes)
						goto Response // return direclty with a 'false positive'
					} else {
						log.Printf("error inserting naw hash from %s: %v", remoteIP, err)
						http.Error(w, "server error", http.StatusInternalServerError)
						return
					}
				}
			}

			// insert handtekening entry into db, get inserted ID
			ipHash := sha256.New()
			ipHash.Write(s.hashingSalt)
			ipHashBytes := ipHash.Sum([]byte(remoteIP))
			mailcheckString := randomString(25)
			mailcheckHashBytes := sha256.New().Sum([]byte(mailcheckString))
			var ID uint64
			err = stmtInsertHandtekening.QueryRow(ipHashBytes, mailHashBytes, mailcheckHashBytes).Scan(&ID)
			if err != nil {
				log.Printf("error inserting handtekening entry in db for %s: %v", remoteIP, err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}

			// save to disk
			err = saver.Save(ID, h)
			if err != nil {
				log.Printf("error saving handtekening for %s: %v", remoteIP, err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}

			// send mail
			md := &mailData{
				Handtekening:    h,
				VerificatieLink: fmt.Sprintf("https://teken.geenpeil.nl/pechtold/verify?mailhash=%s&check=%s", string(mailHashBytes), mailcheckString),
			}
			var bodyBuf = &bytes.Buffer{}
			var htmlBuf = &bytes.Buffer{}
			err = tmplVerificationMailPlainText.Execute(bodyBuf, md)
			if err != nil {
				log.Printf("error executing tmplVerificationMailPlainText: %v", err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			err = tmplVerificationMailHTML.Execute(htmlBuf, md)
			if err != nil {
				log.Printf("error executing tmplVerificationMailHTML: %v", err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			mailMessage := &gophermail.Message{
				From: mailFrom,
				To: []mail.Address{
					{
						Name:    fmt.Sprintf("%s %s %s", h.Voornaam, h.Tussenvoegsel, h.Achternaam),
						Address: h.Email,
					},
				},
				Subject:  mailSubject,
				Body:     bodyBuf.String(),
				HTMLBody: htmlBuf.String(),
			}
			err = gophermail.SendMail(s.options.SMTPServer, nil, mailMessage)
			if err != nil {
				log.Printf("error sending verification mail: %v", err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
		}

	Response:
		s.verbosef("response to request from %s is %t with err %s", remoteIP, out.Success, out.Error)
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			log.Printf("error encoding response json: %v", err)
		}
	}
}
