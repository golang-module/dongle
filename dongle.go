// @Package dongle
// @Description a simple, semantic and developer-friendly golang crypto package
// @Page github.com/dromara/dongle
// @Developer gouguoyin
// @Blog www.gouguoyin.com
// @Email 245629560@qq.com

// Package dongle is a simple, semantic and developer-friendly golang crypto package.
package dongle

import (
	"unsafe"
)

// Version current version
const Version = "1.0.2"

// dongle defines a dongle struct.
type dongle struct {
	src   []byte
	dst   []byte
	Error error
}

var (
	// Encode returns a new Encoder instance
	Encode = newEncoder()
	// Decode returns a new Decoder instance
	Decode = newDecoder()
	// Encrypt returns a new Encrypter instance
	Encrypt = newEncrypter()
	// Decrypt returns a new Decrypter instance
	Decrypt = newDecrypter()
	// Sign returns a new Signer instance
	Sign = newSigner()
	// Verify returns a new Verifier instance
	Verify = newVerifier()
)

// converts string to byte slice without a memory allocation.
func string2bytes(s string) []byte {
	if len(s) == 0 {
		return []byte("")
	}
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func bytes2string(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// split the byte slice by the specified size.
func bytesSplit(buf []byte, size int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/size+1)
	for len(buf) >= size {
		chunk, buf = buf[:size], buf[size:]
		chunks = append(chunks, chunk)
	}
	return chunks
}
