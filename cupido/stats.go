package cupido

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

func (s *Server) newAPIStatsHandlerFunc() http.HandlerFunc {

	stmtSelectTotal, err := s.db.Prepare(`SELECT COUNT(*) FROM handtekeningen`)
	if err != nil {
		log.Fatalf("error preparing stmtSelectTotal: %v", err)
	}

	// stmtSelectVerified, err := s.db.Prepare(`SELECT COUNT(*) FROM handtekeningen WHERE mailcheckdone = true`)
	// if err != nil {
	// 	log.Fatalf("error preparing stmtSelectVerified: %v", err)
	// }

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

		// err = stmtSelectVerified.QueryRow().Scan(&data.Verified)
		// if err != nil {
		// 	log.Printf("error selecting verified handtekeningen: %v", err)
		// 	http.Error(w, "server error", http.StatusInternalServerError)
		// 	return
		// }

		err = json.NewEncoder(w).Encode(&data)
		if err != nil {
			log.Printf("error encoding api/stats output data: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) newAPIPublicStatsHandlerFunc(offset, factor float64) http.HandlerFunc {

	stmtSelectTotal, err := s.db.Prepare(`SELECT COUNT(*) FROM handtekeningen`)
	if err != nil {
		log.Fatalf("error preparing stmtSelectTotal: %v", err)
	}

	var outputBytes = make([]byte, 0) // load empty slice to avoid nilpointer exception when goroutine hasn't filled this yet.
	var outputBytesLock sync.RWMutex

	go func() {
		for {
			var data StatsOutput
			err = stmtSelectTotal.QueryRow().Scan(&data.Total)
			if err != nil {
				log.Printf("error selecting total handtekeningen: %v", err)
			}

			// account for invalid signatures
			data.Total = uint64(offset + (float64(data.Total)-offset)*factor)

			var buf = bytes.Buffer{}
			err = json.NewEncoder(&buf).Encode(&data)
			if err != nil {
				log.Printf("error encoding api/stats output data: %v", err)
			}
			outputBytesLock.Lock()
			outputBytes = buf.Bytes()
			outputBytesLock.Unlock()
			time.Sleep(900 * time.Millisecond)
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		outputBytesLock.RLock()
		w.Write(outputBytes)
		outputBytesLock.RUnlock()
	}
}
