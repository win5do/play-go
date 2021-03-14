package aes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAesCfb(t *testing.T) {
	aesCfb := NewEncryptor("1234567890123456")
	in := "咕咕咕"

	encrypted := aesCfb.Encrypt(in)
	t.Log(encrypted)
	require.NoError(t, aesCfb.Err)

	decrypted := aesCfb.Decrypt(encrypted)
	t.Log(decrypted)
	require.NoError(t, aesCfb.Err)

	require.Equal(t, in, decrypted)
}
