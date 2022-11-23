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
	pub := getRsaPublicKey(publicKey)
	if pub == nil {
		e.Error = invalidPublicKeyError()
		return e
	}
	e.dst, e.Error = rsa.EncryptPKCS1v15(rand.Reader, pub, e.src)
	return e
}

// ByRsa decrypts by rsa with private key.
// 通过 rsa 私钥解密
func (d decrypt) ByRsa(privateKey interface{}, format string) decrypt {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	pri := getRsaPrivateKey(privateKey, format)
	if pri == nil {
		d.Error = invalidPrivateKeyError()
		return d
	}
	d.dst, d.Error = rsa.DecryptPKCS1v15(rand.Reader, pri, d.src)
	return d
}

// getRsaPublicKey gets rsa public key pointer.
// 获取 rsa 公钥
func getRsaPublicKey(publicKey interface{}) *rsa.PublicKey {
	var key []byte
	switch v := publicKey.(type) {
	case string:
		key = string2bytes(v)
	case []byte:
		key = v
	}
	block, _ := pem.Decode(key)
	if block == nil {
		return nil
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	return pub.(*rsa.PublicKey)
}

// getRsaPrivateKey gets rsa private key pointer.
// 获取 rsa 私钥
func getRsaPrivateKey(privateKey interface{}, format string) *rsa.PrivateKey {
	var key []byte
	switch v := privateKey.(type) {
	case string:
		key = string2bytes(v)
	case []byte:
		key = v
	}
	block, _ := pem.Decode(key)
	if block == nil {
		return nil
	}

	if format == PKCS1 {
		pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil
		}
		return pri
	}

	if format == PKCS8 {
		pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil
		}
		return pri.(*rsa.PrivateKey)
	}
	return nil
}
