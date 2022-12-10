package dongle

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// ByRsa encrypts by rsa with public key.
// 通过 rsa 公钥加密
func (e encrypter) ByRsa(publicKey interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	pub, err := ParseRsaPublicKey(publicKey)
	if err != nil {
		e.Error = invalidRsaPublicKeyError()
		return e
	}
	e.dst, e.Error = rsa.EncryptPKCS1v15(rand.Reader, pub, e.src)
	return e
}

// ByRsa decrypts by rsa with private key.
// 通过 rsa 私钥解密
func (d decrypter) ByRsa(privateKey interface{}) decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	pri, err := ParseRsaPrivateKey(privateKey)
	if err != nil {
		d.Error = invalidRsaPrivateKeyError()
		return d
	}
	d.dst, d.Error = rsa.DecryptPKCS1v15(rand.Reader, pri, d.src)
	if d.Error != nil {
		d.Error = invalidRsaPrivateKeyError()
		return d
	}
	return d
}

// ByRsa signs by rsa.
// 通过 rsa 私钥签名
func (s Signer) ByRsa(privateKey interface{}, hash crypto.Hash) Signer {
	if len(s.src) == 0 || s.Error != nil {
		return s
	}
	if !isSupportedHash(hash) {
		s.Error = invalidRsaHashError()
		return s
	}
	pri, err := ParseRsaPrivateKey(privateKey)
	if err != nil {
		s.Error = err
		return s
	}
	hasher := hash.New()
	hasher.Write(s.src)
	s.dst, s.Error = rsa.SignPKCS1v15(rand.Reader, pri, hash, hasher.Sum(nil))
	return s
}

// ByRsa verify sign by rsa with public key.
// 通过 rsa 公钥验签
func (v Verifier) ByRsa(publicKey interface{}, hash crypto.Hash) Verifier {
	if len(v.src) == 0 || v.Error != nil {
		return v
	}
	if !isSupportedHash(hash) {
		v.Error = invalidRsaHashError()
		return v
	}
	pub, err := ParseRsaPublicKey(publicKey)
	if err != nil {
		v.Error = err
		return v
	}
	hasher := hash.New()
	hasher.Write(v.src)
	v.Error = rsa.VerifyPKCS1v15(pub, hash, hasher.Sum(nil), v.sign)
	return v
}

// ParseRsaPublicKey parses rsa public key.
// 解析 rsa 公钥
func ParseRsaPublicKey(publicKey interface{}) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(interface2bytes(publicKey))
	if block == nil {
		return nil, invalidRsaPublicKeyError()
	}
	// pkcs1 pem
	if block.Type == "RSA PUBLIC KEY" {
		return x509.ParsePKCS1PublicKey(block.Bytes)
	}
	// pkcs8 pem
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, invalidRsaPublicKeyError()
	}
	return pub.(*rsa.PublicKey), err
}

// ParseRsaPrivateKey parses rsa private key.
// 解析 rsa 私钥
func ParseRsaPrivateKey(privateKey interface{}) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(interface2bytes(privateKey))
	if block == nil {
		return nil, invalidRsaPrivateKeyError()
	}
	// pkcs1 pem
	if block.Type == "RSA PRIVATE KEY" {
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	}
	// pkcs8 pem
	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, invalidRsaPrivateKeyError()
	}
	return pri.(*rsa.PrivateKey), nil
}

// whether is a supported hash function
// 判断是否是支持的哈希函数
func isSupportedHash(hash crypto.Hash) bool {
	hashes := []crypto.Hash{
		MD5, SHA1, SHA224, SHA256, SHA384, SHA512, RIPEMD160,
	}
	for _, e := range hashes {
		if e == hash {
			return true
		}
	}
	return false
}
