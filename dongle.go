// @Package dongle
// @Description a simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption
// @Page github.com/golang-module/dongle
// @Version v0.0.2
// @Author gouguoyin
// @Blog www.gouguoyin.cn
// @Email contact@gouguoyin.cn

// Package dongle is a simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption.
package dongle

import (
	"unsafe"
)

// dongle defines a dongle struct.
type dongle struct {
	input  []byte
	output []byte
	Error  error
}

var (
	// emptyString defines a empty string
	emptyString = ""
	// emptyBytes defines a empty byte slice
	emptyBytes = []byte{}

	// Encode returns a new encode instance
	Encode = newEncode()
	// Decode returns a new decode instance
	Decode = newDecode()
	// Encrypt returns a new encrypt instance
	Encrypt = newEncrypt()
	// Decrypt returns a new decrypt instance
	Decrypt = newDecrypt()
)

// string2bytes converts string to byte slice without a memory allocation.
func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// bytes2string converts byte slice to string without a memory allocation.
func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
