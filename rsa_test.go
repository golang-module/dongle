package dongle

import (
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

func TestRsa_PKCS1_String(t *testing.T) {
	e := Encrypt.FromString(rsaInput).ByRsa(pkcs1PublicKey)
	assert.Nil(t, e.Error)
	d1 := Decrypt.FromRawString(e.ToRawString()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d1.Error)
	assert.Equal(t, rsaInput, d1.ToString())
	d2 := Decrypt.FromHexString(e.ToHexString()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d2.Error)
	assert.Equal(t, rsaInput, d2.ToString())
	d3 := Decrypt.FromBase64String(e.ToBase64String()).ByRsa(pkcs1PrivateKey)
	assert.Nil(t, d3.Error)
	assert.Equal(t, rsaInput, d3.ToString())

	s := Sign.FromString(rsaInput).ByRsa(pkcs1PrivateKey, SHA384)
	assert.Nil(t, s.Error)
	v1 := Verify.FromRawString(s.ToRawString(), rsaInput).ByRsa(pkcs1PublicKey, SHA384)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v1.ToBool())
	v2 := Verify.FromHexString(s.ToHexString(), rsaInput).ByRsa(pkcs1PublicKey, SHA384)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v2.ToBool())
	v3 := Verify.FromBase64String(s.ToBase64String(), rsaInput).ByRsa(pkcs1PublicKey, SHA384)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v3.ToBool())
}

func TestRsa_PKCS1_Bytes(t *testing.T) {
	e := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs1PublicKey))
	assert.Nil(t, e.Error)
	d1 := Decrypt.FromRawBytes(e.ToRawBytes()).ByRsa([]byte(pkcs1PrivateKey))
	assert.Nil(t, d1.Error)
	assert.Equal(t, []byte(rsaInput), d1.ToBytes())
	d2 := Decrypt.FromHexBytes(e.ToHexBytes()).ByRsa([]byte(pkcs1PrivateKey))
	assert.Nil(t, d2.Error)
	assert.Equal(t, []byte(rsaInput), d2.ToBytes())
	d3 := Decrypt.FromBase64Bytes(e.ToBase64Bytes()).ByRsa([]byte(pkcs1PrivateKey))
	assert.Nil(t, d3.Error)
	assert.Equal(t, []byte(rsaInput), d3.ToBytes())

	s := Sign.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs1PrivateKey), SHA256)
	assert.Nil(t, s.Error)
	v1 := Verify.FromRawBytes(s.ToRawBytes(), []byte(rsaInput)).ByRsa([]byte(pkcs1PublicKey), SHA256)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v1.ToBool())
	v2 := Verify.FromHexBytes(s.ToHexBytes(), []byte(rsaInput)).ByRsa([]byte(pkcs1PublicKey), SHA256)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v2.ToBool())
	v3 := Verify.FromBase64Bytes(s.ToBase64Bytes(), []byte(rsaInput)).ByRsa([]byte(pkcs1PublicKey), SHA256)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v3.ToBool())
}

func TestRsa_PKCS8_String(t *testing.T) {
	e := Encrypt.FromString(rsaInput).ByRsa(pkcs8PublicKey)
	assert.Nil(t, e.Error)
	d1 := Decrypt.FromRawString(e.ToRawString()).ByRsa(pkcs8PrivateKey)
	assert.Nil(t, d1.Error)
	assert.Equal(t, rsaInput, d1.ToString())
	d2 := Decrypt.FromHexString(e.ToHexString()).ByRsa(pkcs8PrivateKey)
	assert.Nil(t, d2.Error)
	assert.Equal(t, rsaInput, d2.ToString())
	d3 := Decrypt.FromBase64String(e.ToBase64String()).ByRsa(pkcs8PrivateKey)
	assert.Nil(t, d3.Error)
	assert.Equal(t, rsaInput, d3.ToString())

	s := Sign.FromString(rsaInput).ByRsa(pkcs8PrivateKey, SHA384)
	assert.Nil(t, s.Error)
	v1 := Verify.FromRawString(s.ToRawString(), "").ByRsa(pkcs8PublicKey, SHA224)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v1.ToBool())
	v2 := Verify.FromHexString(s.ToHexString(), "").ByRsa(pkcs8PublicKey, SHA224)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v2.ToBool())
	v3 := Verify.FromBase64String(s.ToBase64String(), rsaInput).ByRsa(pkcs8PublicKey, SHA384)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v3.ToBool())
}

