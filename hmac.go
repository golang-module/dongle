package dongle

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"gitee.com/golang-package/dongle/md2"
	"github.com/emmansun/gmsm/sm3"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// ByHmacMd2 encrypts by hmac with md2.
// 通过 hmac-md2 加密
func (e Encrypter) ByHmacMd2(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(md2.New, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacMd4 encrypts by hmac with md4.
// 通过 hmac-md4 加密
func (e Encrypter) ByHmacMd4(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(md4.New, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacMd5 encrypts by hmac with md5.
// 通过 hmac-md5 加密
func (e Encrypter) ByHmacMd5(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(md5.New, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha1 encrypts by hmac with sha1.
// 通过 hmac-sha1 加密
func (e Encrypter) ByHmacSha1(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sha1.New, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha3 encrypts by hmac with sha3.
// 通过 hmac-sha3 加密
func (e Encrypter) ByHmacSha3(key interface{}, size int) Encrypter {
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
		e.Error = invalidHashSizeError()
		return e
	}
	h := hmac.New(f, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha224 encrypts by hmac with sha224.
// 通过 hmac-sha224 加密
func (e Encrypter) ByHmacSha224(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sha256.New224, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha256 encrypts by hmac with sha256.
// 通过 hmac-sha256 加密
func (e Encrypter) ByHmacSha256(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sha256.New, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha384 encrypts by hmac with sha384.
// 通过 hmac-sha384 加密
func (e Encrypter) ByHmacSha384(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sha512.New384, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSha512 encrypts by hmac with sha512.
// 通过 hmac-sha512 加密
func (e Encrypter) ByHmacSha512(key interface{}, size ...int) Encrypter {
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
		e.Error = invalidHashSizeError()
		return e
	}
	h := hmac.New(f, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacRipemd160 encrypts by hmac with ripemd160.
// 通过 hmac-ripemd160 加密
func (e Encrypter) ByHmacRipemd160(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(ripemd160.New, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByHmacSm3 encrypts by hmac with sm3.
// 通过 hmac-sm3 加密
func (e Encrypter) ByHmacSm3(key interface{}) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := hmac.New(sm3.New, interface2bytes(key))
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}
