package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/GeenPeil/teken/data"
)

var handtekeningPNGBytes, _ = base64.StdEncoding.DecodeString(`iVBORw0KGgoAAAANSUhEUgAAACkAAAAjCAYAAAAJ+yOQAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH1woJFjko180rngAAABl0RVh0Q29tbWVudABDcmVhdGVkIHdpdGggR0lNUFeBDhcAAAUVSURBVFjDzZh9TJVVHMc/57mXN8EQuwoipqQivg1XpkmbLzO1mGtqkHNKU6alk2UUpqmoga+TWYalmEx6cYAmPpXEqNTZmrPNzDAtjbIsFKepiYBXLs/pj3Mvl5d74Soi97edPb/znO95nu++z/melwfuc0idOVJnON4at/JJXPgsMvcVDKkT63UEpU6XzUlYAQnIL9O4JHWEV5G88hG5IYGKICD790BW72GJN6k4dMEzityEGGRUuMrXzaJW6li8gmTpFk6ZNKQmkL9sRR5Zq0h28kWWbaPYG1Sc9lS0IrVoMlLqqiTEqnsJsUipM6YjCfrnp1IJSEtnZHWBk+TfO5G+ZkW0aEXbTKS1heTNapYtziUIIGMmBPg52yIskD5D5ck7CKusYWlHqBi+ZiZ1gBwYgbTtc6roKHc+RfYNa2Si4AeqZNklsjMKVP/tC8Bkao7xMcMHC1W+Oh/z2XL2PDCSUmdEegGTrTZ4fhSMHuxsO38ZSk7ArRpVHzcUEmLhjg3SdjPxgZhI6ohDGfwuQJo15Pls5+fdv1TdA2S/MOSN3c1NdGAFFVLHp11J2gqZPSxSvXB5fOMx6JjEHSVztrNtQ6KTfLuuRFInKCeZKkCGBCIr8xqTdKjoKPMnOdvqCpHRPetNZJM6PdplTF6/RfryT+gEsGUuBAU0bo+wuK9rGmS9pPK1ezGdv8zO+05S6jy6sZBFFTdgRH9IHNccE9On5frTMcpEVVZY8iFxnm7nPCb5RwU7tnyBJoCsea4xw/s1rj/RvzlmcxIE+cPeo1B0nEJPTKR5qOLoxbmMv10LiWNhRJRr3JMDnHmf7hDapTkmwgIrElSeuovQyhpS20xS6phKfiS/8JhSYP2L7rEjo1BSA6MGuMelPAfRPeHXcthaRLrUCW8Tydu1LHgjVzlxWTyEd3WP7RwAQx5ReWy0e5yvj1ql7CYyny1n1z2TlDohOV+TWfoXRHZXCrQWjnHpajw2jDFDnCayr0Rj74nkxWtkrs7DDyBzDvj7tk5yWCQInIq2FJuTINBPmehgKXlSd81Ha0HFfpv2k3S1EsYOgWmjPJsFoiMgLAQC/VvHRlgg7QX7di6bsFobr7vCiRaOBCcfSyFGqpxBvdppma2DgQuhrALWzcL2ZjyhYgrXWlVS6sSl5BBjM2DuBM8JXroG8RvVS1/LgVpb633MJtiZrPJVeZjPlpPXFGNyQdD82fccXb+PgC6BULwS/Dzcs0xdD0U/wNVKOHYOfEzKIK1Fn+5w5gKcugCXb9D3dBYlb+Xzj1sla6y8mrqLrgArp6tpxdP47kzj+renPe/7zlwI8FUmOnyKwoYm0pqoaNlewoayCuhlgeS4uxtfQ3u3XG8pwrvCqukqn7+NHtZaFrskefUmmzIK1BDYNl9t/+8mPk6Bx/uCvw9MHQmrZ9xd/9QpEBUO5y7C25+zzrESiQYqDk7O5uf3itWUc3hNxxyTD5XC+JVq/jyzlaLe85hcb5wpIyl5+X21/BWlQbfgjiEZGapMdPJPKP+XqNNZfCPsKk6MS6ek+AQ8FKCWNMOwb68lGFJdpVT3DKNJjrPdcINztEHzZzd9RrUVrlcp7JG1/GYGqLYy6auT9gN/DRws9Y4fYT4mCPClt0PJQQeO89OV/6h6uDPXBUjNfl4RAkMTSCGUOEKoH1PCngtRf5X2foZw9jNEAywCQ4Bh0urxhqZRpwnqGjzDEII6k4YM7kR1t2De/R8+QVLKM5yBtQAAAABJRU5ErkJggg==`)

func TestServer(t *testing.T) {
	// start server async but wait for most of the setup to be done
	o := &Options{
		HTTPAddress:            ":8080",
		CaptchaDisable:         true,
		StoragePubkeyFile:      "../../storage/testpub.pem",
		StorageLocation:        "../../storage/testdata",
		PostgresSocketLocation: "/var/run/postgresql",
		HashingSalt:            "FWIrJHXXxb2+iGRluUz/sjS3X1zI3LBGTQJTsYcqu9I=",
	}
	s := New(o)
	setupDoneCh := make(chan struct{})
	go s.Run(setupDoneCh)
	<-setupDoneCh

	handtekening := &data.Handtekening{
		Voornaam:        "Voornaam",
		Tussenvoegsel:   "Tussenvoegsel",
		Achternaam:      "Achternaam",
		Geboortedatum:   "Geboortedatum",
		Geboorteplaats:  "Geboorteplaats",
		Straat:          "Straat",
		Huisnummer:      "Huisnummer",
		Postcode:        "Postcode",
		Woonplaats:      "Woonplaats",
		Handtekening:    handtekeningPNGBytes,
		CaptchaResponse: "foobar",
		Email:           "gjr19912@gmail.com",
	}
	handtekeningJSON, err := json.MarshalIndent(handtekening, "", "\t")
	if err != nil {
		t.Fatalf("error creating test JSON data: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/pechtold/submit", "application/json", bytes.NewBuffer(handtekeningJSON))
	if err != nil {
		t.Fatalf("error making upload request: %v", err)
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		t.Fatalf("http request returned non-200: %d", resp.StatusCode)
	}

	out := &SubmitOutput{}
	err = json.NewDecoder(resp.Body).Decode(out)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("invalid json in response body: %v", err)
	}

	if !out.Success {
		t.Fatalf("error in request data: %s", out.Error)
	}

	// all ok
}
