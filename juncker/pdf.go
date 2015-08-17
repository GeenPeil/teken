package main

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"

	"github.com/GeenPeil/teken/data"
	"github.com/GeertJohan/go.incremental"
)

var pngtp string

func init() {
	p := gofpdf.New("P", "mm", "A4", "")
	pngtp = p.ImageTypeFromMime("image/png")
}

type pdf struct {
	name string
	fpdf *gofpdf.Fpdf
	inc  incremental.Uint64
}

func NewPDF(name string) *pdf {
	return &pdf{
		name: name,
		fpdf: gofpdf.New("P", "mm", "A4", ""),
	}
}

func (p *pdf) AddHandtekening(h *data.Handtekening) error {
	p.fpdf.AddPage()
	p.fpdf.SetFont("Arial", "", 11)
	// p.fpdf.Image(imageFile("logo.png"), 10, 10, 30, 0, false, "", 0, "")
	p.fpdf.Text(50, 20, h.Voornaam)
	p.fpdf.Text(50, 30, h.Tussenvoegsel)
	p.fpdf.Text(50, 40, h.Achternaam)
	p.fpdf.Text(50, 50, h.Geboortedatum)
	p.fpdf.Text(50, 60, h.Geboorteplaats)
	p.fpdf.Text(50, 70, h.Straat)
	p.fpdf.Text(50, 80, h.Huisnummer)
	p.fpdf.Text(50, 90, h.Postcode)
	p.fpdf.Text(50, 100, h.Woonplaats)
	imgName := fmt.Sprintf("h-%d", p.inc.Next())
	p.fpdf.RegisterImageReader(imgName, pngtp, bytes.NewBuffer(h.Handtekening))
	p.fpdf.Image(imgName, 50, 110, 0, 0, false, pngtp, 0, "")

	return nil
}

func (p *pdf) Render(filename string) error {
	return p.fpdf.OutputFileAndClose(filename)
}
