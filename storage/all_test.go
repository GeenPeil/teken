package storage

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/GeenPeil/teken/data"
)

func TestStorage(t *testing.T) {
	h := &data.Handtekening{
		Voornaam:        "Voornaam",
		Tussenvoegsel:   "Tussenvoegsel",
		Achternaam:      "Achternaam",
		Geboortedatum:   "Geboortedatum",
		Geboorteplaats:  "Geboorteplaats",
		Straat:          "Straat",
		Huisnummer:      "Huisnummer",
		Postcode:        "Postcode",
		Woonplaats:      "Woonplaats",
		Handtekening:    []byte("Handtekening"),
		CaptchaResponse: "foobar",
	}

	saver, err := NewSaver("testpub.pem", "testdata")
	if err != nil {
		fmt.Printf("error creating new saver: %v\n", err)
		return
	}

	err = saver.Save(1, h)
	if err != nil {
		fmt.Printf("error saving: %v\n", err)
		return
	}

	fetcher, err := NewFetcher("testkey.pem", "testdata")
	if err != nil {
		fmt.Printf("error creating new fetcher: %v\n", err)
		return
	}

	h2, err := fetcher.Fetch(1)
	if err != nil {
		fmt.Printf("error fetching: %v\n", err)
		return
	}

	if !reflect.DeepEqual(h, h2) {
		t.Error("saved handtekening doesn't match fetched handtekening")
	}

	if h2.CaptchaResponse != "" {
		t.Error("CaptchaResponse on fetched handtekening is not empty")
	}
}