func TestRsa_PKCS8_Bytes(t *testing.T) {
	e := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs8PublicKey))
	assert.Nil(t, e.Error)
	d1 := Decrypt.FromRawBytes(e.ToRawBytes()).ByRsa([]byte(pkcs8PrivateKey))
	assert.Nil(t, d1.Error)
	assert.Equal(t, []byte(rsaInput), d1.ToBytes())
	d2 := Decrypt.FromHexBytes(e.ToHexBytes()).ByRsa([]byte(pkcs8PrivateKey))
	assert.Nil(t, d2.Error)
	assert.Equal(t, []byte(rsaInput), d2.ToBytes())
	d3 := Decrypt.FromBase64Bytes(e.ToBase64Bytes()).ByRsa([]byte(pkcs8PrivateKey))
	assert.Nil(t, d3.Error)
	assert.Equal(t, []byte(rsaInput), d3.ToBytes())

	s := Sign.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs8PrivateKey), SHA256)
	assert.Nil(t, s.Error)
	v1 := Verify.FromRawBytes(s.ToRawBytes(), []byte("")).ByRsa([]byte(pkcs8PublicKey), MD5)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v1.ToBool())
	v2 := Verify.FromHexBytes(s.ToHexBytes(), []byte(rsaInput)).ByRsa([]byte(pkcs8PublicKey), SHA256)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v2.ToBool())
	v3 := Verify.FromBase64Bytes(s.ToBase64Bytes(), []byte(rsaInput)).ByRsa([]byte(pkcs8PublicKey), SHA256)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v3.ToBool())
}

func TestRsa_Empty_String(t *testing.T) {
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

	s := Sign.FromString(empty).ByRsa(pkcs1PrivateKey, SHA384)
	assert.Nil(t, s.Error)
	v1 := Verify.FromRawString(s.ToRawString(), empty).ByRsa(pkcs1PublicKey, SHA224)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v1.ToBool())
	v2 := Verify.FromHexString(s.ToHexString(), empty).ByRsa(pkcs1PublicKey, SHA224)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v2.ToBool())
	v3 := Verify.FromBase64String(s.ToBase64String(), empty).ByRsa(pkcs1PublicKey, SHA384)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v3.ToBool())
}

func TestRsa_Empty_Bytes(t *testing.T) {
	empty := ""
	e := Encrypt.FromBytes([]byte(empty)).ByRsa([]byte(pkcs1PublicKey))
	assert.Nil(t, e.Error)
	d1 := Decrypt.FromRawBytes(e.ToRawBytes()).ByRsa([]byte(pkcs1PrivateKey))
	assert.Nil(t, d1.Error)
	assert.Equal(t, []byte(empty), d1.ToBytes())
	d2 := Decrypt.FromHexBytes(e.ToHexBytes()).ByRsa([]byte(pkcs1PrivateKey))
	assert.Nil(t, d2.Error)
	assert.Equal(t, []byte(empty), d2.ToBytes())
	d3 := Decrypt.FromBase64Bytes(e.ToBase64Bytes()).ByRsa([]byte(pkcs1PrivateKey))
	assert.Nil(t, d3.Error)
	assert.Equal(t, []byte(empty), d3.ToBytes())

	s := Sign.FromBytes([]byte(empty)).ByRsa([]byte(pkcs1PrivateKey), SHA256)
	assert.Nil(t, s.Error)
	v1 := Verify.FromRawBytes(s.ToRawBytes(), []byte(empty)).ByRsa([]byte(pkcs1PublicKey), MD5)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v1.ToBool())
	v2 := Verify.FromHexBytes(s.ToHexBytes(), []byte(empty)).ByRsa([]byte(pkcs1PublicKey), SHA256)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v2.ToBool())
	v3 := Verify.FromBase64Bytes(s.ToBase64Bytes(), []byte(empty)).ByRsa([]byte(pkcs1PublicKey), SHA256)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v3.ToBool())
}

