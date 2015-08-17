package main

import (
	"fmt"
	"log"

	"github.com/GeenPeil/teken/storage"
)

func main() {
	parseFlags()

	fetcher, err := storage.NewFetcher(flags.StoragePrivkeyFile, flags.StorageLocation)
	if err != nil {
		log.Fatalf("error creating new fetcher: %v\n", err)
		return
	}

	pdf := NewPDF("hi")
	var pdfFilename string
	var handtekeningenIDList []uint64

	if flags.Single != nil {
		pdfFilename = fmt.Sprintf("single-%d.pdf", *flags.Single)
		handtekeningenIDList = []uint64{*flags.Single}
	} else if flags.Partition != nil {
		pdfFilename = fmt.Sprintf("partition-%d.pdf", *flags.Partition)
		handtekeningenIDList, err = fetcher.ListPartition(*flags.Partition)
		if err != nil {
			log.Fatalf("error listing handtekeningen for partition %d: %v", *flags.Partition, err)
		}
	}

	for _, ID := range handtekeningenIDList {
		h, err := fetcher.Fetch(ID)
		if err != nil {
			log.Printf("error fetching: %v\n", err)
			continue
		}
		err = pdf.AddHandtekening(h)
		if err != nil {
			log.Printf("error rendering handtekening %d: %v", ID, err)
			continue
		}
	}

	err = pdf.Render(pdfFilename)
	if err != nil {
		log.Fatalf("error rendering final pdf: %v", err)
	}

	log.Printf("pdf `%s` successfully rendered\n", pdfFilename)
	if len(handtekeningenIDList) != 1000 {
		log.Printf("WARNING!! partitie had slechts %d handtekeningen, i.p.v. 1000. Gaat alles goed? Neem bij twijfel contact op met Geert-Johan.", len(handtekeningenIDList))
	}
}

func verbosef(format string, args ...interface{}) {
	if flags.Verbose {
		log.Printf(format, args...)
	}
}
