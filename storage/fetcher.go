package storage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
func (f *Fetcher) Fetch(n uint64) (*data.Handtekening, error) {
	filename, foldername := fileFolderByNumber(n)
	file, err := os.Open(filepath.Join(f.datapath, foldername, filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read gph encoded data
	gphBuf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// unmarshal into GPH structure
	gph := &data.GPH{}
	err = proto.Unmarshal(gphBuf, gph)
	if err != nil {
		return nil, err
	}

	// new sha1 hasher for use with rsa encryption
	rsaHasher := sha1.New() // TODO: use sync.Pool
	aesKey, err := rsa.DecryptOAEP(rsaHasher, nil, f.privkey, gph.RSAEncryptedAESKey, nil)
	if err != nil {
		return nil, err
	}

	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	cfbDecrypter := cipher.NewCFBDecrypter(aesBlock, gph.AESCFBIV)
	cfbDecrypter.XORKeyStream(gph.AESEncryptedData, gph.AESEncryptedData)
	hBuf := gph.AESEncryptedData

	h := &data.Handtekening{}
	err = proto.Unmarshal(hBuf, h)
	if err != nil {
		return nil, err
	}

	return h, nil
}

// ListPartition returns singles that were saved in a partition
func (f *Fetcher) ListPartition(n uint64) ([]uint64, error) {
	idList := make([]uint64, 0, 1000)

	partitionPath := filepath.Join(f.datapath, folderByPartition(n))
	err := filepath.Walk(partitionPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if path == partitionPath {
				return nil
			}
			return filepath.SkipDir
		}

		filename := info.Name()
		if !strings.HasSuffix(filename, ".gph") {
			log.Printf("skipping unrecognized file `%s` in partition %d\n", filename, n)
			return nil
		}

		id, err := strconv.ParseUint(filename[:len(filename)-4], 10, 64)
		if err != nil {
			log.Printf("skipping file %s in partition %d because of bad filename format, expecting <number>.gph\n", filename, n)
			return nil
		}
		idList = append(idList, id)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return idList, nil
}
