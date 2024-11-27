package dongle

import (
	"crypto"
	"fmt"

	"github.com/dromara/dongle/rsa"
)

type rsaHash crypto.Hash

const (
	MD5 rsaHash = 2 + iota
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	RIPEMD160 = 19
)

type RsaError struct {
}

func NewRsaError() RsaError {
	return RsaError{}
}

func (e RsaError) PublicKeyError() error {
	return fmt.Errorf("rsa: invalid public key, please make sure the public key is valid")
}

func (e RsaError) PrivateKeyError() error {
	return fmt.Errorf("rsa: invalid private key, please make sure the private key is valid")
}

var rsaError = NewRsaError()

// ByRsa encrypts by rsa with public key or private key.
func (e Encrypter) ByRsa(rsaKey []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	keyPair := rsa.NewKeyPair()
	keyPair.SetPublicKey(rsaKey)
	keyPair.SetPrivateKey(rsaKey)

	if keyPair.IsPrivateKey() {
		e.dst, e.Error = keyPair.EncryptByPrivateKey(e.src)
		return e
	}
	e.dst, e.Error = keyPair.EncryptByPublicKey(e.src)
	return e
}

// ByRsa decrypts by rsa with private key or public key.
func (d Decrypter) ByRsa(rsaKey []byte) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	keyPair := rsa.NewKeyPair()
	keyPair.SetPublicKey(rsaKey)
	keyPair.SetPrivateKey(rsaKey)
	if keyPair.IsPublicKey() {
		d.dst, d.Error = keyPair.DecryptByPublicKey(d.src)
		return d
	}
	d.dst, d.Error = keyPair.DecryptByPrivateKey(d.src)
	return d
}

// ByRsa signs by rsa with private key.
func (s Signer) ByRsa(privateKey []byte, hash rsaHash) Signer {
	if len(s.src) == 0 || s.Error != nil {
		return s
	}
	keyPair := rsa.NewKeyPair()
	keyPair.SetPrivateKey(privateKey)
	keyPair.SetHash(crypto.Hash(hash))

	s.dst, s.Error = keyPair.SignByPrivateKey(s.src)
	return s
}

// ByRsa verify sign by rsa with public key.
func (v Verifier) ByRsa(publicKey []byte, hash rsaHash) Verifier {
	if len(v.src) == 0 || v.Error != nil {
		return v
	}
	keyPair := rsa.NewKeyPair()
	keyPair.SetPublicKey(publicKey)
	keyPair.SetHash(crypto.Hash(hash))

	v.Error = keyPair.VerifyByPublicKey(v.src, v.sign)
	return v
}
