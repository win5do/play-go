package src

import (
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
)

// public key
func pk(filename string) (openpgp.EntityList, error) {
	// Read in public key
	pkFile, _ := os.Open(filename)
	defer pkFile.Close()
	pkEntityList, err := openpgp.ReadArmoredKeyRing(pkFile)
	if err != nil {
		log.Println("err =", err)
		return nil, err
	}

	return pkEntityList, nil
}

// secret key
func sk(filename string, passphrase string) (openpgp.EntityList, error) {
	skFile, err := os.Open(ASK)
	if err != nil {
		return nil, err
	}
	defer skFile.Close()
	skEntityList, err := openpgp.ReadArmoredKeyRing(skFile)
	if err != nil {
		return nil, err
	}

	DecryptPrivateKey(skEntityList, passphrase)

	return skEntityList, nil
}

func DecryptPrivateKey(entityList openpgp.EntityList, passphrase string) {
	if passphrase == "" {
		return
	}

	entity := entityList[0]

	passphraseByte := []byte(passphrase)
	log.Println("Decrypting private key using passphrase")
	entity.PrivateKey.Decrypt(passphraseByte)
	for _, subkey := range entity.Subkeys {
		subkey.PrivateKey.Decrypt(passphraseByte)
	}
	log.Println("Finished decrypting private key using passphrase")
}
