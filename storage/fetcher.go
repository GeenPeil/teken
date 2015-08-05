package storage

import (
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/GeenPeil/teken/data"
	"github.com/gogo/protobuf/proto"
)

// Fetcher is responsible for fetching and decrypting+decoding stored handtekeningen.
type Fetcher struct {
	privkey  *rsa.PrivateKey
	datapath string
}

// NewFetcher creates a new *Fetcher instance with private key from given privkeyFilename.
// The returned Fetcher will use datapath to lookup stored files.
func NewFetcher(privkeyFilename string, datapath string) (*Fetcher, error) {
	privkeyFile, err := os.Open(privkeyFilename)
	if err != nil {
		return nil, err
	}
	defer privkeyFile.Close()

	privkeyPem, err := ioutil.ReadAll(privkeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(privkeyPem))
	rsaPriv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %v", err)
	}

	f := &Fetcher{
		privkey:  rsaPriv,
		datapath: datapath,
	}
	return f, nil
}

// Fetch loads a handtekening from file by number
func (f *Fetcher) Fetch(n int) (*data.Handtekening, error) {
	filename, foldername := fileFolderByNumber(n)
	file, err := os.Open(filepath.Join(f.datapath, foldername, filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	encryptedData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	hash := sha1.New()

	bufData, err := rsa.DecryptOAEP(hash, nil, f.privkey, encryptedData, nil)
	if err != nil {
		return nil, err
	}

	h := &data.Handtekening{}

	err = proto.Unmarshal(bufData, h)
	if err != nil {
		return nil, err
	}

	return h, nil
}
