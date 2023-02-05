package openssl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	publicKey1  []byte
	privateKey1 []byte
	publicKey8  []byte
	privateKey8 []byte
)

func TestRsa_GenKeyPair(t *testing.T) {
	publicKey1, privateKey1 = RSA.GenKeyPair(PKCS1, 1024)
	fmt.Printf("生成 1024 字节 PKCS#1 格式公钥：\n%s", publicKey1)
	fmt.Printf("生成 1024 字节 PKCS#1 格式私钥：\n%s", privateKey1)

	publicKey8, privateKey8 = RSA.GenKeyPair(PKCS8, 2048)
	fmt.Printf("生成 2048 字节 PKCS#8 格式公钥：\n%s", publicKey8)
	fmt.Printf("生成 2048 字节 PKCS#8 格式私钥：\n%s", privateKey8)
}

func TestRsa_VerifyKeyPair(t *testing.T) {
	publicKey1, privateKey1 = RSA.GenKeyPair(PKCS1, 1024)
	publicKey8, privateKey8 = RSA.GenKeyPair(PKCS8, 2048)

	assert.Equal(t, true, RSA.VerifyKeyPair(publicKey1, privateKey1))
	assert.Equal(t, true, RSA.VerifyKeyPair(publicKey8, privateKey8))
	assert.Equal(t, false, RSA.VerifyKeyPair(publicKey1, privateKey8))
	assert.Equal(t, false, RSA.VerifyKeyPair(publicKey8, privateKey1))
	assert.Equal(t, false, RSA.VerifyKeyPair(publicKey1, []byte("xxx")))
	assert.Equal(t, false, RSA.VerifyKeyPair([]byte("xxx"), privateKey1))
}

func TestRsa_IsPublicKey(t *testing.T) {
	publicKey1, privateKey1 = RSA.GenKeyPair(PKCS1, 1024)
	publicKey8, privateKey8 = RSA.GenKeyPair(PKCS8, 2048)

	assert.Equal(t, true, RSA.IsPublicKey(publicKey1))
	assert.Equal(t, true, RSA.IsPrivateKey(privateKey1))
	assert.Equal(t, false, RSA.IsPublicKey(privateKey1))
	assert.Equal(t, false, RSA.IsPrivateKey(publicKey1))

	assert.Equal(t, true, RSA.IsPublicKey(publicKey8))
	assert.Equal(t, true, RSA.IsPrivateKey(privateKey8))
	assert.Equal(t, false, RSA.IsPublicKey(privateKey8))
	assert.Equal(t, false, RSA.IsPrivateKey(publicKey8))

	assert.Equal(t, false, RSA.IsPublicKey([]byte("xxx")))
	assert.Equal(t, false, RSA.IsPrivateKey([]byte("xxx")))
}

func TestRsa_FormatPublicKey(t *testing.T) {
	publicKey1 := RSA.FormatPublicKey(PKCS1, []byte("MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEwcujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFndk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE="))
	fmt.Printf("格式化后的 PKCS#1 格式公钥：\n%s", publicKey1)

	publicKey8 := RSA.FormatPublicKey(PKCS8, []byte("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHqX1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJy4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMezHC1outlM6x+/BB0BSQIDAQAB"))
	fmt.Printf("格式化后的 PKCS#8 格式公钥：\n%s", publicKey8)
}

