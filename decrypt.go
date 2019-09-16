package rails5cook

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	keyIterNum = 1000
	keySize    = 32
)

func (cookie *Cookie) Decrypt() {
	cookie.Decode()

	key := cookie.GenerateKey()
	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, cookie.Iv, cookie.Data, nil)

	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(plaintext, &cookie.Content); err != nil {
		panic(err)
	}
}

func (cookie *Cookie) Decode() {

	encryptedData := strings.Split(cookie.Value, "--")

	data, _ := base64.StdEncoding.DecodeString(encryptedData[0])
	iv, _ := base64.StdEncoding.DecodeString(encryptedData[1])
	authTag, _ := base64.StdEncoding.DecodeString(encryptedData[2])

	cookie.Data = []byte(append(data, authTag...))
	cookie.Iv = []byte(iv)
}

func (cookie *Cookie) GenerateKey() []byte {
	return pbkdf2.Key([]byte(cookie.SecretKeyBase), []byte(cookie.Salt), keyIterNum, keySize, sha1.New)
}
