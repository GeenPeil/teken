package storage

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/GeenPeil/teken/data"
	"github.com/gogo/protobuf/proto"
)

// Saver is responsible for encoding+encrypting uploaded handtekeningen to storage.
type Saver struct {
	pubkey   *rsa.PublicKey
	datapath string
}

// NewSaver returns a new *Saver instance with a public key loaded from the given pubkeyFilename.
// The returned Saver will use datapash to write files to, using the specified directory structure.
func NewSaver(pubkeyFilename string, datapath string) (*Saver, error) {
	pubkeyFile, err := os.Open(pubkeyFilename)
	if err != nil {
		return nil, err
	}
	defer pubkeyFile.Close()

	pubkeyPem, err := ioutil.ReadAll(pubkeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(pubkeyPem))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA public key: %v", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("value returned from ParsePKIXPublicKey was not an RSA public key")
	}

	s := &Saver{
		pubkey:   rsaPub,
		datapath: datapath,
	}
	return s, nil
}

// Save stores the given handtekening in an encrypted file.
// NOTE!: Save clears the CaptchaResponse field.
func (s *Saver) Save(n int, h *data.Handtekening) error {
	h.CaptchaResponse = ""

	bufData, err := proto.Marshal(h)
	if err != nil {
		return err
	}

	hash := sha1.New()

	encryptedData, err := rsa.EncryptOAEP(hash, rand.Reader, s.pubkey, bufData, nil)
	if err != nil {
		return err
	}

	filename, foldername := fileFolderByNumber(n)
	foldername = filepath.Join(s.datapath, foldername)

	err = os.MkdirAll(foldername, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(foldername, filename))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(encryptedData)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
