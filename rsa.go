package dongle

import (
	"crypto"
	"github.com/golang-module/dongle/rsa"
)

// defines rsa hash enum type.
// 定义 rsa 哈希算法枚举类型
type rsaHash crypto.Hash

// rsa hash constants
// rsa 哈希算法枚举值
const (
	MD5 rsaHash = 2 + iota
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	RIPEMD160 = 19
)

// ByRsa encrypts by rsa with public key.
// 通过 rsa 密钥加密
func (e Encrypter) ByRsa(rsaKey interface{}) Encrypter {
	if len(e.src) == 0 {
		return e
	}
	keyPair := rsa.NewKeyPair()
	keyPair.SetPublicKey(interface2bytes(rsaKey))
	keyPair.SetPrivateKey(interface2bytes(rsaKey))

	if rsa.IsPrivateKey(interface2bytes(rsaKey)) {
		e.dst, e.Error = keyPair.EncryptByPrivateKey(e.src)
		return e
	}
	e.dst, e.Error = keyPair.EncryptByPublicKey(e.src)
	return e
}

// ByRsa decrypts by rsa with private key.
// 通过 rsa 密钥解密
func (d Decrypter) ByRsa(rsaKey interface{}) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	keyPair := rsa.NewKeyPair()
	keyPair.SetPublicKey(interface2bytes(rsaKey))
	keyPair.SetPrivateKey(interface2bytes(rsaKey))

	if rsa.IsPublicKey(interface2bytes(rsaKey)) {
		d.dst, d.Error = keyPair.DecryptByPublicKey(d.src)
		return d
	}
	d.dst, d.Error = keyPair.DecryptByPrivateKey(d.src)
	return d
}

// ByRsa signs by rsa.
// 通过 rsa 私钥签名
func (s Signer) ByRsa(privateKey interface{}, hash rsaHash) Signer {
	if len(s.src) == 0 || s.Error != nil {
		return s
	}
	keyPair := rsa.NewKeyPair()
	keyPair.SetPrivateKey(interface2bytes(privateKey))

	s.dst, s.Error = keyPair.SignByPrivateKey(s.src, crypto.Hash(hash))
	return s
}

// ByRsa verify sign by rsa with public key.
// 通过 rsa 公钥验签
func (v Verifier) ByRsa(publicKey interface{}, hash rsaHash) Verifier {
	if len(v.src) == 0 || v.Error != nil {
		return v
	}
	keyPair := rsa.NewKeyPair()
	keyPair.SetPublicKey(interface2bytes(publicKey))

	v.Error = keyPair.VerifyByPublicKey(v.src, v.sign, crypto.Hash(hash))
	return v
}
