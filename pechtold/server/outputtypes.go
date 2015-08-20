package server

//go:generate ffjson -nodecoder $GOFILE

type SubmitOutput struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type StatsOutput struct {
	Total    uint64 `json:"total"`
	Verified uint64 `json:"verified"`
}
