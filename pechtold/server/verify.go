package server

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"log"
	"net"
	"net/http"
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

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
		if xRealIP := r.Header.Get("X-Real-IP"); xRealIP != "" {
			remoteIP = xRealIP
		}

		mailHashBytes, err := base64.URLEncoding.DecodeString(r.FormValue("mailhash"))
		if err != nil {
			w.Write([]byte("Ongeldige verificatie link."))
			return
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
				w.Write([]byte("Je handtekening is al geverifieerd. Bedankt."))
				return
			}
			w.Write([]byte("Dit email address is inmiddels gebruikt voor het verifieren van een andere handtekening. Teken a.u.b. opnieuw met een ander email adres."))
			return
		}

		res, err := stmtUpdateMailCheck.Exec(mailHashBytes, checkHashBytes)
		if err != nil {
			log.Printf("error updating mail check for %s: %v", remoteIP, err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		n, _ := res.RowsAffected()
		if n != 1 {
			w.Write([]byte("Ongeldige verificatie link."))
			return
		}

		w.Write([]byte("Bedankt voor het tekenen!"))
	}

}
