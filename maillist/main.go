package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/huandu/xstrings"

	"github.com/GeenPeil/teken/storage"
)

func main() {
	parseFlags()

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	var fixMail bool
	var fixMailMap map[string]string
	if flags.FixMail != "" {
		fixMail = true
		log.Printf("WARNING Only writing handtekeningen to output.csv which are present in the fixmail csv file")
		fixMailMap = make(map[string]string)
		fixMailFile, err := os.Open(flags.FixMail)
		if err != nil {
			log.Fatalf("error opening fixmail file: %v", err)
		}
		fixMailReader := csv.NewReader(fixMailFile)
		fixMailReader.ReuseRecord = true
		var line uint64
		for {
			line++
			record, err := fixMailReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("error reading from fixmail csv: %v", err)
			}
			if len(record) != 2 {
				log.Fatalf("invalid length record (expect 2 fields) in fixmail csv at line %d", line)
			}
			fixMailMap[record[0]] = record[1]
		}
	}

	fetcher, err := storage.NewFetcher(flags.StoragePrivkeyFile, flags.StorageLocation)
	if err != nil {
		log.Fatalf("error creating new fetcher: %v\n", err)
		return
	}

	chIDs := make(chan uint64)
	go func() {
		if flags.Single != nil {
			chIDs <- *flags.Single
			close(chIDs)
		} else if flags.Partition != nil {
			handtekeningenIDList, err := fetcher.ListPartition(*flags.Partition)
			if err != nil {
				log.Fatalf("error listing handtekeningen for partition %d: %v", *flags.Partition, err)
			}
			for _, id := range handtekeningenIDList {
				chIDs <- id
			}
			close(chIDs)
		} else if flags.All || flags.StartID != nil {
			var startID = uint64(1)
			if flags.StartID != nil {
				startID = *flags.StartID
			}
			startPartition := (startID-1)/1000 + 1
			for i := startPartition; true; i++ {
				fmt.Printf("walking partition %04d\n", i)
				handtekeningenIDList, err := fetcher.ListPartition(i)
				if err != nil {
					log.Printf("error listing handtekeningen for partition %d: %v", i, err)
					break
				}
				for _, id := range handtekeningenIDList {
					if id < startID {
						continue
					}
					chIDs <- id
				}
			}
			close(chIDs)
		}
	}()

	var spreadsheetLock sync.Mutex
	spreadsheetFile, err := os.Create(flags.CSVFile)
	if err != nil {
		log.Fatalf("error opening csv file: %v", err)
	}
	spreadsheet := csv.NewWriter(spreadsheetFile)
	spreadsheet.Write([]string{"ID", "Voornaam", "Achternaam", "Email"})

	var numWorkers = (runtime.NumCPU() - 1) * 2
	wgDone := sync.WaitGroup{}
	wgDone.Add(numWorkers)
	for n := 0; n < numWorkers; n++ {
		go func(n int) {
			for {
				id, ok := <-chIDs
				if !ok {
					break
				}
				h, err := fetcher.Fetch(id)
				if err != nil {
					log.Printf("error fetching %06d\n", id)
					continue
				}

				// Combine tussenvoegsel and achternaam into single field
				var achternaam string
				if h.Tussenvoegsel != "" {
					achternaam = xstrings.FirstRuneToUpper(strings.ToLower(h.Tussenvoegsel)) + " " + xstrings.FirstRuneToUpper(strings.ToLower(h.Achternaam))
				} else {
					achternaam = xstrings.FirstRuneToUpper(strings.ToLower(h.Achternaam))
				}

				// --fixmail procedure
				var email = h.Email
				email = strings.TrimSpace(h.Email)
				email = strings.ToLower(h.Email)
				if fixMail {
					correctedMail, needsFix := fixMailMap[email]
					if !needsFix {
						continue // skip this handtekening, doesn't need fix
					}
					email = correctedMail
				}

				// write field to output csv file
				fields := []string{strconv.FormatUint(id, 10), xstrings.FirstRuneToUpper(strings.ToLower(h.Voornaam)), achternaam, email}
				spreadsheetLock.Lock()
				err = spreadsheet.Write(fields)
				if err != nil {
					log.Printf("error writing fields for %06d: %v\n", id, err)
				}
				spreadsheetLock.Unlock()
			}

			fmt.Printf("worker %02d done\n", n)
			wgDone.Done()
		}(n)
	}
	wgDone.Wait()

	spreadsheet.Flush()
	err = spreadsheetFile.Close()
	if err != nil {
		log.Printf("error closing csv file: %v", err)
	}
}
