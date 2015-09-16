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
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.4.0/css/font-awesome.min.css">
		<link href='https://fonts.googleapis.com/css?family=Montserrat' rel='stylesheet' type='text/css'>
		<script>
		function shareLink(element) {
			var href = element.getAttribute('href');
			var newWindow = window.open(href, 'newwindow', 'left=300, top=200, width=600, height=475');
			return !newWindow;
		}
		</script>
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
			.geenpeil-share {
				background-color: white;
				display: inline-block;
				padding: 7px;
				font-family: 'Montserrat', 'Arial', sans-serif;
				box-shadow: 1px 1px 1px 1px rgba(80,80,80,0.6);
				border-radius: 4px;
				text-align: center;
			}
			.geenpeil-share .share-button {
				transition: all 0.3s ease;
				font-size: 40px;
				margin: 4px;
				text-decoration: none;
				text-shadow: 2px 2px 2px #ccc;
			}
			.geenpeil-share .share-title {
				color: black;
				display: inline-block;
				margin: 5px;
			}
			.geenpeil-share .share-button:hover {
				text-shadow: 3px 3px 5px #bbb;
			}
			.geenpeil-share .share-button.facebook {
				color: #395693;
			}
			.geenpeil-share .share-button.twitter {
				color: #39a9e0;
			}
			.geenpeil-share .share-button.googleplus {
				color: #d14836;
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
			<br/>
			<div class="geenpeil-share">
				<span class="share-title">Help en <b>DEEL GEENPEIL</b> met iedereen!</span>
				<br/>
				<a title="Deel op Facebook" class="share-button facebook fa fa-facebook-official" href="https://www.facebook.com/sharer/sharer.php?u=teken.geenpeil.nl" onclick="return shareLink(this);"></a>
				<a title="Deel op Twitter" class="share-button twitter fa fa-twitter-square" href="https://twitter.com/home?status=Ik heb getekend! Red de democratie en teken ook op https%3A%2F%2Fteken.geenpeil.nl %23geenpeil" onclick="return shareLink(this);"></a>
				<a title="Deel op Google+" class="share-button googleplus fa fa-google-plus-square" href="https://plus.google.com/share?url=www.geenpeil.nl" onclick="return shareLink(this);"></a>
				</div>
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
