package src

import (
	"bytes"
	"crypto"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func PgpEncryptAndSign(inputString string, signFlag, armorFlag bool) (*bytes.Buffer, error) {
	log.Println("Secret to hide:", inputString)
	log.Println("Public Keyring:", APK)

	pkEntityList, err := pk(APK)
	if err != nil {
		return nil, err
	}

	var signed *openpgp.Entity
	if signFlag {
		skEntityList, err := sk(ASK, PASSPHRASE)
		if err != nil {
			return nil, err
		}
		signed = skEntityList[0]
	} else {
		signed = nil
	}

	// encrypt string
	buf := new(bytes.Buffer)

	var writer io.Writer
	if armorFlag {
		armorWriter, err := armor.Encode(buf, BLOCK_TYPE, nil)
		if err != nil {
			return nil, fmt.Errorf("Error creating OpenPGP armor: %v", err)
		}
		defer armorWriter.Close()
		writer = armorWriter
	} else {
		writer = buf
	}

	config := &packet.Config{
		DefaultHash: crypto.MD4,
	}
	config = nil

	w, err := openpgp.Encrypt(writer, pkEntityList, signed, nil, config)
	if err != nil {
		log.Println("err =", err)
		return nil, err
	}
	defer w.Close()

	_, err = w.Write([]byte(inputString))
	if err != nil {
		return nil, err
	}

	// 需要两个defer close之后才能调用string
	// str := buf.String()

	return buf, nil
}
