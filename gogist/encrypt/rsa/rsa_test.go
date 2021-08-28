package rsa

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

var b64cipher = `MxLVJhRU3eleH9o4uGZIT3BK1LZE7W+QqiQTYoIgcFcO2hdpP8shZ8SiTmmVlJlXx+UGqZB+2ZJNvD2kalGgFuq1/vteYI2wT0vdexa1uq4a+oDvXc6Km7WYZq2EAAhhfN4ugIT7qXgmhRPZIb1dzD2uJSJbhJQxoRikfIPitrLGIybQB5a3peSsqF249oRXzb1ly6FF336FAyAOWn/30sQDg75RSGg/ROQ5p9wECm17PzWfeW5N/Q26rNojFkSOW5Y08YoydgQlh2I+pLNnbSlTYBz1C+uKqpK6G4KlE/bdUAu5+dz+RyFsvhMXm4s25PwhaWGLAulQdj0enJU2QQ==`
var origData = "咕咕咕"

func TestRsa(t *testing.T) {
	t.Run("encrypt", func(t *testing.T) {
		ciphertext, err := RsaEncrypt([]byte(origData), publicKey)
		require.NoError(t, err)

		t.Logf("%s", base64.StdEncoding.EncodeToString(ciphertext))

		// 相同的源文，每次加密后的密文都不相同，不能根据加密后的数据是否相等来判断加密算法是否正确
		// 要把加密后得数据再解码，如果和源文一致，说明加密算法正确
		actual, err := RsaDecrypt(ciphertext, privateKey)
		require.NoError(t, err)

		require.Equal(t, origData, string(actual))
	})

	t.Run("decrypt", func(t *testing.T) {
		actual, err := RsaDecryptBase64(b64cipher, privateKey)
		require.NoError(t, err)

		require.Equal(t, origData, string(actual))
	})
}

var privateKey = []byte(`
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDBBUyGQxTy/V+e
sKXJc6tXY1Dz51r6QUm587BvFJyfvX4CTTdq1altzHoTEEQWVZqdLLq8Y4/rv1l9
cUGFIkXSqfbZCY37//rp7+uo3CzNAmL58haB9GN3msoyZDnHYUOKJG5VPUI5AcHC
WPnLyeXNEqLYKHcXB8K445690/0ruY/yv0Ef3TC8BEM/aoGj9A+Tuxwnfl/ebLQd
9gG6qcXehf8H6Kliy4LQL5BoNDq1Fjy6U1WpAZD26CsZ4mpEE69KaHjq2lrPjV9f
IV1kaJkibDDaxKI/AyCsUq3PRrpuPao7ZhD8AcgCOiN9gBFQlx+eJESM7iBP1PHk
rrmuoHh9AgMBAAECggEAS8NEsj6KboY9jhBQQq6ARqDZGaNp9mjCA5c/XZp9j8XV
PBoK8ohDJLHqCKmN0CZxtdfkxCVgTGyjN7XIvfUh6vqDxdUWQh/L4R8nJPSnMSEK
sIjxhLjkggHj87ubYkMvO33pQNP2tmGWKC68fd0VslTqTmYzuHp3WBMgL+qRqSNR
WO1SXznAGxxwp4BFbXwQRdgX3HxuUV03QjJhEH5jpeKCUAG542FX95m6otUZQZ5Q
oRHWzVHlVxlh4IWzPkFkzhsDm6iYCHlg+qx/CLRvF9rzX+3rxl8vwQHOVgS0kB2j
ngPSSYbO9VaeR/y40U3ToTPDIDDs5gkGI266eIqtEQKBgQD48dd+aeIKFMxldX3Y
dlUhaSsv7xcNgs9mYT9iA9Ajo0xWTdZU5Ved2T0Mo2XaALaZoWACjS9QBp+KVppK
hpxbR709QAUU9aRO2fVdjM/5PduCXwX//2E2ltRjV3Qhx3nwSE8KdCsYI9sDNcvn
X0b33eIRTfe1SRmqJQMeXtZE0wKBgQDGfbbOArpkL8BeFoEXtI1lzsNgXzZzTOP0
iPzPsXEZQIm1mW3oUqrqFyJrmahmm9tLr9/wmCHXcI4fIbVlpFdpguO+Al35Spcq
rEmjhHG8Jy6atdmqcih7OBH7aakrz3cj1xI+4llosSKLGL2/1hk2SlDiwZv3Hj0M
yiu6ITQ7bwKBgDQhkNrsGWW7DUKB5CScQx+IAn87NjyiFASAgOFpowCThbzc9/E3
Ra3MfDhVT1Ljq1lorc5a+nlrqaUylYTdw4hp8XsOGQJWnl3UzskUX8j15y3Z8Eu3
kwf3deqHuc4JE8P5oHOz+AJKH+cAMP72uMnXMSwmg6T29eTnc0u82Gn5AoGBALlL
imfu7Amy+iXbZE+44Xb0+jbxK2efk/4oE7nfr/Ee6m949701zxAQ4LUKEMcZIDMb
0DZODumLzWEKmsh0DpYU9n3ev9OE6nGPvH7FmVdITJ1ovfX79AZzZXYKQT8AwfvX
PtUtgMHW+qRXSGvZdlfUMIY+r1aWB+/0m7V6YcRJAoGBAPMixzmAuINbtNZIVUew
lC8nHPtsqEMhXQaRAftY0VD+zE+SmHdloozIHrhsXRwTE8uOj5L1MovCgGARBMA0
tLcwGBd1skFfgAZz+o5OYBmbxRdOEH6fA2qRgiwKzsFWtMhh7DuL6PDhQ3aNkGMe
xPRxGmOxkSIrT0BdOo2DVMU9
-----END PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwQVMhkMU8v1fnrClyXOr
V2NQ8+da+kFJufOwbxScn71+Ak03atWpbcx6ExBEFlWanSy6vGOP679ZfXFBhSJF
0qn22QmN+//66e/rqNwszQJi+fIWgfRjd5rKMmQ5x2FDiiRuVT1COQHBwlj5y8nl
zRKi2Ch3FwfCuOOevdP9K7mP8r9BH90wvARDP2qBo/QPk7scJ35f3my0HfYBuqnF
3oX/B+ipYsuC0C+QaDQ6tRY8ulNVqQGQ9ugrGeJqRBOvSmh46tpaz41fXyFdZGiZ
Imww2sSiPwMgrFKtz0a6bj2qO2YQ/AHIAjojfYARUJcfniREjO4gT9Tx5K65rqB4
fQIDAQAB
-----END PUBLIC KEY-----
`)
