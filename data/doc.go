// Package data contains the Handtekening struct which can be unmarshalled from json (http api) and marshalled with proto (storage).
package data

//go:generate protoc --go_out=. handtekening.proto
//go:generate ffjson -noencoder handtekening.pb.go

//go:generate protoc --go_out=. gph.proto
