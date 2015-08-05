package server

//go:generate ffjson -nodecoder $GOFILE

type uploadOutput struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