func TestRsa_FormatPrivateKey(t *testing.T) {
	privateKey1 := RSA.FormatPrivateKey(PKCS1, []byte("MIICXQIBAAKBgQCtdjE3fOKpAmc6eIi1I/GElIJWhTLBZb/SpSC+Pl7bONDyt/sG7FzxMHLoxnlJnYDHmYiu0iy0s/EHhNL+1bPnBiuAbOKi7TBgojusBLsrddCrGn08gZLRZ3ZP/oAS6+wf2/vfh2jCCKCVHbBZLHQIN3MU0qcFQKXvfJaBTUX6gwIDAQABAoGAFwAfEo56t5JcAcLNzccQVVYj2jkbO820G8hNiSxYA5WLD0QaAxcAU/Lqqbb3ii1aUB0ppJS13NgnU6nnGGdZzUYBG1Hai6EkVyCGrI4amQ93AaVdKncL8gJ4RZAmYzPPUwSMEESsu24pS1NF1G1Y8C+28b/Wr0oqOsCvL6PhsMECQQDwsPJJoWRx7ZJwE1K5KLT0cXKyrIpyXY3I6tyA5imCzOzccf3d1vDgB0L9sdSO7bG3ceSwpAeiWEbg5jGZemPzAkEAuH6U4pEI4AMbWnatpK55Rc235NDgmT3VyIuRaKC02YXAZ+jznFepXMd4DTli4R9r3j2YVhUpyDVbdQpFH98DMQJAQpOvcU6DSkA80WOG7lCkPTlkUKgJY7kdDwZoF/+SW+vzWMbvQf3CgzV/Ak2+TgrRrbyDVZkJw45HjM4fyiRgoQJBALH//qlxgPyQQs3O/s2KQBsm1auAE5IF5MLuVUZ69sF/mBko2hEXSqHnGV645TuKU0pCZz12ga9WO3z6gaK0SaECQQDah1pKt9ViBBy4USXK3OWXEloHuTwmyr9AbLqqI5tQ2eNuH0NkuJYQmnXmHLbKOELoYocldEBXmkzPXSN+X9kV"))
	fmt.Printf("格式化后的 PKCS#1 格式私钥：\n%s", privateKey1)

	privateKey8 := RSA.FormatPrivateKey(PKCS8, []byte("MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAKrNk1r1Wtx7DJTrAOhXtj2QAepfVUrQHdFvoY2ZB7jMsR9x7txVNoutzhUZMqXfm0AMbVxEeq1obhL9a22mIZkGHEnLgyk5dvp4g+JUuyfaUv6smjld1tKveDKPEQ5BD3uKG3DiUN3nAyjhsg67DUu0x7McLWi62UzrH78EHQFJAgMBAAECgYAeo3nHWzPNURVUsUMcan96U5bEYA2AugxfQVMNf2HvOGidZ2adh3udWrQY/MglERNcTd5gKriG2rDEH0liBecIrNKsBL4lV+qHEGRUcnDDdtUBdGInEU8lve5keDgmX+/huXSRJ+3tYA5u9j+32RquVczvIdtb5XnBLUl61k0osQJBAON5+eJjtw6xpn+pveU92BSHvaJYVyrLHwUjR07aNKb7GlGVM3MGf1FCa8WQUo9uUzYxGLtg5Qf3sqwOrwPd5UsCQQDAOF/zWqGuY3HfV/1wgiXiWp8rc+S8tanMj5M37QQbYW5YLjUmJImoklVahv3qlgLZdEN5ZSueM5jfoSFtNts7AkBKoRDvSiGbi4MBbTHkzLZgfewkH/FxE7S4nctePk553fXTgCyh9ya8BRuQdHnxnpNkOxVPHEnnpEcVFbgrf5gjAkB7KmRI4VTiEfRgINhTJAG0VU7SH/N7+4cufPzfA+7ywG5c8Fa79wOB0SoB1KeUjcSLo5Ssj2fwea1F9dAeU90LAkBJQFofveaDa3YlN4EQZOcCvJKmg7xwWuGxFVTZDVVEws7UCQbEOEEXZrNd9x0IF5kpPLR+rxuaRPgUNaDGIh5o"))
	fmt.Printf("格式化后的 PKCS#8 格式私钥：\n%s", privateKey8)
}

func TestRsa_ParsePublicKey(t *testing.T) {
	publicKey1, privateKey1 = RSA.GenKeyPair(PKCS1, 1024)
	publicKey8, privateKey8 = RSA.GenKeyPair(PKCS8, 2048)

	pub1, err1 := RSA.ParsePublicKey(publicKey1)
	assert.Nil(t, err1)
	pri1, err1 := RSA.ParsePrivateKey(privateKey1)
	assert.Nil(t, err1)
	assert.Equal(t, pub1, &pri1.PublicKey)

	pub8, err8 := RSA.ParsePublicKey(publicKey8)
	assert.Nil(t, err8)
	pri8, err8 := RSA.ParsePrivateKey(privateKey8)
	assert.Nil(t, err8)
	assert.Equal(t, pub8, &pri8.PublicKey)
}

