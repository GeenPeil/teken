package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/GeenPeil/teken/storage"
)

func main() {
	parseFlags()

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

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
	spreadsheet.Write([]string{"ID", "Voornaam", "Tussenvoegsel", "Achternaam", "Email"})

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
				fields := []string{strconv.FormatUint(id, 10), h.Voornaam, h.Tussenvoegsel, h.Achternaam, h.Email}
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
