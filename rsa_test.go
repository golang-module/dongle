package dongle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	rsaInput       = "hello world"
	pkcs1PublicKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEw
cujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFn
dk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE=
-----END RSA PUBLIC KEY-----`)
	pkcs1PrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)
	pkcs8PublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`)
	pkcs8PrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
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
-----END PRIVATE KEY-----`)
)

func TestRsa_EncryptWithPublicKy(t *testing.T) {
	e1 := Encrypt.FromString(rsaInput).ByRsa(pkcs1PublicKey)
	assert.Nil(t, e1.Error)
	d11 := Decrypt.FromRawString(e1.ToRawString()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d11.Error)
	assert.Equal(t, rsaInput, d11.ToString())
	d12 := Decrypt.FromHexString(e1.ToHexString()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d12.Error)
	assert.Equal(t, rsaInput, d12.ToString())
	d13 := Decrypt.FromBase64String(e1.ToBase64String()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d13.Error)
	assert.Equal(t, rsaInput, d13.ToString())

	e8 := Encrypt.FromString(rsaInput).ByRsa(pkcs8PublicKey)
	assert.Nil(t, e8.Error)
	d81 := Decrypt.FromRawString(e8.ToRawString()).ByRsa(pkcs8PrivateKey)
	assert.Nil(t, d81.Error)
	assert.Equal(t, rsaInput, d11.ToString())
	d82 := Decrypt.FromHexString(e8.ToHexString()).ByRsa(pkcs8PrivateKey)
	assert.Nil(t, d82.Error)
	assert.Equal(t, rsaInput, d12.ToString())
	d83 := Decrypt.FromBase64String(e8.ToBase64String()).ByRsa(pkcs8PrivateKey)
	assert.Nil(t, d83.Error)
	assert.Equal(t, rsaInput, d83.ToString())
}

func TestRsa_EncryptWithPrivateKy(t *testing.T) {
	e1 := Encrypt.FromString(rsaInput).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, e1.Error)
	d11 := Decrypt.FromRawString(e1.ToRawString()).ByRsa(pkcs1PublicKey)
	assert.Nil(t, d11.Error)
	assert.Equal(t, rsaInput, d11.ToString())
	d12 := Decrypt.FromHexString(e1.ToHexString()).ByRsa(pkcs1PublicKey)
	assert.Nil(t, d12.Error)
	assert.Equal(t, rsaInput, d12.ToString())
	d13 := Decrypt.FromBase64String(e1.ToBase64String()).ByRsa(pkcs1PublicKey)
	assert.Nil(t, d13.Error)
	assert.Equal(t, rsaInput, d13.ToString())

	e8 := Encrypt.FromString(rsaInput).ByRsa(pkcs8PrivateKey)
	assert.Nil(t, e8.Error)
	d81 := Decrypt.FromRawString(e8.ToRawString()).ByRsa(pkcs8PublicKey)
	assert.Nil(t, d81.Error)
	assert.Equal(t, rsaInput, d81.ToString())
	d82 := Decrypt.FromHexString(e8.ToHexString()).ByRsa(pkcs8PublicKey)
	assert.Nil(t, d82.Error)
	assert.Equal(t, rsaInput, d82.ToString())
	d83 := Decrypt.FromBase64String(e8.ToBase64String()).ByRsa(pkcs8PublicKey)
	assert.Nil(t, d83.Error)
	assert.Equal(t, rsaInput, d83.ToString())
}

func TestRsa_SignWithPrivateKy(t *testing.T) {
	s1 := Sign.FromString(rsaInput).ByRsa(pkcs1PrivateKey, SHA224)
	assert.Nil(t, s1.Error)
	v11 := Verify.FromRawString(s1.ToRawString(), rsaInput).ByRsa(pkcs1PublicKey, SHA224)
	assert.Nil(t, v11.Error)
	assert.Equal(t, true, v11.ToBool())
	v12 := Verify.FromHexString(s1.ToHexString(), rsaInput).ByRsa(pkcs1PublicKey, SHA224)
	assert.Nil(t, v12.Error)
	assert.Equal(t, true, v12.ToBool())
	v13 := Verify.FromBase64String(s1.ToBase64String(), rsaInput).ByRsa(pkcs1PublicKey, SHA224)
	assert.Nil(t, v13.Error)
	assert.Equal(t, true, v13.ToBool())

	s8 := Sign.FromString(rsaInput).ByRsa(pkcs8PrivateKey, SHA256)
	assert.Nil(t, s8.Error)
	v81 := Verify.FromRawString(s8.ToRawString(), rsaInput).ByRsa(pkcs8PublicKey, SHA256)
	assert.Nil(t, v81.Error)
	assert.Equal(t, true, v81.ToBool())
	v82 := Verify.FromHexString(s8.ToHexString(), rsaInput).ByRsa(pkcs8PublicKey, SHA256)
	assert.Nil(t, v82.Error)
	assert.Equal(t, true, v12.ToBool())
	v83 := Verify.FromBase64String(s8.ToBase64String(), rsaInput).ByRsa(pkcs8PublicKey, SHA256)
	assert.Nil(t, v83.Error)
	assert.Equal(t, true, v83.ToBool())
}

func TestRsa_Empty_Src(t *testing.T) {
	empty := ""
	e := Encrypt.FromString(empty).ByRsa(pkcs1PublicKey)
	assert.Nil(t, e.Error)
	d1 := Decrypt.FromRawString(e.ToRawString()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d1.Error)
	assert.Equal(t, empty, d1.ToString())
	d2 := Decrypt.FromHexString(e.ToHexString()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d2.Error)
	assert.Equal(t, empty, d2.ToString())
	d3 := Decrypt.FromBase64String(e.ToBase64String()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d3.Error)
	assert.Equal(t, empty, d3.ToString())

	s := Sign.FromString(empty).ByRsa(pkcs1PrivateKey, MD5)
	assert.Nil(t, s.Error)
	v1 := Verify.FromRawString(s.ToRawString(), empty).ByRsa(pkcs1PublicKey, MD5)
	assert.Nil(t, v1.Error)
	assert.Equal(t, false, v1.ToBool())
	v2 := Verify.FromHexString(s.ToHexString(), empty).ByRsa(pkcs1PublicKey, MD5)
	assert.Nil(t, v2.Error)
	assert.Equal(t, false, v2.ToBool())
	v3 := Verify.FromBase64String(s.ToBase64String(), empty).ByRsa(pkcs1PublicKey, MD5)
	assert.Nil(t, v1.Error)
	assert.Equal(t, false, v3.ToBool())
}

func TestRsa_PublicKey_Error(t *testing.T) {
	invalidRsaKey := []byte("xxxx")
	invalidPublicKey := `-----BEGIN PUBLIC KEY-----
