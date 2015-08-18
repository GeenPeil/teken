package main

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"

	"github.com/GeenPeil/teken/data"
	"github.com/GeertJohan/go.incremental"
	"github.com/GeertJohan/go.rice"
)

var (
	pngtp string
	jpgtp string
)

var (
	formName  = "form.jpg"
	formBytes []byte
)

func init() {
	p := gofpdf.New("P", "mm", "A4", "")
	pngtp = p.ImageTypeFromMime("image/png")
	jpgtp = p.ImageTypeFromMime("image/jpeg")

	formBytes = rice.MustFindBox("forms").MustBytes("definitief.jpg")
}

type pdf struct {
	name string
	fpdf *gofpdf.Fpdf
	inc  incremental.Uint64
}

func NewPDF(name string) *pdf {
	p := &pdf{
		name: name,
		fpdf: gofpdf.New("P", "pt", "A4", ""),
	}
	p.fpdf.RegisterImageReader(formName, jpgtp, bytes.NewBuffer(formBytes))
	return p
}

func (p *pdf) AddHandtekening(h *data.Handtekening) error {
	p.fpdf.AddPage()
	pageSizeX, pageSizeY := p.fpdf.GetPageSize()
	p.fpdf.Image(formName, 0, 0, pageSizeX, pageSizeY, false, jpgtp, 0, "")
	p.fpdf.SetFont("Arial", "", 14)
	// p.fpdf.Image(imageFile("logo.png"), 10, 10, 30, 0, false, "", 0, "")
	p.FormText(37, 393, h.Voornaam)
	p.FormText(305, 393, h.Tussenvoegsel)
	p.FormText(37, 428, h.Achternaam)
	p.FormText(37, 466, h.Straat)
	p.FormText(423, 466, h.Huisnummer)

	p.FormText(37, 500, h.Postcode[:4])
	var postcodeCijfers string
	if len(h.Postcode) == 6 {
		postcodeCijfers = h.Postcode[4:]
	} else if len(h.Postcode) == 7 {
		postcodeCijfers = h.Postcode[5:]
	}
	p.FormText(105, 500, postcodeCijfers)

	p.FormText(187, 500, h.Woonplaats)
	p.FormText(37, 536, h.Geboortedatum[:2])
	p.FormText(76, 536, h.Geboortedatum[3:5])
	p.FormText(116, 536, h.Geboortedatum[6:10])
	p.FormText(187, 536, h.Geboorteplaats)

	imgName := fmt.Sprintf("h-%d", p.inc.Next())
	p.fpdf.RegisterImageReader(imgName, pngtp, bytes.NewBuffer(h.Handtekening))
	sx := 426.0
	sy := 399.0
	ex := 550.0 - sx
	ey := ex / 2.25
	p.fpdf.Image(imgName, sx, sy, ex, ey, false, pngtp, 0, "")

	return nil
}

func (p *pdf) FormText(x, y float64, text string) {
	for _, c := range text {
		// p.fpdf.Text(x, y, string(c))
		p.fpdf.MoveTo(x, y)
		p.fpdf.CellFormat(12, 20, string(c), "", 0, "CA", false, 0, "")
		x += 14.8
	}
}

func (p *pdf) Render(filename string) error {
	return p.fpdf.OutputFileAndClose(filename)
}
