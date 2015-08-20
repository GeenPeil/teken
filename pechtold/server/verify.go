package server

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
	"text/template"
)

var (
	tmplResultaat = template.Must(template.New("plain").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>GeenPeil verificatie</title>
		<style>
			body {
				background-color: #04339a;
				color: white;
				font-family: Ariel, sans-serif;
			}
			.logo{
				display: inline-block;
				background-image: url(/pechtold-static/geenpeillogo.png);
				background-size: cover;
				width: 55px;
				height: 55px;
			}
			.logo-text{
				display: inline-block;
				margin-bottom: 35px;
				vertical-align: middle;
			}
		</style>
	</head>
	<body>
		<center>
			<h1><div class="logo"></div><span class="logo-text">GEENPEIL</span></h1>
			<h3>{{.Text}}</h3>
			{{if .Verdrietig}}
				<img src="/pechtold-static/verdrietig.jpg" />
			{{end}}
		</center>
	</body>
</html>
`))
)

func (s *Server) newVerifyHandlerFunc() http.HandlerFunc {

	stmtSelectMailUsed, err := s.db.Prepare(`SELECT mailcheckhash FROM handtekeningen WHERE mailhash = $1 AND mailcheckdone = true`)
	if err != nil {
		log.Fatalf("error preparing stmtSelectMailUsed: %v", err)
	}

	stmtUpdateMailCheck, err := s.db.Prepare(`UPDATE handtekeningen SET mailcheckdone = true WHERE mailhash = $1 AND mailcheckhash = $2`)
	if err != nil {
		log.Fatalf("error preparing stmtUpdateMailCheck: %v", err)
	}

	type pageData struct {
		Text       string
		Verdrietig bool
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		remoteIP := s.resolveRemoteIP(r)

		out := &pageData{}

		{
			mailHashBytes, err := base64.URLEncoding.DecodeString(r.FormValue("mailhash"))
			if err != nil {
				out.Text = "Ongeldige verificatie link."
				goto SendResult

			}
			check := r.FormValue("check")
			checkHash := sha256.New()
			checkHash.Write([]byte(check))
			checkHashBytes := checkHash.Sum(nil)

			var existingCheck []byte
			err = stmtSelectMailUsed.QueryRow(mailHashBytes).Scan(&existingCheck)
			if err != nil && err != sql.ErrNoRows {
				log.Printf("error checking mail used for remote ip %s: %v", remoteIP, err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			} else if err == nil {
				// either the link has been used, or the same email was verified earlier
				if bytes.Equal(checkHashBytes, existingCheck) {
					out.Text = "Je handtekening is al geverifieerd."
					out.Verdrietig = true
					goto SendResult

				}
				out.Text = "Dit email address is inmiddels gebruikt voor het verifieren van een andere handtekening. Teken a.u.b. opnieuw met een ander email adres."
				goto SendResult
			}

			res, err := stmtUpdateMailCheck.Exec(mailHashBytes, checkHashBytes)
			if err != nil {
				log.Printf("error updating mail check for %s: %v", remoteIP, err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			n, _ := res.RowsAffected()
			if n != 1 {
				out.Text = "Ongeldige verificatie link."
				goto SendResult

			}

			out.Text = "Je handtekening is succesvol geverifieerd."
			out.Verdrietig = true
			goto SendResult
		}

	SendResult:
		err = tmplResultaat.Execute(w, out)
		if err != nil {
			log.Printf("error rendering template for %s: %v", remoteIP, err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}
}
