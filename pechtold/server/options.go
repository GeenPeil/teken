package server

// Options is used to create a new server instance.
// It's meant to be used with github.com/jessevdk/go-flags, but is seperated from the actual parsing to allow overrides in tests and benchmarks.
type Options struct {
	Verbose bool `long:"verbose" short:"v" description:"Show verbose debug information"`

	HTTPAddress string `long:"http-address" description:"HTTP address to listen on" default:":8080"`

	CaptchaSecret  string `long:"captcha-secret" description:"Google reCaptcha secret" default:"6Lfl0QoTAAAAAFKK76skXuJwlt5x2U_R8Lf7nHLP"` // TODO: remove default, make mandatory. -- testing server secret (localhost only), corresponding site key = 6Lfl0QoTAAAAAGA3RbwPfNj2th6gDYLEf0im51RY
	CaptchaDisable bool   `long:"captcah-disable" description:"Disable captcha check"`

	StoragePubkeyFile string `long:"storage-pubkey-file" description:"storage public key" default:"../storage/testpub.pem"` // TODO: remove default, make mandatory
	StorageLocation   string `long:"storage-location" desciption:"storage location" default:"../storage/testdata"`          // TODO: remove default, make mandatory
}