func TestRsa_PublicKey_Error(t *testing.T) {
	e1 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte("xxxx"))
	assert.Equal(t, invalidRsaPublicKeyError(), e1.Error)
	e2 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(`-----BEGIN PUBLIC KEY-----
xxxx
-----END PUBLIC KEY-----`))
	assert.Equal(t, invalidRsaPublicKeyError(), e2.Error)
	e3 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(`-----BEGIN RSA PUBLIC KEY-----
	xxxx
	-----END RSA PUBLIC KEY-----`))
	assert.Equal(t, invalidRsaPublicKeyError(), e3.Error)
	e4 := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs8PrivateKey))
	assert.Equal(t, invalidRsaPublicKeyError(), e4.Error)

	v1 := Verify.FromRawBytes([]byte("xxxx"), []byte(rsaInput)).ByRsa([]byte("xxxx"), SHA512)
	assert.Equal(t, invalidRsaPublicKeyError(), v1.Error)
	v2 := Verify.FromRawBytes([]byte("xxxx"), []byte(rsaInput)).ByRsa([]byte(`-----BEGIN PUBLIC KEY-----
xxxx
-----END PUBLIC KEY-----`), SHA512)
	assert.Equal(t, invalidRsaPublicKeyError(), v2.Error)
	v3 := Verify.FromRawBytes([]byte("xxxx"), []byte(rsaInput)).ByRsa([]byte(`-----BEGIN RSA PUBLIC KEY-----
	xxxx
	-----END RSA PUBLIC KEY-----`), SHA512)
	assert.Equal(t, invalidRsaPublicKeyError(), v3.Error)
	v4 := Verify.FromRawBytes([]byte("xxxx"), []byte(rsaInput)).ByRsa([]byte(pkcs8PrivateKey), SHA512)
	assert.Equal(t, invalidRsaPublicKeyError(), v4.Error)

	v5 := Verify.FromHexString("xxxx", rsaInput).ByRsa(pkcs8PublicKey, SHA512)
	assert.Equal(t, invalidDecodingError("hex"), v5.Error)
	v6 := Verify.FromHexBytes([]byte("xxxx"), []byte(rsaInput)).ByRsa(pkcs8PublicKey, SHA512)
	assert.Equal(t, invalidDecodingError("hex"), v6.Error)
	v7 := Verify.FromBase64String("xxxxxx", rsaInput).ByRsa(pkcs8PublicKey, SHA512)
	assert.Equal(t, invalidDecodingError("base64"), v7.Error)
	v8 := Verify.FromBase64Bytes([]byte("xxxxxx"), []byte(rsaInput)).ByRsa(pkcs8PublicKey, SHA512)
	assert.Equal(t, invalidDecodingError("base64"), v8.Error)
}

func TestRsa_PrivateKey_Error(t *testing.T) {
	e := Encrypt.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs1PublicKey))

	d1 := Decrypt.FromHexBytes(e.ToHexBytes()).ByRsa([]byte("xxxx"))
	assert.Equal(t, invalidRsaPrivateKeyError(), d1.Error)
	d2 := Decrypt.FromBase64Bytes(e.ToBase64Bytes()).ByRsa([]byte(`-----BEGIN RSA PRIVATE KEY-----
	xxxx
	-----END RSA PRIVATE KEY-----`))
	assert.Equal(t, invalidRsaPrivateKeyError(), d2.Error)
	d3 := Decrypt.FromBase64Bytes(e.ToBase64Bytes()).ByRsa([]byte(`-----BEGIN PRIVATE KEY-----
	xxxx
	-----END PRIVATE KEY-----`))
	assert.Equal(t, invalidRsaPrivateKeyError(), d3.Error)
	d4 := Decrypt.FromHexBytes(e.ToHexBytes()).ByRsa([]byte(pkcs8PrivateKey))
	assert.Equal(t, invalidRsaPrivateKeyError(), d4.Error)

	s1 := Sign.FromBytes([]byte(rsaInput)).ByRsa([]byte("xxxx"), SHA1)
	assert.Equal(t, invalidRsaPrivateKeyError(), s1.Error)
	s2 := Sign.FromBytes([]byte(rsaInput)).ByRsa([]byte(`-----BEGIN RSA PRIVATE KEY-----
	xxxx
	-----END RSA PRIVATE KEY-----`), SHA1)
	assert.Equal(t, invalidRsaPrivateKeyError(), s2.Error)
	s3 := Sign.FromBytes([]byte(rsaInput)).ByRsa([]byte(`-----BEGIN PRIVATE KEY-----
	xxxx
	-----END PRIVATE KEY-----`), SHA1)
	assert.Equal(t, invalidRsaPrivateKeyError(), s3.Error)
	s4 := Sign.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs8PublicKey), SHA1)
	assert.Equal(t, invalidRsaPrivateKeyError(), s4.Error)
}

func TestRsa_Hash_Error(t *testing.T) {
	s := Sign.FromBytes([]byte(rsaInput)).ByRsa([]byte(pkcs1PrivateKey), MD4)
	assert.Equal(t, invalidRsaHashError(), s.Error)

	v := Verify.FromRawBytes(s.ToRawBytes(), []byte(rsaInput)).ByRsa([]byte(pkcs8PublicKey), MD4)
	assert.Equal(t, invalidRsaHashError(), v.Error)
}
