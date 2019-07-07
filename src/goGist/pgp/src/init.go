package src

import (
	_ "golang.org/x/crypto/ripemd160"
)

const (
	BLOCK_TYPE = "PGP MESSAGE"

	PASSPHRASE = "test123456"
	prefix     = "../file/"
	ASK        = prefix + "ask.asc"
	APK        = prefix + "apk.asc"
)
