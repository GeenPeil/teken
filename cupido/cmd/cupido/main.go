package main

import (
	"fmt"
	"os"

	"github.com/GeenPeil/teken/cupido"

	goflags "github.com/jessevdk/go-flags"
)

func main() {

	flags := &cupido.Options{}

	parser := goflags.NewParser(flags, goflags.HelpFlag|goflags.PrintErrors)

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

	s := cupido.New(flags)
	s.Run(nil)
}
