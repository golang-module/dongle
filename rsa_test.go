package dongle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rsaInput       = "hello world"
	pkcs1PublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCtdjE3fOKpAmc6eIi1I/GElIJW
hTLBZb/SpSC+Pl7bONDyt/sG7FzxMHLoxnlJnYDHmYiu0iy0s/EHhNL+1bPnBiuA
bOKi7TBgojusBLsrddCrGn08gZLRZ3ZP/oAS6+wf2/vfh2jCCKCVHbBZLHQIN3MU
0qcFQKXvfJaBTUX6gwIDAQAB
-----END PUBLIC KEY-----`
	pkcs8PublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`
	pkcs1PrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCtdjE3fOKpAmc6eIi1I/GElIJWhTLBZb/SpSC+Pl7bONDyt/sG
7FzxMHLoxnlJnYDHmYiu0iy0s/EHhNL+1bPnBiuAbOKi7TBgojusBLsrddCrGn08
gZLRZ3ZP/oAS6+wf2/vfh2jCCKCVHbBZLHQIN3MU0qcFQKXvfJaBTUX6gwIDAQAB
AoGAFwAfEo56t5JcAcLNzccQVVYj2jkbO820G8hNiSxYA5WLD0QaAxcAU/Lqqbb3
ii1aUB0ppJS13NgnU6nnGGdZzUYBG1Hai6EkVyCGrI4amQ93AaVdKncL8gJ4RZAm
YzPPUwSMEESsu24pS1NF1G1Y8C+28b/Wr0oqOsCvL6PhsMECQQDwsPJJoWRx7ZJw
E1K5KLT0cXKyrIpyXY3I6tyA5imCzOzccf3d1vDgB0L9sdSO7bG3ceSwpAeiWEbg
5jGZemPzAkEAuH6U4pEI4AMbWnatpK55Rc235NDgmT3VyIuRaKC02YXAZ+jznFep
XMd4DTli4R9r3j2YVhUpyDVbdQpFH98DMQJAQpOvcU6DSkA80WOG7lCkPTlkUKgJ
Y7kdDwZoF/+SW+vzWMbvQf3CgzV/Ak2+TgrRrbyDVZkJw45HjM4fyiRgoQJBALH/
/qlxgPyQQs3O/s2KQBsm1auAE5IF5MLuVUZ69sF/mBko2hEXSqHnGV645TuKU0pC
Zz12ga9WO3z6gaK0SaECQQDah1pKt9ViBBy4USXK3OWXEloHuTwmyr9AbLqqI5tQ
2eNuH0NkuJYQmnXmHLbKOELoYocldEBXmkzPXSN+X9kV
-----END RSA PRIVATE KEY-----`
	pkcs8PrivateKey = `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAKrNk1r1Wtx7DJTr
AOhXtj2QAepfVUrQHdFvoY2ZB7jMsR9x7txVNoutzhUZMqXfm0AMbVxEeq1obhL9
a22mIZkGHEnLgyk5dvp4g+JUuyfaUv6smjld1tKveDKPEQ5BD3uKG3DiUN3nAyjh
sg67DUu0x7McLWi62UzrH78EHQFJAgMBAAECgYAeo3nHWzPNURVUsUMcan96U5bE
YA2AugxfQVMNf2HvOGidZ2adh3udWrQY/MglERNcTd5gKriG2rDEH0liBecIrNKs
BL4lV+qHEGRUcnDDdtUBdGInEU8lve5keDgmX+/huXSRJ+3tYA5u9j+32RquVczv
Idtb5XnBLUl61k0osQJBAON5+eJjtw6xpn+pveU92BSHvaJYVyrLHwUjR07aNKb7
GlGVM3MGf1FCa8WQUo9uUzYxGLtg5Qf3sqwOrwPd5UsCQQDAOF/zWqGuY3HfV/1w
giXiWp8rc+S8tanMj5M37QQbYW5YLjUmJImoklVahv3qlgLZdEN5ZSueM5jfoSFt
Nts7AkBKoRDvSiGbi4MBbTHkzLZgfewkH/FxE7S4nctePk553fXTgCyh9ya8BRuQ
dHnxnpNkOxVPHEnnpEcVFbgrf5gjAkB7KmRI4VTiEfRgINhTJAG0VU7SH/N7+4cu
fPzfA+7ywG5c8Fa79wOB0SoB1KeUjcSLo5Ssj2fwea1F9dAeU90LAkBJQFofveaD
a3YlN4EQZOcCvJKmg7xwWuGxFVTZDVVEws7UCQbEOEEXZrNd9x0IF5kpPLR+rxua
RPgUNaDGIh5o
-----END PRIVATE KEY-----`
)

func TestRsa_PKCS1_ToString(t *testing.T) {
	e1 := Encrypt.FromString("").ByRsa(pkcs1PublicKey)
	d1 := Decrypt.FromHexString(e1.ToHexString()).ByRsa(pkcs1PrivateKey, PKCS1)
	assert.Nil(t, d1.Error)
	assert.Equal(t, "", d1.ToString())

	e2 := Encrypt.FromString(rsaInput).ByRsa(pkcs1PublicKey)
	d2 := Decrypt.FromHexString(e2.ToHexString()).ByRsa(pkcs1PrivateKey, PKCS1)
	assert.Nil(t, d2.Error)
	assert.Equal(t, rsaInput, d2.ToString())

}

func TestRsa_PKCS1_ToBytes(t *testing.T) {
	e1 := Encrypt.FromBytes([]byte("")).ByRsa([]byte(pkcs1PublicKey))
	d1 := Decrypt.FromHexBytes(e1.ToHexBytes()).ByRsa([]byte(pkcs1PrivateKey), PKCS1)
	assert.Nil(t, d1.Error)
	assert.Equal(t, []byte(""), d1.ToBytes())

	e2 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs1PublicKey))
	d2 := Decrypt.FromHexBytes(e2.ToHexBytes()).ByRsa([]byte(pkcs1PrivateKey), PKCS1)
	assert.Nil(t, d2.Error)
	assert.Equal(t, []byte(rsaInput), d2.ToBytes())
}

func TestRsa_PKCS8_ToString(t *testing.T) {
	e1 := Encrypt.FromString("").ByRsa(pkcs8PublicKey)
	d1 := Decrypt.FromHexString(e1.ToHexString()).ByRsa(pkcs8PrivateKey, PKCS1)
	assert.Nil(t, d1.Error)
	assert.Equal(t, "", d1.ToString())

	e2 := Encrypt.FromString(rsaInput).ByRsa(pkcs8PublicKey)
	d2 := Decrypt.FromHexString(e2.ToHexString()).ByRsa(pkcs8PrivateKey, PKCS8)
	assert.Nil(t, d2.Error)
	assert.Equal(t, rsaInput, d2.ToString())
}

func TestRsa_PKCS8_ToBytes(t *testing.T) {
	e1 := Encrypt.FromBytes([]byte("")).ByRsa([]byte(pkcs8PublicKey))
	d1 := Decrypt.FromHexBytes(e1.ToHexBytes()).ByRsa([]byte(pkcs8PrivateKey), PKCS1)
	assert.Nil(t, d1.Error)
	assert.Equal(t, []byte(""), d1.ToBytes())

	e2 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs8PublicKey))
	d2 := Decrypt.FromHexBytes(e2.ToHexBytes()).ByRsa([]byte(pkcs8PrivateKey), PKCS8)
	assert.Nil(t, d2.Error)
	assert.Equal(t, []byte(rsaInput), d2.ToBytes())
}

func TestRsa_FromString_Error(t *testing.T) {
	e1 := Encrypt.FromString(rsaInput).ByRsa("xxxx")
	assert.Equal(t, invalidPublicKeyError(), e1.Error)

	e2 := Encrypt.FromString(rsaInput).ByRsa(`-----BEGIN PUBLIC KEY-----
