package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image/png"
	"log"
	"net"
	"net/http"

	"github.com/GeenPeil/teken/data"
)

var (
	fieldErr   = "form values missing or invalid"
	captchaErr = "captcha invalid"
	imgErr     = "image invalid"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	h := &data.Handtekening{}
	err := json.NewDecoder(r.Body).Decode(h)
	if err != nil {
		http.Error(w, "input json error", http.StatusInternalServerError)
		log.Printf("error decoding json in request from %s: %v", r.RemoteAddr, err)
	}

	out := &uploadOutput{}

	{
		// check captcha
		remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
		valid, err := captcha.Verify(h.CaptchaResponse, remoteIP)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			log.Printf("error verifying captcha: %v", err)
			return
		}
		if !valid {
			out.Error = captchaErr
			log.Printf("invalid captcha in request from %s", r.RemoteAddr)
			goto Response
		}

		// checken of alle data is ingevuld
		if len(h.Voornaam) == 0 {
			out.Error = fieldErr
			log.Printf("missing field voornaam in request from %s", r.RemoteAddr)
			goto Response
		}
		if len(h.Achternaam) == 0 {
			out.Error = fieldErr
			log.Printf("missing field achternaam in request from %s", r.RemoteAddr)
			goto Response
		}
		if len(h.Geboortedatum) == 0 {
			out.Error = fieldErr
			log.Printf("missing field geboortedatum in request from %s", r.RemoteAddr)
			goto Response
		}
		if len(h.Geboorteplaats) == 0 {
			out.Error = fieldErr
			log.Printf("missing field geboorteplaats in request from %s", r.RemoteAddr)
			goto Response
		}
		if len(h.Straat) == 0 {
			out.Error = fieldErr
			log.Printf("missing field straat in request from %s", r.RemoteAddr)
			goto Response
		}
		if len(h.Huisnummer) == 0 {
			out.Error = fieldErr
			log.Printf("missing field huisnummer in request from %s", r.RemoteAddr)
			goto Response
		}
		if len(h.Postcode) == 0 {
			out.Error = fieldErr
			log.Printf("missing field postcode in request from %s", r.RemoteAddr)
			goto Response
		}
		if len(h.Woonplaats) == 0 {
			out.Error = fieldErr
			log.Printf("missing field woonplaats in request from %s", r.RemoteAddr)
			goto Response
		}
		if len(h.Handtekening) == 0 {
			out.Error = fieldErr
			log.Printf("missing field handtekening in request from %s", r.RemoteAddr)
			goto Response
		}

		// check (decode, etc.) handtekening
		hImgPNG := make([]byte, base64.StdEncoding.DecodedLen(len(h.Handtekening)))
		_, err = base64.StdEncoding.Decode(hImgPNG, h.Handtekening)
		if err != nil {
			out.Error = imgErr
			log.Printf("invalid base64 image received from %s: %v", r.RemoteAddr, err)
			goto Response
		}

		_, err = png.Decode(bytes.NewBuffer(hImgPNG))
		if err != nil {
			out.Error = imgErr
			log.Printf("invalid image from %s: %v", r.RemoteAddr, err)
			goto Response
		}

		// all ok
		out.Success = true

		//++ naw hash check (false positive) TODO: handtekeningen.nashash_ID refers to nawhash.ID, voor e.v.t. opzoekwerk ???? nuttig of niet?
		//++ insert handtekening entry into db, get inserted ID
		//++ save to disk
	}

Response:
	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Printf("error encoding response json: %v", err)
	}
}
