package dongle

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"

	"github.com/dromara/dongle/md2"
	"github.com/emmansun/gmsm/sm3"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type HmacError struct {
}

func NewHmacError() HmacError {
	return HmacError{}
}

func (e HmacError) SizeError() error {
	return fmt.Errorf("hmac: invalid size, the size is unsupported")
}

// ByHmacMd2 encrypts by hmac with md2.
func (e Encrypter) ByHmacMd2(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(md2.New, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacMd4 encrypts by hmac with md4.
func (e Encrypter) ByHmacMd4(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(md4.New, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacMd5 encrypts by hmac with md5.
func (e Encrypter) ByHmacMd5(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(md5.New, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha1 encrypts by hmac with sha1.
func (e Encrypter) ByHmacSha1(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sha1.New, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha3 encrypts by hmac with sha3.
func (e Encrypter) ByHmacSha3(key []byte, size int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	var f func() hash.Hash
	switch size {
	case 224:
		f = sha3.New224
	case 256:
		f = sha3.New256
	case 384:
		f = sha3.New384
	case 512:
		f = sha3.New512
	default:
		hmacError := HmacError{}
		e.Error = hmacError.SizeError()
		return e
	}
	h := hmac.New(f, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha224 encrypts by hmac with sha224.
func (e Encrypter) ByHmacSha224(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sha256.New224, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha256 encrypts by hmac with sha256.
func (e Encrypter) ByHmacSha256(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sha256.New, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha384 encrypts by hmac with sha384.
func (e Encrypter) ByHmacSha384(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sha512.New384, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha512 encrypts by hmac with sha512.
func (e Encrypter) ByHmacSha512(key []byte, size ...int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	var f func() hash.Hash
	if len(size) == 0 {
		size = []int{512}
	}
	switch size[0] {
	case 512:
		f = sha512.New
	case 224:
		f = sha512.New512_224
	case 256:
		f = sha512.New512_256
	default:
		hmacError := HmacError{}
		e.Error = hmacError.SizeError()
		return e
	}
	h := hmac.New(f, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacRipemd160 encrypts by hmac with ripemd160.
func (e Encrypter) ByHmacRipemd160(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(ripemd160.New, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSm3 encrypts by hmac with sm3.
func (e Encrypter) ByHmacSm3(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sm3.New, key)
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}
