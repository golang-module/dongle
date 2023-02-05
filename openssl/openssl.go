// Package openssl for key certificate management
package openssl

import (
	"bytes"
	"unsafe"
)

// defines key format enum type.
// 定义密钥格式枚举类型
type keyFormat string

// key format constants.
// 密钥格式常量
const (
	PKCS1 keyFormat = "pkcs1"
	PKCS8 keyFormat = "pkcs8"
)

var (
	RSA = newRsaKeyPair()
)

// converts string to byte slice without a memory allocation.
// 将字符串转换为字节切片
func string2bytes(s string) []byte {
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
	return *(*string)(unsafe.Pointer(&b))
}

// split the string by the specified size.
// 按照指定长度分割字符串
func stringSplit(s string, n int) string {
	substr, strings := "", ""
	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		substr = substr + string(r)
		if (i+1)%n == 0 {
			strings = strings + substr + "\n"
			substr = ""
		} else if (i + 1) == l {
			strings = strings + substr + "\n"
		}
	}
	return strings
}
