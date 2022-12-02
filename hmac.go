package dongle

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/tjfoc/gmsm/sm3"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

// ByHmacMd4 encrypts by hmac with md4.
// 通过 hmac-md4 加密
func (e encrypt) ByHmacMd4(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(md4.New, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByHmacMd5 encrypts by hmac with md5.
// 通过 hmac-md5 加密
func (e encrypt) ByHmacMd5(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(md5.New, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByHmacSha1 encrypts by hmac with sha1.
// 通过 hmac-sha1 加密
func (e encrypt) ByHmacSha1(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(sha1.New, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByHmacSha224 encrypts by hmac with sha224.
// 通过 hmac-sha224 加密
func (e encrypt) ByHmacSha224(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(sha256.New224, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByHmacSha256 encrypts by hmac with sha256.
// 通过 hmac-sha256 加密
func (e encrypt) ByHmacSha256(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(sha256.New, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByHmacSha384 encrypts by hmac with sha384.
// 通过 hmac-sha384 加密
func (e encrypt) ByHmacSha384(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(sha512.New384, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByHmacSha512 encrypts by hmac with sha512.
// 通过 hmac-sha512 加密
func (e encrypt) ByHmacSha512(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(sha512.New, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByHmacRipemd160 encrypts by hmac with ripemd160.
// 通过 hmac-ripemd160 加密
func (e encrypt) ByHmacRipemd160(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(ripemd160.New, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByHmacSm3 encrypts by hmac with sm3.
// 通过 hmac-sm3 加密
func (e encrypt) ByHmacSm3(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := hmac.New(sm3.New, interface2bytes(key))
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}
