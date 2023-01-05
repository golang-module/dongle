package rsa

import (
	"crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rsaInput       = "hello world"
	pkcs1PublicKey = `-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEw
cujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFn
dk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE=
-----END RSA PUBLIC KEY-----`
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
	pkcs8PublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`
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

func TestRsa_PKCS1_Encrypt(t *testing.T) {
	keyPair := NewKeyPair()
	keyPair.SetPublicKey([]byte(pkcs1PublicKey))
	keyPair.SetPrivateKey([]byte(pkcs1PrivateKey))

	dst1, err1 := keyPair.EncryptByPublicKey([]byte(rsaInput))
	assert.Nil(t, err1)
	dst2, err2 := keyPair.DecryptByPrivateKey(dst1)
	assert.Nil(t, err2)
	assert.Equal(t, []byte(rsaInput), dst2)

	dst3, err3 := keyPair.EncryptByPrivateKey([]byte(rsaInput))
	assert.Nil(t, err3)
	dst4, err4 := keyPair.DecryptByPublicKey(dst3)
	assert.Nil(t, err4)
	assert.Equal(t, []byte(rsaInput), dst4)
}

func TestRsa_PKCS1_Sign(t *testing.T) {
	keyPair := NewKeyPair()
	keyPair.SetPublicKey([]byte(pkcs1PublicKey))
	keyPair.SetPrivateKey([]byte(pkcs1PrivateKey))
	keyPair.SetHash(crypto.SHA224)

	dst1, err1 := keyPair.SignByPrivateKey([]byte(rsaInput))
	assert.Nil(t, err1)
	err2 := keyPair.VerifyByPublicKey([]byte(rsaInput), dst1)
	assert.Nil(t, err2)
	assert.Equal(t, nil, err2)
}

func TestRsa_PKCS8_Encrypt(t *testing.T) {
	keyPair := NewKeyPair()
	keyPair.SetPublicKey([]byte(pkcs8PublicKey))
	keyPair.SetPrivateKey([]byte(pkcs8PrivateKey))

	dst1, err1 := keyPair.EncryptByPublicKey([]byte(rsaInput))
	assert.Nil(t, err1)
	dst2, err2 := keyPair.DecryptByPrivateKey(dst1)
	assert.Nil(t, err2)
	assert.Equal(t, []byte(rsaInput), dst2)

	dst3, err3 := keyPair.EncryptByPrivateKey([]byte(rsaInput))
	assert.Nil(t, err3)
	dst4, err4 := keyPair.DecryptByPublicKey(dst3)
	assert.Nil(t, err4)
	assert.Equal(t, []byte(rsaInput), dst4)
}

func TestRsa_PKCS8_Sign(t *testing.T) {
	keyPair := NewKeyPair()
	keyPair.SetPublicKey([]byte(pkcs8PublicKey))
	keyPair.SetPrivateKey([]byte(pkcs8PrivateKey))
	keyPair.SetHash(crypto.SHA384)

	dst1, err1 := keyPair.SignByPrivateKey([]byte(rsaInput))
	assert.Nil(t, err1)
	err2 := keyPair.VerifyByPublicKey([]byte(rsaInput), dst1)
	assert.Nil(t, err2)
	assert.Equal(t, nil, err2)
}

func TestRsa_Empty_Src(t *testing.T) {
	keyPair := NewKeyPair()
	keyPair.SetPublicKey([]byte(pkcs1PublicKey))
	keyPair.SetPrivateKey([]byte(pkcs1PrivateKey))

	empty := ""
	dst1, err1 := keyPair.EncryptByPublicKey([]byte(empty))
	assert.Nil(t, err1)
	dst2, err2 := keyPair.DecryptByPrivateKey(dst1)
	assert.Nil(t, err2)
	assert.Equal(t, []byte(empty), dst2)

	dst3, err3 := keyPair.EncryptByPrivateKey([]byte(empty))
	assert.Nil(t, err3)
	dst4, err4 := keyPair.DecryptByPublicKey(dst3)
	assert.Nil(t, err4)
	assert.Equal(t, []byte(empty), dst4)
}

func TestRsa_IsKey(t *testing.T) {
	keyPair := NewKeyPair()

	keyPair.SetPublicKey([]byte(pkcs1PublicKey))
	keyPair.SetPrivateKey([]byte(pkcs1PrivateKey))
	assert.Equal(t, true, keyPair.IsPublicKey())
	assert.Equal(t, true, keyPair.IsPrivateKey())

	keyPair.SetPublicKey([]byte(pkcs8PublicKey))
	keyPair.SetPrivateKey([]byte(pkcs8PrivateKey))
	assert.Equal(t, true, keyPair.IsPublicKey())
	assert.Equal(t, true, keyPair.IsPrivateKey())

	keyPair.SetPublicKey([]byte(pkcs1PrivateKey))
	keyPair.SetPrivateKey([]byte(pkcs1PublicKey))
	assert.Equal(t, false, keyPair.IsPublicKey())
	assert.Equal(t, false, keyPair.IsPrivateKey())

	keyPair.SetPublicKey([]byte(pkcs8PrivateKey))
	keyPair.SetPrivateKey([]byte(pkcs8PublicKey))
	assert.Equal(t, false, keyPair.IsPublicKey())
	assert.Equal(t, false, keyPair.IsPrivateKey())

	keyPair.SetPublicKey([]byte(""))
	keyPair.SetPrivateKey([]byte(""))
	assert.Equal(t, false, keyPair.IsPublicKey())
	assert.Equal(t, false, keyPair.IsPrivateKey())

	keyPair.SetPublicKey([]byte("xxx"))
	keyPair.SetPrivateKey([]byte("xxx"))
	assert.Equal(t, false, keyPair.IsPublicKey())
	assert.Equal(t, false, keyPair.IsPrivateKey())
}

func TestRsa_PublicKey_Error(t *testing.T) {
	invalidPublicKey := `-----BEGIN PUBLIC KEY-----
xxxx
-----END PUBLIC KEY-----`

	keyPair := NewKeyPair()
	keyPair.SetPrivateKey([]byte(pkcs1PrivateKey))
	keyPair.SetHash(crypto.SHA1)

	sign, _ := keyPair.SignByPrivateKey([]byte(rsaInput))

	keyPair.SetPublicKey([]byte(pkcs1PrivateKey))
	_, err1 := keyPair.EncryptByPublicKey([]byte(rsaInput))
	assert.Equal(t, invalidPublicKeyError(), err1)
	_, err2 := keyPair.DecryptByPublicKey(sign)
	assert.Equal(t, invalidPublicKeyError(), err2)
	err3 := keyPair.VerifyByPublicKey([]byte(rsaInput), sign)
	assert.Equal(t, invalidPublicKeyError(), err3)

	keyPair.SetPublicKey([]byte(invalidPublicKey))
	_, err4 := keyPair.EncryptByPublicKey([]byte(rsaInput))
	assert.Equal(t, invalidPublicKeyError(), err4)
	_, err5 := keyPair.DecryptByPublicKey(sign)
	assert.Equal(t, invalidPublicKeyError(), err5)
	err6 := keyPair.VerifyByPublicKey([]byte(rsaInput), sign)
	assert.Equal(t, invalidPublicKeyError(), err6)

	keyPair.SetPrivateKey([]byte(pkcs1PrivateKey))
	_, err7 := keyPair.EncryptByPublicKey([]byte(rsaInput))
	assert.Equal(t, invalidPublicKeyError(), err7)
	_, err8 := keyPair.DecryptByPublicKey([]byte(rsaInput))
	assert.Equal(t, invalidPublicKeyError(), err8)
	err9 := keyPair.VerifyByPublicKey([]byte(rsaInput), sign)
	assert.Equal(t, invalidPublicKeyError(), err9)
}

func TestRsa_PrivateKey_Error(t *testing.T) {
	invalidPrivateKey := `-----BEGIN PRIVATE KEY-----
xxxx
-----END PRIVATE KEY-----`

	keyPair := NewKeyPair()
	keyPair.SetHash(crypto.SHA512)

	keyPair.SetPrivateKey([]byte(invalidPrivateKey))
	_, err1 := keyPair.EncryptByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err1)
	_, err2 := keyPair.DecryptByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err2)
	_, err3 := keyPair.SignByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err3)

	keyPair.SetPrivateKey([]byte(pkcs1PublicKey))
	_, err4 := keyPair.EncryptByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err4)
	_, err5 := keyPair.DecryptByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err5)
	_, err6 := keyPair.SignByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err6)

	keyPair.SetPrivateKey([]byte(pkcs1PublicKey))
	_, err7 := keyPair.EncryptByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err7)
	_, err8 := keyPair.DecryptByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err8)
	_, err9 := keyPair.SignByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidPrivateKeyError(), err9)
}

func TestRsa_Parse_Error(t *testing.T) {
	invalidPublicKey := "xxxx"
	invalidPrivateKey := "xxxx"

	keyPair := NewKeyPair()

	keyPair.SetPublicKey([]byte(invalidPublicKey))
	keyPair.SetPrivateKey([]byte(invalidPrivateKey))

	_, err1 := keyPair.ParsePublicKey()
	assert.Equal(t, invalidPublicKeyError(), err1)
	_, err2 := keyPair.ParsePrivateKey()
	assert.Equal(t, invalidPrivateKeyError(), err2)
}

func TestRsa_Hash_Error(t *testing.T) {
	keyPair := NewKeyPair()
	keyPair.SetPublicKey([]byte(pkcs1PublicKey))
	keyPair.SetPrivateKey([]byte(pkcs1PrivateKey))
	keyPair.SetHash(crypto.MD4)

	_, err1 := keyPair.SignByPrivateKey([]byte(rsaInput))
	assert.Equal(t, invalidHashError(), err1)
	err2 := keyPair.VerifyByPublicKey([]byte(rsaInput), []byte(rsaInput))
	assert.Equal(t, invalidHashError(), err2)
}
