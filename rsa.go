package dongle

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// ByRsa encrypts by rsa with public key.
// 通过 rsa 公钥加密
func (e encrypt) ByRsa(publicKey interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var key []byte
	switch v := publicKey.(type) {
	case string:
		key = string2bytes(v)
	case []byte:
		key = v
	}
	block, _ := pem.Decode(key)
	if block == nil {
		e.Error = invalidPublicKeyError()
		return e
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		e.Error = invalidPublicKeyError()
		return e
	}
	e.dst, e.Error = rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), e.src)
	return e
}

// ByRsa decrypts by rsa with private key.
// 通过 rsa 私钥解密
func (d decrypt) ByRsa(privateKey interface{}, format ...string) decrypt {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	var key []byte
	switch v := privateKey.(type) {
	case string:
		key = string2bytes(v)
	case []byte:
		key = v
	}
	block, _ := pem.Decode(key)
	if block == nil {
		d.Error = invalidPrivateKeyError()
		return d
	}

	if len(format) == 0 {
		format = []string{PKCS8}
	}

	if format[0] == PKCS1 {
		pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			d.Error = invalidPrivateKeyError()
			return d
		}
		d.dst, d.Error = rsa.DecryptPKCS1v15(rand.Reader, pri, d.src)
	}

	if format[0] == PKCS8 {
		pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			d.Error = invalidPrivateKeyError()
			return d
		}
		d.dst, d.Error = rsa.DecryptPKCS1v15(rand.Reader, pri.(*rsa.PrivateKey), d.src)
	}

	return d
}
