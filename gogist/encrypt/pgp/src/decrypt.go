package src

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func PgpDecrypt(encString string, armorFlag bool) (string, error) {
	log.Println("Secret Keyring:", ASK)

	skEntityList, err := sk(ASK, PASSPHRASE)
	if err != nil {
		return "", err
	}

	pkEntityList, err := pk(APK)
	if err != nil {
		log.Printf("err = %s\n", err)
		return "", err
	}

	keyring := append(skEntityList, pkEntityList...)

	var reader io.Reader
	reader = bytes.NewBufferString(encString)
	if armorFlag {
		block, err := armor.Decode(reader)
		if err != nil {
			return "", err
		}
		reader = block.Body
	}

	// Decrypt it with the contents of the private key
	md, err := openpgp.ReadMessage(reader, keyring, nil, nil)
	if err != nil {
		return "", err
	}

	bts, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return "", err
	}
	decStr := string(bts)

	log.Printf("md = %+v\n", md)
	log.Printf("Signature = %+v\n", md.Signature)
	if md.SignatureError != nil {
		log.Printf("err = %s\n", md.SignatureError)
	}

	return decStr, nil
}
