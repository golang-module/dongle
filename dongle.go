// @Package dongle
// @Description a simple, semantic and developer-friendly golang package for encoding&decoding, encryption&decryption and signature&verification
// @Page github.com/golang-module/dongle
// @Developer gouguoyin
// @Blog www.gouguoyin.com
// @Email contact@gouguoyin.com

// Package dongle is a simple, semantic and developer-friendly golang package for encoding&decoding, encryption&decryption and signature&verification.
package dongle

import (
	"unsafe"
)

// Version current version
// 当前版本号
const Version = "0.2.2"

// dongle defines a dongle struct.
// 定义 dongle 结构体
type dongle struct {
	src   []byte
	dst   []byte
	Error error
}

var (
	// Encode returns a new Encoder instance
	// 返回 encoder 实例
	Encode = newEncoder()
	// Decode returns a new Decoder instance
	// 返回 decoder 实例
	Decode = newDecoder()
	// Encrypt returns a new Encrypter instance
	// 返回 encrypter 实例
	Encrypt = newEncrypter()
	// Decrypt returns a new Decrypter instance
	// 返回 decrypter 实例
	Decrypt = newDecrypter()
	// Sign returns a new Signer instance
	// 返回 signer 实例
	Sign = newSigner()
	// Verify returns a new Verifier instance
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
