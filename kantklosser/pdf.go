package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"

	"github.com/GeenPeil/teken/data"
	incremental "github.com/GeertJohan/go.incremental"
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

const (
	firstOffsetY  = -2
	secondOffsetY = 217
)

func init() {
	p := gofpdf.New("P", "mm", "A4", "")
	pngtp = p.ImageTypeFromMime("image/png")
	jpgtp = p.ImageTypeFromMime("image/jpeg")

	formBytes = rice.MustFindBox("forms").MustBytes("33506-definitief-vl-nl-1.jpg")
}

type pdf struct {
	name      string
	fpdf      *gofpdf.Fpdf
	translate func(in string) string

	inc incremental.Uint64
}

func NewPDF(name string) *pdf {
	p := &pdf{
		name: name,
		fpdf: gofpdf.New("P", "pt", "A4", ""),
	}
	p.translate = p.fpdf.UnicodeTranslatorFromDescriptor("")
	p.fpdf.RegisterImageReader(formName, jpgtp, bytes.NewBuffer(formBytes))
	return p
}

var oldDate = time.Date(1900, 01, 01, 01, 01, 01, 01, time.UTC)
var dateFormat = `02-01-2006`

func (p *pdf) AddHandtekening(h *data.Handtekening) error {
	if !flags.SkipAgeCheck {
		date, err := time.Parse(dateFormat, h.Geboortedatum)
		if err != nil {
			var errPlausibleDate error
			date, errPlausibleDate = time.Parse(dateFormat, h.Geboortedatum[3:5]+"-"+h.Geboortedatum[0:2]+"-"+h.Geboortedatum[6:])
			if errPlausibleDate != nil {
				fmt.Printf("Skipping invalid birthdate %s: %v\n", h.Geboortedatum, err)
				return nil
			}
			fmt.Printf("Graceful on date %s, accepted as %s\n", h.Geboortedatum, date.Format(dateFormat))
		}
		if date.Before(oldDate) {
			fmt.Printf("Skipping old birthdate %s (%s %s %s)\n", h.Geboortedatum, h.Voornaam, h.Achternaam, h.Email)
			return nil
		}
	}

	pagePos := p.inc.Last() % 2
	if pagePos == 0 {
		p.fpdf.AddPage()
		pageSizeX, pageSizeY := p.fpdf.GetPageSize()
		p.fpdf.Image(formName, 0, 0, pageSizeX, pageSizeY, false, jpgtp, 0, "")
		p.fpdf.SetFont("Arial", "", 14)
	}
	offsetY := float64(firstOffsetY + (secondOffsetY * int(pagePos)))
	// p.fpdf.Image(imageFile("logo.png"), 10, 10, 30, 0, false, "", 0, "")
	p.FormText(37, offsetY+393, h.Voornaam)
	p.FormText(305, offsetY+393, h.Tussenvoegsel)
	p.FormText(37, offsetY+428, h.Achternaam)
	p.FormText(37, offsetY+466, h.Straat)
	p.FormText(423, offsetY+466, h.Huisnummer)

	p.FormText(37, offsetY+500, h.Postcode[:4])
	var postcodeCijfers string
	if len(h.Postcode) == 6 {
		postcodeCijfers = h.Postcode[4:]
	} else if len(h.Postcode) == 7 {
		postcodeCijfers = h.Postcode[5:]
	}
	p.FormText(105, offsetY+500, postcodeCijfers)

	p.FormText(187, offsetY+500, h.Woonplaats)
	p.FormText(37, offsetY+536, h.Geboortedatum[:2])
	p.FormText(76, offsetY+536, h.Geboortedatum[3:5])
	p.FormText(116, offsetY+536, h.Geboortedatum[6:10])
	p.FormText(187, offsetY+536, h.Geboorteplaats)

	imgName := fmt.Sprintf("h-%d", p.inc.Next())
	p.fpdf.RegisterImageReader(imgName, pngtp, bytes.NewBuffer(h.Handtekening))
	sx := 426.0
	sy := 400.0 + offsetY
	ex := 550.0 - sx
	ey := ex / 2.25
	p.fpdf.Image(imgName, sx, sy, ex, ey, false, pngtp, 0, "")

	return nil
}

func (p *pdf) FormText(x, y float64, text string) {
	for _, c := range text {
		// p.fpdf.Text(x, y, string(c))
		p.fpdf.MoveTo(x, y)
		p.fpdf.CellFormat(12, 20, p.translate(string(c)), "", 0, "CA", false, 0, "")
		x += 14.8
	}
}

func (p *pdf) Render(filename string) error {
	return p.fpdf.OutputFileAndClose(filename)
}
