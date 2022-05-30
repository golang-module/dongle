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
const Version = "0.1.1"

// dongle defines a dongle struct.
// 定义 dongle 结构体
type dongle struct {
	src   []byte
	dst   []byte
	Error error
}

var (
	// Encode returns a new encode instance
	Encode = newEncode()
	// Decode returns a new decode instance
	Decode = newDecode()
	// Encrypt returns a new encrypt instance
	Encrypt = newEncrypt()
	// Decrypt returns a new decrypt instance
	Decrypt = newDecrypt()
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
