// @Package dongle
// @Description a simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption
// @Page github.com/golang-module/dongle
// @Developer gouguoyin
// @Blog www.gouguoyin.cn
// @Email contact@gouguoyin.cn

// Package dongle is a simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption.
package dongle

import (
	"unsafe"
)

// Version current version
// 当前版本号
const Version = "0.2.1"

// dongle defines a dongle struct.
// 定义 dongle 结构体
type dongle struct {
	src   []byte
	dst   []byte
	Error error
}

var (
	// Encode returns a new encoder instance
	// 返回 encoder 实例
	Encode = newEncoder()
	// Decode returns a new decoder instance
	// 返回 decoder 实例
	Decode = newDecoder()
	// Encrypt returns a new encrypter instance
	// 返回 encrypter 实例
	Encrypt = newEncrypter()
	// Decrypt returns a new decrypter instance
	// 返回 decrypter 实例
	Decrypt = newDecrypter()
	// Sign returns a new signer instance
	// 返回 signer 实例
	Sign = newSigner()
	// Verify returns a new verifier instance
	// 返回 verifier 实例
	Verify = newVerifier()
)

// converts string to byte slice without a memory allocation.
// 将字符串转换为字节切片
func string2bytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// converts byte slice to string without a memory allocation.
// 将字节切片转换为字符串
func bytes2string(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// converts interface to byte slice.
// 将接口转换为字节切片
func interface2bytes(i interface{}) (b []byte) {
	switch v := i.(type) {
	case string:
		b = string2bytes(v)
	case []byte:
		b = v
	}
	return
}

// gets Cipher instance.
// 获取 Cipher 对象
func getCipher(mode cipherMode, padding cipherPadding, key, iv interface{}) (cipher *Cipher) {
	cipher = NewCipher()
	cipher.SetMode(mode)
	cipher.SetPadding(padding)
	cipher.SetKey(key)
	cipher.SetIV(iv)
	return
}
