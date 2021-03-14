package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

// AesCfbMode AES CFB 加密
type AesCfbMode struct {
	Key []byte
	Err error
}

func NewEncryptor(key string) *AesCfbMode {
	return &AesCfbMode{
		Key: []byte(key),
		Err: nil,
	}
}

func (a *AesCfbMode) Encrypt(plaintext string) string {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		a.Err = err
		log.Printf("err: %s", err)
		return ""
	}
	pbt := []byte(plaintext)
	ciphertext := make([]byte, aes.BlockSize+len(pbt))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		a.Err = err
		log.Printf("err: %s", err)
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], pbt)

	return hex.EncodeToString(ciphertext)
}

func (a *AesCfbMode) Decrypt(ciphertext string) string {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		a.Err = err
		log.Printf("err: %s", err)
		return ""
	}

	cbt, err := hex.DecodeString(ciphertext)
	if err != nil {
		a.Err = err
		log.Printf("err: %s", err)
		return ""
	}

	if len(cbt) < aes.BlockSize {
		err = fmt.Errorf("the length of the ciphertext is too short")
		a.Err = err
		log.Printf("err: %s", err)
		return ""
	}
	iv := cbt[:aes.BlockSize]
	cbt = cbt[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cbt, cbt)
	return string(cbt)
}
