package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type Rails5Cookie struct {
	Value         string
	Data          []byte
	Iv            []byte
	SecretKeyBase string
	Salt          string
	Content       *Content
}

type Content struct {
	SessionID string `json:"session_id"`
	Csrf      string `json:"_csrf_token"`
}

const (
	keyIterNum = 1000
	keySize    = 32
)

func (cookie *Rails5Cookie) Init() {
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

func (cookie *Rails5Cookie) Decode() {

	encryptedData := strings.Split(cookie.Value, "--")

	data, _ := base64.StdEncoding.DecodeString(encryptedData[0])
	iv, _ := base64.StdEncoding.DecodeString(encryptedData[1])
	authTag, _ := base64.StdEncoding.DecodeString(encryptedData[2])

	cookie.Data = []byte(append(data, authTag...))
	cookie.Iv = []byte(iv)
}

func (cookie *Rails5Cookie) GenerateKey() []byte {
	return pbkdf2.Key([]byte(cookie.SecretKeyBase), []byte(cookie.Salt), keyIterNum, keySize, sha1.New)
}
