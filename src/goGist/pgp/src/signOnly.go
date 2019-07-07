package src

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func PgpSign(inputString string, armorFlag bool) (*bytes.Buffer, error) {
	skEntityList, err := sk(ASK, PASSPHRASE)
	if err != nil {
		return nil, err
	}

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

	w, err := openpgp.Sign(writer, skEntityList[0], nil /* no hints */, nil)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	_, err = w.Write([]byte(inputString))
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func PgpDetachSign(inputString string, armorFlag bool) (*bytes.Buffer, error) {
	skEntityList, err := sk(ASK, PASSPHRASE)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	if armorFlag {
		openpgp.ArmoredDetachSign(buf, skEntityList[0], bytes.NewBufferString(inputString), nil)
	} else {
		openpgp.DetachSign(buf, skEntityList[0], bytes.NewBufferString(inputString), nil)
	}

	return buf, nil
}
