package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/GeenPeil/teken/data"
	"github.com/GeenPeil/teken/pechtold/server"
)

const workers = 800
const posts = 500

func main() {

	var handtekeningPNGBytes, err = base64.StdEncoding.DecodeString(`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`)
	if err != nil {
		panic(err)
	}

	var wgSetup sync.WaitGroup
	var wgDone sync.WaitGroup
	wgSetup.Add(workers)
	wgDone.Add(workers)
	startCh := make(chan struct{})

	for n := 0; n < workers; n++ {
		go func(n int) {
			handtekening := &data.Handtekening{
				Voornaam:       fmt.Sprintf("Voornaam %d", n),
				Tussenvoegsel:  "Tussenvoegsel",
				Achternaam:     "",
				Geboortedatum:  "Geboortedatum",
				Geboorteplaats: "Geboorteplaats",
				Straat:         "Straat",
				Huisnummer:     "Huisnummer",
				Postcode:       "Postcode",
				Woonplaats:     "Woonplaats",
				Handtekening:   handtekeningPNGBytes,
			}

			// setup done, wait for start signal
			wgSetup.Done()
			<-startCh

			for i := 0; i < posts; i++ {
				handtekening.Achternaam = fmt.Sprintf("Achternaam %d", i)
				handtekeningJSON, err := json.MarshalIndent(handtekening, "", "\t")
				if err != nil {
					log.Fatalf("error creating test JSON data: %v", err)
				}

				resp, err := http.Post("http://localhost:8080/pechtold/submit", "application/json", bytes.NewBuffer(handtekeningJSON))
				if err != nil {
					log.Fatalf("error making upload request: %v", err)
				}
				if resp.StatusCode != 200 {
					resp.Body.Close()
					log.Fatalf("http request returned non-200: %d", resp.StatusCode)
				}

				out := &server.SubmitOutput{}
				err = json.NewDecoder(resp.Body).Decode(out)
				resp.Body.Close()
				if err != nil {
					log.Fatalf("invalid json in response body: %v", err)
				}

				if !out.Success {
					log.Fatalf("error in request data: %s", out.Error)
				}

				fmt.Printf("done %03d-%04d\n", n, i)
			}

			wgDone.Done()
		}(n)
	}

	// wait until all goroutines are started
	wgSetup.Wait()
	tStart := time.Now()
	// signal start
	close(startCh)
	wgDone.Wait()
	tEnd := time.Now()

	fmt.Printf("completed in %s\n", tEnd.Sub(tStart).String())
}
