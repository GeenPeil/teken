package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/GeenPeil/teken/data"
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
		for i := uint64(1); true; i++ {
			fmt.Printf("walking partition %04d\n", i)
			handtekeningenIDList, err := fetcher.ListPartition(i)
			if err != nil {
				log.Printf("error listing handtekeningen for partition %d: %v", i, err)
				break
			}
			for _, id := range handtekeningenIDList {
				chIDs <- id
			}
		}
		close(chIDs)
	}()

	type checklist struct {
		lock sync.Mutex
		list map[string]uint
	}
	var checks []*checklist
	for i := 0; i < 256; i++ {
		checks = append(checks, &checklist{
			list: make(map[string]uint),
		})
	}

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
				hash := hashHandtekening(h)
				cl := checks[int(hash[0])]
				cl.lock.Lock()
				cl.list[hash]++
				cl.lock.Unlock()
			}

			fmt.Printf("worker %02d done\n", n)
			wgDone.Done()
		}(n)
	}
	wgDone.Wait()

	var total uint
	for i := 0; i < 256; i++ {
		cl := checks[i]
		for _, count := range cl.list {
			if count > 1 {
				total += count
				fmt.Printf("%d, ", count)
			}
		}
	}
	fmt.Println("")
	fmt.Println(total)
}

func hashHandtekening(h *data.Handtekening) string {
	nawHash := sha256.New()
	nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Voornaam))))
	nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Achternaam))))
	nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Geboortedatum))))
	nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Geboorteplaats))))
	nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Straat))))
	nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Huisnummer))))
	nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Postcode))))
	nawHash.Write(bytes.ToLower(bytes.TrimSpace([]byte(h.Woonplaats))))
	nawHashBytes := nawHash.Sum(nil)
	return string(nawHashBytes)
}
