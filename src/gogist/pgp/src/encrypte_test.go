package src

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	rq := require.New(t)

	encBuf, err := PgpEncryptAndSign("golang", true, true)
	rq.Nil(err)
	t.Log(encBuf.String())

	err = ioutil.WriteFile("../file/golang.pgp", encBuf.Bytes(), 0644)
	rq.Nil(err)
}

func TestDecrypt(t *testing.T) {
	rq := require.New(t)

	enc, err := ioutil.ReadFile("../file/golang.pgp")
	rq.Nil(err)

	dec, err := PgpDecrypt(string(enc), true)
	rq.Nil(err)
	t.Log(dec)
}
