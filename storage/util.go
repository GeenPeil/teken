package storage

import (
	"crypto/aes"
	"encoding/binary"
	"fmt"
)

// fileFolderByNumber returns the filename and foldername by ID
func fileFolderByNumber(n uint64) (string, string) {
	return fmt.Sprintf("%06d.gph", n), fmt.Sprintf("%04d", n/1000+1)
}

// iv used in AES-CFB encryption
func ivFromNumber(n uint64) []byte {
	iv := make([]byte, aes.BlockSize) // TODO: use sync.Pool
	binary.LittleEndian.PutUint64(iv[:len(iv)-(64/8)], uint64(n))
	return iv
}
