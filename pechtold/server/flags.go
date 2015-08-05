package server

import (
	"fmt"
	"log"
	"os"

	goflags "github.com/jessevdk/go-flags"
)

var flags struct {
	Verbose       bool   `long:"verbose" short:"v" description:"Show verbose debug information"`
	HTTPAddress   string `long:"http-address" description:"HTTP address to listen on" default:":8080"`
	CaptchaSecret string `long:"captcha-secret" description:"Google reCaptcha secret" default:"6Lfl0QoTAAAAAFKK76skXuJwlt5x2U_R8Lf7nHLP"` // testing server secret (localhost only), corresponding site key = 6Lfl0QoTAAAAAGA3RbwPfNj2th6gDYLEf0im51RY
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
}

func verbosef(format string, args ...interface{}) {
	if flags.Verbose {
		log.Printf(format, args...)
	}
}
