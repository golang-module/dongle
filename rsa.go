package dongle

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type pkcsVersion string

// returns an invalid public key error
// 返回无效的公钥错误
var invalidPublicKeyError = func() error {
	return fmt.Errorf("invalid public key, please make sure the public key is valid")
}

// returns an invalid private key error
// 返回无效的私钥错误
var invalidPrivateKeyError = func() error {
	return fmt.Errorf("invalid private key, please make sure the private key is valid")
}

// pkcs version constants
// 证书版本常量
const (
	PKCS1 pkcsVersion = "pkcs1"
	PKCS8 pkcsVersion = "pkcs8"
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
func (d decrypt) ByRsa(privateKey interface{}, version pkcsVersion) decrypt {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	pri := getRsaPrivateKey(privateKey, version)
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
	var keyBytes []byte
	switch v := publicKey.(type) {
	case string:
		keyBytes = string2bytes(v)
	case []byte:
		keyBytes = v
	}
	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
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
func getRsaPrivateKey(privateKey interface{}, version pkcsVersion) *rsa.PrivateKey {
	var keyBytes []byte
	switch v := privateKey.(type) {
	case string:
		keyBytes = string2bytes(v)
	case []byte:
		keyBytes = v
	}
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil
	}

	if version == PKCS1 {
		pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil
		}
		return pri
	}

	if version == PKCS8 {
		pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil
		}
		return pri.(*rsa.PrivateKey)
	}
	return nil
}
