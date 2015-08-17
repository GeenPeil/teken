package server

//go:generate ffjson -nodecoder $GOFILE

type SubmitOutput struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
