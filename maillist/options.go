package main

import (
	"fmt"
	"os"

	goflags "github.com/jessevdk/go-flags"
)

var flags struct {
	Verbose bool `long:"verbose" short:"v" description:"Show verbose debug information"`

	Single    *uint64 `short:"s" long:"single" description:"Single handtekening to render"`
	Partition *uint64 `short:"p" long:"partitie" description:"Partition of handtekeningen to render"`
	All       bool    `short:"a" long:"all" description:"Get all mail addresses"`
	StartID   *uint64 `long:"start-id" description:"Starting id for the scanning"`

	StoragePrivkeyFile string `long:"storage-privkey-file" description:"storage private key" default:"../storage/testpriv.pem"` // TODO: remove default, make mandatory
	StorageLocation    string `long:"storage-location" desciption:"storage location" default:"../storage/testdata"`             // TODO: remove default, make mandatory

	CSVFile string `long:"csv-file" description:"csv file location" default:"output.csv"` // TODO: remove default, make mandatory
}

func parseFlags() {
	parser := goflags.NewParser(&flags, goflags.HelpFlag|goflags.PrintErrors)

	// parse flags
	args, err := parser.Parse()
	if err != nil {
		// assert the err to be a flags.Error
		flagError, ok := err.(*goflags.Error)
		if !ok {
			// not a flags error
			os.Exit(1)
		}
		if flagError.Type == goflags.ErrHelp {
			// exitcode 0 when user asked for help
			fmt.Println("Read the README.md for more usage information.")
			os.Exit(0)
		}
		// error on left-over arguments
		if len(args) > 0 {
			fmt.Printf("unexpected arguments: %s\n", args)
			os.Exit(0)
		}
		// error on left-over arguments
		if len(args) > 0 {
			fmt.Printf("unexpected arguments: %s\n", args)
			os.Exit(0)
		}
		if flagError.Type == goflags.ErrUnknownFlag {
			fmt.Println("run with --help to view available options")
		}
		os.Exit(1)
	}

	// error on left-over arguments
	if len(args) > 0 {
		fmt.Printf("unexpected arguments: %s\n", args)
		os.Exit(0)
	}

	// required flags
	if !flags.All && flags.Single == nil && flags.Partition == nil && flags.StartID == nil {
		fmt.Println("Require either --single, --partition, --all or --start-id flag to render a single or partition of handtekeningen to pdf.")
		os.Exit(42)
	}
}
