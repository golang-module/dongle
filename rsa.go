package dongle

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// defines rsa hash enum type.
// 定义 rsa 哈希算法枚举类型
type rsaHash crypto.Hash

// rsa hash constants
// rsa 哈希算法枚举值
const (
	MD4 rsaHash = 1 + iota
	MD5
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	RIPEMD160 = 19
)

// ByRsa encrypts by rsa with public key.
// 通过 rsa 公钥加密
func (e Encrypter) ByRsa(publicKey interface{}) Encrypter {
	if len(e.src) == 0 {
		return e
	}
	pub, err := parseRsaPublicKey(publicKey)
	if err != nil {
		e.Error = invalidRsaPublicKeyError()
		return e
	}
	buffer := bytes.NewBufferString("")
	for _, chunk := range split(e.src, pub.Size()-11) {
		e.dst, e.Error = rsa.EncryptPKCS1v15(rand.Reader, pub, chunk)
		buffer.Write(e.dst)
	}
	e.dst = buffer.Bytes()
	return e
}

// ByRsa decrypts by rsa with private key.
// 通过 rsa 私钥解密
func (d Decrypter) ByRsa(privateKey interface{}) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	pri, err := parseRsaPrivateKey(privateKey)
	if err != nil {
		d.Error = invalidRsaPrivateKeyError()
		return d
	}
	buffer := bytes.NewBufferString("")
	for _, chunk := range split(d.src, pri.Size()) {
		d.dst, d.Error = rsa.DecryptPKCS1v15(rand.Reader, pri, chunk)
		if d.Error != nil {
			d.Error = invalidRsaPrivateKeyError()
			return d
		}
		buffer.Write(d.dst)
	}
	d.dst = buffer.Bytes()
	return d
}

// ByRsa signs by rsa.
// 通过 rsa 私钥签名
func (s Signer) ByRsa(privateKey interface{}, hash rsaHash) Signer {
	if len(s.src) == 0 || s.Error != nil {
		return s
	}
	if !hash.isSupported() {
		s.Error = invalidRsaHashError()
		return s
	}
	pri, err := parseRsaPrivateKey(privateKey)
	if err != nil {
		s.Error = err
		return s
	}
	hasher := crypto.Hash(hash).New()
	hasher.Write(s.src)
	s.dst, s.Error = rsa.SignPKCS1v15(rand.Reader, pri, crypto.Hash(hash), hasher.Sum(nil))
	return s
}

// ByRsa verify sign by rsa with public key.
// 通过 rsa 公钥验签
func (v Verifier) ByRsa(publicKey interface{}, hash rsaHash) Verifier {
	if len(v.src) == 0 || v.Error != nil {
		return v
	}
	if !hash.isSupported() {
		v.Error = invalidRsaHashError()
		return v
	}
	pub, err := parseRsaPublicKey(publicKey)
	if err != nil {
		v.Error = err
		return v
	}
	hasher := crypto.Hash(hash).New()
	hasher.Write(v.src)
	v.Error = rsa.VerifyPKCS1v15(pub, crypto.Hash(hash), hasher.Sum(nil), v.sign)
	return v
}

// parseRsaPublicKey parses rsa public key.
// 解析 rsa 公钥
func parseRsaPublicKey(publicKey interface{}) (*rsa.PublicKey, error) {
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

// parseRsaPrivateKey parses rsa private key.
// 解析 rsa 私钥
func parseRsaPrivateKey(privateKey interface{}) (*rsa.PrivateKey, error) {
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

// whether is a rsa supported hash algorithm
// 判断是否是 rsa 支持的哈希算法
func (hash rsaHash) isSupported() bool {
	hashes := []rsaHash{
		MD5, SHA1, SHA224, SHA256, SHA384, SHA512, RIPEMD160,
	}
	for _, algo := range hashes {
		if algo == hash {
			return true
		}
	}
	return false
}

// split the string by the specified size
// 按照指定长度分割字符串
func split(buf []byte, size int) [][]byte {
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