xxxx
-----END PUBLIC KEY-----`)
	assert.Equal(t, invalidPublicKeyError(), e2.Error)

	e3 := Encrypt.FromString(rsaInput).ByRsa(pkcs1PublicKey)
	d3 := Decrypt.FromHexString(e3.ToHexString()).ByRsa(pkcs1PrivateKey, PKCS8)
	assert.Equal(t, invalidPrivateKeyError(), d3.Error)

	e4 := Encrypt.FromString(rsaInput).ByRsa(pkcs8PublicKey)
	d4 := Decrypt.FromHexString(e4.ToHexString()).ByRsa(pkcs8PrivateKey, PKCS1)
	assert.Equal(t, invalidPrivateKeyError(), d4.Error)

	e5 := Encrypt.FromString(rsaInput).ByRsa(pkcs1PublicKey)
	d5 := Decrypt.FromHexString(e5.ToHexString()).ByRsa("xxxx", PKCS1)
	assert.Equal(t, invalidPrivateKeyError(), d5.Error)

	e6 := Encrypt.FromString(rsaInput).ByRsa(pkcs1PublicKey)
	d6 := Decrypt.FromHexString(e6.ToHexString()).ByRsa(`-----BEGIN PRIVATE KEY-----
xxxx
-----END PRIVATE KEY-----`, PKCS8)
	assert.Equal(t, invalidPrivateKeyError(), d6.Error)
}

func TestRsa_FromBytes_Error(t *testing.T) {
	e1 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte("xxxx"))
	assert.Equal(t, invalidPublicKeyError(), e1.Error)

	e2 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(`-----BEGIN PUBLIC KEY-----
xxxx
-----END PUBLIC KEY-----`))
	assert.Equal(t, invalidPublicKeyError(), e2.Error)

	e3 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs1PublicKey))
	d3 := Decrypt.FromHexString(e3.ToHexString()).ByRsa([]byte(pkcs1PrivateKey), PKCS8)
	assert.Equal(t, invalidPrivateKeyError(), d3.Error)

	e4 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs8PublicKey))
	d4 := Decrypt.FromHexBytes(e4.ToHexBytes()).ByRsa([]byte(pkcs8PrivateKey), PKCS1)
	assert.Equal(t, invalidPrivateKeyError(), d4.Error)

	e5 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs1PublicKey))
	d5 := Decrypt.FromHexBytes(e5.ToHexBytes()).ByRsa("xxxx", PKCS1)
	assert.Equal(t, invalidPrivateKeyError(), d5.Error)

	e6 := Encrypt.FromString(rsaInput).ByRsa([]byte(pkcs1PublicKey))
	d6 := Decrypt.FromHexBytes(e6.ToHexBytes()).ByRsa([]byte(`-----BEGIN PRIVATE KEY-----
xxxx
-----END PRIVATE KEY-----`), PKCS8)
	assert.Equal(t, invalidPrivateKeyError(), d6.Error)

	e7 := Encrypt.FromString(rsaInput).ByRsa(pkcs1PublicKey)
	d7 := Decrypt.FromHexString(e7.ToHexString()).ByRsa(`-----BEGIN PRIVATE KEY-----
xxxx
-----END PRIVATE KEY-----`, "xxxx")
	assert.Equal(t, invalidPrivateKeyError(), d7.Error)
}
