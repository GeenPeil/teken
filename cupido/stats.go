package cupido

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) newAPIStatsHandlerFunc() http.HandlerFunc {

	stmtSelectTotal, err := s.db.Prepare(`SELECT COUNT(*) FROM handtekeningen`)
	if err != nil {
		log.Fatalf("error preparing stmtSelectTotal: %v", err)
	}

	stmtSelectVerified, err := s.db.Prepare(`SELECT COUNT(*) FROM handtekeningen WHERE mailcheckdone = true`)
	if err != nil {
		log.Fatalf("error preparing stmtSelectVerified: %v", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		remoteIP := s.resolveRemoteIP(r)

		apiKey := r.FormValue("apikey")
		if apiKey != s.options.APIKey {
			log.Printf("invalid api key from %s", remoteIP)
			http.Error(w, "invalid key", http.StatusUnauthorized)
			return
		}

		var data StatsOutput

		err = stmtSelectTotal.QueryRow().Scan(&data.Total)
		if err != nil {
			log.Printf("error selecting total handtekeningen: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		err = stmtSelectVerified.QueryRow().Scan(&data.Verified)
		if err != nil {
			log.Printf("error selecting verified handtekeningen: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(&data)
		if err != nil {
			log.Printf("error encoding api/stats output data: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}
}
