package openssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

var (
	// returns an invalid RSA public key error.
	// 返回无效的 RSA 公钥错误
	invalidRsaPublicKeyError = func() error {
		return fmt.Errorf("openssl/rsa: invalid rsa public key, please make sure the public key is valid")
	}
	// returns an invalid RSA private key error.
	// 返回无效的 RSA 私钥错误
	invalidRSAPrivateKeyError = func() error {
		return fmt.Errorf("openssl/rsa: invalid rsa private key, please make sure the private key is valid")
	}
)

// rsaKeyPair defines a rsaKeyPair struct.
// 定义 rsaKeyPair 结构体
type rsaKeyPair struct {
	publicKey  []byte
	privateKey []byte
}

// newRsaKeyPair returns a new rsaKeyPair instance.
// 初始化 rsaKeyPair 结构体
func newRsaKeyPair() *rsaKeyPair {
	return &rsaKeyPair{}
}

// GenPKCS1KeyPair generates PKCS#1 key pair.
// 生成 PKCS#1 格式密钥对
func (r rsaKeyPair) GenPKCS1KeyPair(bits int) (publicKey, privateKey []byte) {
	pri, _ := rsa.GenerateKey(rand.Reader, bits)
	privateBytes := x509.MarshalPKCS1PrivateKey(pri)
	privateKey = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateBytes,
	})
	publicBytes := x509.MarshalPKCS1PublicKey(&pri.PublicKey)
	publicKey = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicBytes,
	})
	return
}

// GenPKCS8KeyPair generates PKCS#8 key pair.
// 生成 PKCS#8 格式密钥对
func (r rsaKeyPair) GenPKCS8KeyPair(bits int) (publicKey, privateKey []byte) {
	pri, _ := rsa.GenerateKey(rand.Reader, bits)

	privateBytes, _ := x509.MarshalPKCS8PrivateKey(pri)
	privateKey = pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateBytes,
	})
	publicBytes, _ := x509.MarshalPKIXPublicKey(&pri.PublicKey)
	publicKey = pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicBytes,
	})
	return
}

// VerifyKeyPair verify key pair matches.
// 验证密钥对是否匹配
func (r rsaKeyPair) VerifyKeyPair(publicKey, privateKey []byte) bool {
	pub, err := r.ExportPublicKey(privateKey)
	if err != nil {
		return false
	}

	if string(publicKey) == string(pub) {
		return true
	}
	return false
}

// IsPublicKey whether is a public key.
// 是否是公钥
func (r rsaKeyPair) IsPublicKey(publicKey []byte) bool {
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
func (r rsaKeyPair) IsPrivateKey(privateKey []byte) bool {
	block, _ := pem.Decode(privateKey)
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

// ParsePublicKey parses public key.
// 解析公钥
func (r rsaKeyPair) ParsePublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, invalidRsaPublicKeyError()
	}
	if block.Type == "RSA PUBLIC KEY" {
		pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, invalidRsaPublicKeyError()
		}
		return pub, nil
	}
	if block.Type == "PUBLIC KEY" {
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, invalidRsaPublicKeyError()
		}
		return pub.(*rsa.PublicKey), err
	}
	return nil, invalidRsaPublicKeyError()
}

// ParsePrivateKey parses private key.
// 解析私钥
func (r rsaKeyPair) ParsePrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, invalidRSAPrivateKeyError()
	}
	if block.Type == "RSA PRIVATE KEY" {
		pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, invalidRSAPrivateKeyError()
		}
		return pri, err
	}
	if block.Type == "PRIVATE KEY" {
		pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, invalidRSAPrivateKeyError()
		}
		return pri.(*rsa.PrivateKey), err
	}
	return nil, invalidRSAPrivateKeyError()
}

// ExportPublicKey exports public key from private key.
// 从私钥里导出公钥
func (r rsaKeyPair) ExportPublicKey(privateKey []byte) (publicKey []byte, err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = invalidRSAPrivateKeyError()
		return
	}

	pri, err := r.ParsePrivateKey(privateKey)
	if err != nil {
		return
	}

	blockType, blockBytes := "", []byte("")
	if block.Type == "RSA PRIVATE KEY" {
		blockType = "RSA PUBLIC KEY"
		blockBytes = x509.MarshalPKCS1PublicKey(&pri.PublicKey)
	}
	if block.Type == "PRIVATE KEY" {
		blockType = "PUBLIC KEY"
		blockBytes, err = x509.MarshalPKIXPublicKey(&pri.PublicKey)
	}
	publicKey = pem.EncodeToMemory(&pem.Block{
		Type:  blockType,
		Bytes: blockBytes,
	})
	return
}
