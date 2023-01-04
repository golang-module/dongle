// Package rsa implements crypto/rsa
package rsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
)

var emptyByte = []byte("")

var (
	// returns an invalid public key error
	// 返回无效的公钥错误
	invalidPublicKeyError = func() error {
		return fmt.Errorf("rsa: invalid public key, please make sure the public key is valid")
	}
	// returns an invalid private key error
	// 返回无效的私钥错误
	invalidPrivateKeyError = func() error {
		return fmt.Errorf("rsa: invalid private key, please make sure the private key is valid")
	}
)

// KeyPair defines a KeyPair struct.
// 定义 KeyPair 结构体
type KeyPair struct {
	publicKey  []byte
	privateKey []byte
}

// NewKeyPair returns a new KeyPair instance.
// 初始化 keyPair 结构体
func NewKeyPair() *KeyPair {
	return &KeyPair{}
}

// SetPublicKey sets public key.
// 设置公钥
func (k *KeyPair) SetPublicKey(publicKey []byte) {
	k.publicKey = publicKey
}

// SetPrivateKey sets private key.
// 设置私钥
func (k *KeyPair) SetPrivateKey(privateKey []byte) {
	k.privateKey = privateKey
}

// EncryptByPublicKey encrypts by public key.
// 通过公钥加密
func (k *KeyPair) EncryptByPublicKey(src []byte) (dst []byte, err error) {
	if len(src) == 0 {
		return emptyByte, nil
	}
	if !IsPublicKey(k.publicKey) {
		dst, err = emptyByte, invalidPublicKeyError()
		return
	}
	pub, err := k.ParsePublicKey()
	if err != nil {
		dst, err = emptyByte, invalidPublicKeyError()
		return
	}
	buffer := bytes.NewBufferString("")
	for _, chunk := range bytesSplit(src, pub.Size()-11) {
		dst, err = rsa.EncryptPKCS1v15(rand.Reader, pub, chunk)
		buffer.Write(dst)
	}
	dst = buffer.Bytes()
	return
}

// DecryptByPrivateKey encrypts by private key.
// 通过私钥解密
func (k *KeyPair) DecryptByPrivateKey(src []byte) (dst []byte, err error) {
	if len(src) == 0 {
		dst = emptyByte
		return
	}
	if !IsPrivateKey(k.privateKey) {
		dst, err = emptyByte, invalidPrivateKeyError()
		return
	}
	pri, err := k.ParsePrivateKey()
	if err != nil {
		dst, err = emptyByte, invalidPrivateKeyError()
		return
	}
	buffer := bytes.NewBufferString("")
	for _, chunk := range bytesSplit(src, pri.Size()) {
		dst, err = rsa.DecryptPKCS1v15(rand.Reader, pri, chunk)
		buffer.Write(dst)
	}
	dst = buffer.Bytes()
	return
}

// EncryptByPrivateKey encrypts by private key.
// 通过私钥加密
func (k *KeyPair) EncryptByPrivateKey(src []byte) (dst []byte, err error) {
	if len(src) == 0 {
		dst = emptyByte
		return
	}
	if !IsPrivateKey(k.privateKey) {
		dst, err = emptyByte, invalidPrivateKeyError()
		return
	}
	pri, err := k.ParsePrivateKey()
	if err != nil {
		dst, err = emptyByte, invalidPrivateKeyError()
		return
	}
	buffer := bytes.NewBufferString("")
	for _, chunk := range bytesSplit(src, pri.Size()) {
		dst, err = rsa.SignPKCS1v15(nil, pri, crypto.Hash(0), chunk)
		buffer.Write(dst)
	}
	dst = buffer.Bytes()
	return
}

// DecryptByPublicKey encrypts by public key.
// 通过公钥解密
func (k *KeyPair) DecryptByPublicKey(src []byte) (dst []byte, err error) {
	if len(src) == 0 {
		dst = emptyByte
		return
	}
	if !IsPublicKey(k.publicKey) {
		dst, err = emptyByte, invalidPublicKeyError()
		return
	}
	pub, err := k.ParsePublicKey()
	if err != nil {
		dst, err = emptyByte, invalidPublicKeyError()
		return
	}

	bigInt := new(big.Int)
	bigInt.Exp(new(big.Int).SetBytes(src), big.NewInt(int64(pub.E)), pub.N)
	dst = leftUnPad(leftPad(bigInt.Bytes(), pub.Size()))
	return
}

// SignByPrivateKey signs by private key.
// 通过私钥签名
func (k *KeyPair) SignByPrivateKey(src []byte, hash crypto.Hash) (dst []byte, err error) {
	pri, err := k.ParsePrivateKey()
	if err != nil {
		dst, err = emptyByte, invalidPrivateKeyError()
		return
	}
	hasher := hash.New()
	hasher.Write(src)
	dst, err = rsa.SignPKCS1v15(rand.Reader, pri, hash, hasher.Sum(nil))
	return
}

// VerifyByPublicKey verify by public key.
// 通过公钥验签
func (k *KeyPair) VerifyByPublicKey(src, sign []byte, hash crypto.Hash) (err error) {
	pub, err := k.ParsePublicKey()
	if err != nil {
		return
	}
	hasher := hash.New()
	hasher.Write(src)
	return rsa.VerifyPKCS1v15(pub, hash, hasher.Sum(nil), sign)
}

// ParsePublicKey parses public key.
// 解析公钥
func (k *KeyPair) ParsePublicKey() (*rsa.PublicKey, error) {
	block, _ := pem.Decode(k.publicKey)
	if block == nil {
		return nil, invalidPublicKeyError()
	}
	// pkcs1 pem
	if block.Type == "RSA PUBLIC KEY" {
		return x509.ParsePKCS1PublicKey(block.Bytes)
	}
	// pkcs8 pem
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, invalidPublicKeyError()
	}
	return pub.(*rsa.PublicKey), err
}

// ParsePrivateKey parses private key.
// 解析私钥
func (k *KeyPair) ParsePrivateKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(k.privateKey)
	if block == nil {
		return nil, invalidPrivateKeyError()
	}
	// pkcs1 pem
	if block.Type == "RSA PRIVATE KEY" {
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	}
	// pkcs8 pem
	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, invalidPrivateKeyError()
	}
	return pri.(*rsa.PrivateKey), nil
}

// IsPublicKey whether is a public key.
// 是否是公钥
func IsPublicKey(publicKey []byte) bool {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false
	}
	if block.Type == "RSA PUBLIC KEY" {
		return true
	}
	if block.Type == "PUBLIC KEY" {
		return true
	}
	return false
}

// IsPrivateKey whether is a private key.
// 是否是私钥
func IsPrivateKey(publicKey []byte) bool {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false
	}
	if block.Type == "RSA PRIVATE KEY" {
		return true
	}
	if block.Type == "PRIVATE KEY" {
		return true
	}
	return false
}

// left padding
// 左填充
func leftPad(src []byte, size int) (dst []byte) {
	dst = make([]byte, size)
	copy(dst[len(dst)-len(src):], src)
	return
}

// remove left padding
// 移出左填充
func leftUnPad(src []byte) (dst []byte) {
	n := len(src)
	t := 2
	for i := 2; i < n; i++ {
		if src[i] == 0xff {
			t = t + 1
		} else {
			if src[i] == src[0] {
				t = t + int(src[1])
			}
			break
		}
	}
	dst = make([]byte, n-t)
	copy(dst, src[t:])
	return
}

// split the byte slice by the specified size
// 按照指定长度分割字符串
func bytesSplit(buf []byte, size int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/size+1)
	for len(buf) >= size {
		chunk, buf = buf[:size], buf[size:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