func TestRsa_ExportPublicKey(t *testing.T) {
	publicKey1, privateKey1 = RSA.GenKeyPair(PKCS1, 1024)
	publicKey8, privateKey8 = RSA.GenKeyPair(PKCS8, 2048)

	actual1, err1 := RSA.ExportPublicKey(privateKey1)
	assert.Nil(t, err1)
	assert.Equal(t, publicKey1, actual1)

	actual8, err8 := RSA.ExportPublicKey(privateKey8)
	assert.Nil(t, err8)
	assert.Equal(t, publicKey8, actual8)
}

func TestRsa_CompressKey(t *testing.T) {
	key, err := RSA.CompressKey([]byte(`-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEw
cujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFn
dk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE=
-----END RSA PUBLIC KEY-----`))
	assert.Nil(t, err)
	fmt.Printf("压缩后的密钥：\n%s\n", string(key))
}

func TestError_ParsePublicKey(t *testing.T) {
	publicKey1, privateKey1 = RSA.GenKeyPair(PKCS1, 1024)
	publicKey8, privateKey8 = RSA.GenKeyPair(PKCS8, 2048)

	_, err := RSA.ParsePublicKey(privateKey1)
	assert.Equal(t, invalidRsaPublicKeyError(), err)

	_, err = RSA.ParsePublicKey(privateKey8)
	assert.Equal(t, invalidRsaPublicKeyError(), err)

	_, err = RSA.ParsePublicKey([]byte("xxx"))
	assert.Equal(t, invalidRsaPublicKeyError(), err)

	_, err = RSA.ParsePublicKey([]byte(`-----BEGIN RSA PUBLIC KEY-----
xxxx
-----END RSA PUBLIC KEY-----`))
	assert.Equal(t, invalidRsaPublicKeyError(), err)

	_, err = RSA.ParsePublicKey([]byte(`-----BEGIN PUBLIC KEY-----
xxxx
-----END PUBLIC KEY-----`))
	assert.Equal(t, invalidRsaPublicKeyError(), err)
}

func TestError_ParsePrivateKey(t *testing.T) {
	publicKey1, privateKey1 = RSA.GenKeyPair(PKCS1, 1024)
	publicKey8, privateKey8 = RSA.GenKeyPair(PKCS8, 2048)

	_, err := RSA.ParsePrivateKey(publicKey1)
	assert.Equal(t, invalidRSAPrivateKeyError(), err)

	_, err = RSA.ParsePrivateKey(publicKey8)
	assert.Equal(t, invalidRSAPrivateKeyError(), err)

	_, err = RSA.ParsePrivateKey([]byte("xxx"))
	assert.Equal(t, invalidRSAPrivateKeyError(), err)

	_, err = RSA.ParsePrivateKey([]byte(`-----BEGIN RSA PRIVATE KEY-----
xxxx
-----END RSA PRIVATE KEY-----`))
	assert.Equal(t, invalidRSAPrivateKeyError(), err)

	_, err = RSA.ParsePrivateKey([]byte(`-----BEGIN PRIVATE KEY-----
xxxx
-----END PRIVATE KEY-----`))
	assert.Equal(t, invalidRSAPrivateKeyError(), err)
}

func TestError_ExportPublicKey(t *testing.T) {
	_, err := RSA.ExportPublicKey([]byte("xxx"))
	assert.Equal(t, invalidRSAPrivateKeyError(), err)

	_, err = RSA.ExportPublicKey([]byte(`-----BEGIN RSA PUBLIC KEY-----
xxxx
-----END RSA PUBLIC KEY-----`))
	assert.Equal(t, invalidRSAPrivateKeyError(), err)

	_, err = RSA.ExportPublicKey([]byte(`-----BEGIN PUBLIC KEY-----
xxxx
-----END PUBLIC KEY-----`))
	assert.Equal(t, invalidRSAPrivateKeyError(), err)
}

func TestError_CompressKey(t *testing.T) {
	_, err := RSA.CompressKey([]byte("xxx"))
	assert.Equal(t, invalidRSAKeyError(), err)
}