xxxx
-----END PUBLIC KEY-----`

	e1 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa(invalidRsaKey)
	assert.Equal(t, rsaError.PublicKeyError(), e1.Error)
	e2 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(invalidPublicKey))
	assert.Equal(t, rsaError.PublicKeyError(), e2.Error)

	v1 := Verify.FromRawBytes([]byte("xxxx"), []byte(rsaInput)).ByRsa(invalidRsaKey, SHA1)
	assert.Equal(t, rsaError.PublicKeyError(), v1.Error)
	v2 := Verify.FromRawBytes([]byte("xxxx"), []byte(rsaInput)).ByRsa([]byte(invalidPublicKey), SHA224)
	assert.Equal(t, rsaError.PublicKeyError(), v2.Error)

	v3 := Verify.FromHexString("xxxx", rsaInput).ByRsa(pkcs8PublicKey, SHA256)
	assert.Equal(t, decodeError.ModeError("hex"), v3.Error)
	v4 := Verify.FromHexBytes([]byte("xxxx"), []byte(rsaInput)).ByRsa(pkcs8PublicKey, SHA384)
	assert.Equal(t, decodeError.ModeError("hex"), v4.Error)
	v5 := Verify.FromBase64String("xxxxxx", rsaInput).ByRsa(pkcs8PublicKey, SHA512)
	assert.Equal(t, decodeError.ModeError("base64"), v5.Error)
	v6 := Verify.FromBase64Bytes([]byte("xxxxxx"), []byte(rsaInput)).ByRsa(pkcs8PublicKey, RIPEMD160)
	assert.Equal(t, decodeError.ModeError("base64"), v6.Error)
}

func TestRsa_PrivateKey_Error(t *testing.T) {
	invalidRsaKey := []byte("xxxx")
	invalidPrivateKey := []byte(`-----BEGIN PRIVATE KEY-----
xxxx
-----END PRIVATE KEY-----`)

	e := Encrypt.FromBytes([]byte(rsaInput)).ByRsa(pkcs1PublicKey)

	d1 := Decrypt.FromHexBytes(e.ToHexBytes()).ByRsa(invalidRsaKey)
	assert.Equal(t, rsaError.PrivateKeyError(), d1.Error)
	d2 := Decrypt.FromBase64Bytes(e.ToBase64Bytes()).ByRsa(invalidPrivateKey)
	assert.Equal(t, rsaError.PrivateKeyError(), d2.Error)

	s1 := Sign.FromBytes([]byte(rsaInput)).ByRsa(invalidRsaKey, SHA1)
	assert.Equal(t, rsaError.PrivateKeyError(), s1.Error)
	s2 := Sign.FromBytes([]byte(rsaInput)).ByRsa(invalidPrivateKey, SHA1)
	assert.Equal(t, rsaError.PrivateKeyError(), s2.Error)
}
