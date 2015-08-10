package storage

import (
	"crypto/aes"
	"crypto/cipher"
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

const aesBlockSize = 256 / 8

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
func (s *Saver) Save(n uint64, h *data.Handtekening) error {
	// empty captcha response, we're not saving it.
	h.CaptchaResponse = ""

	// gpa protobuf is saved to disk
	gph := data.GPH{}

	// TODO: don't return errors from crypto/*, instead: print them and return ErrCryptoError

	// marshall handtekening to proto buf
	hBuf, err := proto.Marshal(h)
	if err != nil {
		return err
	}

	// read new aes key
	aesKey := make([]byte, aesBlockSize) // TODO: use sync.Pool
	_, err = rand.Read(aesKey)
	if err != nil {
		return err
	}
	// create AES cipher.Block
	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}

	// create new CFB encrypter and encrypt hbuf to itself (re-using allocation)
	cfbEncrypter := cipher.NewCFBEncrypter(aesBlock, ivFromNumber(n))
	cfbEncrypter.XORKeyStream(hBuf, hBuf)
	gph.AESEncryptedData = hBuf

	// use sha1 as hash in rsa
	rsaHasher := sha1.New() // TODO: use sync.Pool

	gph.RSAEncryptedAESKey, err = rsa.EncryptOAEP(rsaHasher, rand.Reader, s.pubkey, aesKey, nil)
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
	gphBuf, err := proto.Marshal(&gph)
	if err != nil {
		return err
	}
	_, err = file.Write(gphBuf)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
