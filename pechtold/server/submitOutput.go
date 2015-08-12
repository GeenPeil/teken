package server

//go:generate ffjson -nodecoder $GOFILE

type submitOutput struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
