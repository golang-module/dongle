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
	"golang.org/x/crypto/sha3"
	"hash"
)

// ByHmacMd4 encrypts by hmac with md4.
// 通过 hmac-md4 加密
func (e encrypter) ByHmacMd4(key interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	mac := hmac.New(md4.New, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacMd5 encrypts by hmac with md5.
// 通过 hmac-md5 加密
func (e encrypter) ByHmacMd5(key interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	mac := hmac.New(md5.New, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacSha1 encrypts by hmac with sha1.
// 通过 hmac-sha1 加密
func (e encrypter) ByHmacSha1(key interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	mac := hmac.New(sha1.New, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacSha3 encrypts by hmac with sha3.
// 通过 hmac-sha3 加密
func (e encrypter) ByHmacSha3(key interface{}, digests int) encrypter {
	if len(e.src) == 0 {
		return e
	}
	var hashFun func() hash.Hash
	switch digests {
	case 224:
		hashFun = sha3.New224
	case 256:
		hashFun = sha3.New256
	case 384:
		hashFun = sha3.New384
	case 512:
		hashFun = sha3.New512
	default:
		e.Error = invalidHashDigestsError()
		return e
	}
	mac := hmac.New(hashFun, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacSha224 encrypts by hmac with sha224.
// 通过 hmac-sha224 加密
func (e encrypter) ByHmacSha224(key interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	mac := hmac.New(sha256.New224, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacSha256 encrypts by hmac with sha256.
// 通过 hmac-sha256 加密
func (e encrypter) ByHmacSha256(key interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	mac := hmac.New(sha256.New, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacSha384 encrypts by hmac with sha384.
// 通过 hmac-sha384 加密
func (e encrypter) ByHmacSha384(key interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	mac := hmac.New(sha512.New384, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacSha512 encrypts by hmac with sha512.
// 通过 hmac-sha512 加密
func (e encrypter) ByHmacSha512(key interface{}, digests ...int) encrypter {
	if len(e.src) == 0 {
		return e
	}
	var hashFun func() hash.Hash
	if len(digests) == 0 {
		digests = []int{512}
	}
	switch digests[0] {
	case 512:
		hashFun = sha512.New
	case 224:
		hashFun = sha512.New512_224
	case 256:
		hashFun = sha512.New512_256
	default:
		e.Error = invalidHashDigestsError()
		return e
	}
	mac := hmac.New(hashFun, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacRipemd160 encrypts by hmac with ripemd160.
// 通过 hmac-ripemd160 加密
func (e encrypter) ByHmacRipemd160(key interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	mac := hmac.New(ripemd160.New, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}

// ByHmacSm3 encrypts by hmac with sm3.
// 通过 hmac-sm3 加密
func (e encrypter) ByHmacSm3(key interface{}) encrypter {
	if len(e.src) == 0 {
		return e
	}
	mac := hmac.New(sm3.New, interface2bytes(key))
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}
