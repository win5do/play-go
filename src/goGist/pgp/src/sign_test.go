package src

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

const InputString = "hello world"

func TestPgpSign(t *testing.T) {
	rq := require.New(t)
	buf, err := PgpSign(InputString, true)
	rq.Nil(err)
	t.Log(buf.String())

	err = ioutil.WriteFile("../file/golangSign.pgp", buf.Bytes(), 0644)
	rq.Nil(err)

	buf, err = PgpSign(InputString, false)
	rq.Nil(err)
	t.Log(buf.String())
}

func TestPgpDetachSign(t *testing.T) {
	rq := require.New(t)
	buf, err := PgpDetachSign(InputString, true)
	rq.Nil(err)
	t.Log(buf.String())

	err = ioutil.WriteFile("../file/detachSign.pgp", buf.Bytes(), 0644)
	rq.Nil(err)

	buf, err = PgpDetachSign(InputString, false)
	rq.Nil(err)
	t.Log(buf.String())
}
