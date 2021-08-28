package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
)

func ParsePKCS1Or8PrivateKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}
	pk1, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return pk1, nil
	}

	pk8Iface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err == nil {
		if pk8, ok := pk8Iface.(*rsa.PrivateKey); ok {
			return pk8, nil
		}
	}
	return nil, errors.New("parse private key error")
}

func RsaEncrypt(plaintext, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Printf("x509.ParsePKIXPublicKey error: %s\n", err)
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, plaintext)
}

func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	priv, err := ParsePKCS1Or8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RsaDecryptBase64(b64Cipher string, privateKey []byte) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(b64Cipher)
	if err != nil {
		log.Printf("err: %s", err)
		return nil, err
	}

	return RsaDecrypt(cipherText, privateKey)
}
