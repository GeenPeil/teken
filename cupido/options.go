package cupido

// Options is used to create a new server instance.
// It's meant to be used with github.com/jessevdk/go-flags, but is seperated from the actual parsing to allow overrides in tests and benchmarks.
type Options struct {
	Verbose bool `long:"verbose" env:"VERBOSE" short:"v" description:"Show verbose debug information"`

	HTTPAddress string `long:"http-address" env:"HTTP_ADDRESS" description:"HTTP address to listen on" default:":8002"`

	CaptchaSecret  string `long:"captcha-secret" env:"CAPTCHA_SECRET" description:"Google reCaptcha secret" default:"6Lfl0QoTAAAAAFKK76skXuJwlt5x2U_R8Lf7nHLP"` // TODO: remove default, make mandatory. -- testing server secret (localhost only), corresponding site key = 6Lfl0QoTAAAAAGA3RbwPfNj2th6gDYLEf0im51RY
	CaptchaDisable bool   `long:"captcha-disable" env:"CAPTCHA_DISABLE" description:"Disable captcha check"`

	PostgresSocketLocation string `long:"postgres-socket-location" env:"POSTGRES_SOCKET_LOCATION" description:"PostgreSQL Unix socket location" default:"/var/run/postgresql"`

	StoragePubkeyFile string `long:"storage-pubkey-file" env:"STORAGE_PUBKEY_FILE" description:"storage public key" default:"./storage/testpub.pem"` // TODO: remove default, make mandatory
	StorageLocation   string `long:"storage-location" env:"STORAGE_LOCATION" desciption:"storage location" default:"./storage/testdata"`             // TODO: remove default, make mandatory

	HashingSalt string `long:"hashing-salt" env:"HASHING_SALT" description:"base-64 encoded hashing salt"`

	SMTPHost     string `long:"smtp-host" env:"SMTP_HOST" description:"smtp server to use" default:"localhost:25"`
	SMTPUsername string `long:"smtp-username" env:"SMTP_USERNAME" description:"Username for SMTP auth"`
	SMTPPassword string `long:"smtp-password" env:"SMTP_PASSWORD" description:"Password for SMTP auth"`

	TotalOffset float64 `long:"total-offset" env:"TOTAL_OFFSET" description:"Offset for the totals calculation, to account for invalid signatures."`
	TotalFactor float64 `long:"total-factor" env:"TOTAL_FACTOR" description:"Factor for the totals calculation, to account for invalid signatures."`

	APIKey string `long:"api-key" env:"API_KEY" description:"api authentication key"`
}
