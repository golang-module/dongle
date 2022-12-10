package dongle

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"gitee.com/golang-package/dongle/md2"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// hash algorithm constants
// 哈希算法常量
const (
	MD4       = crypto.MD4
	MD5       = crypto.MD5
	SHA1      = crypto.SHA1
	SHA224    = crypto.SHA224
	SHA256    = crypto.SHA256
	SHA384    = crypto.SHA384
	SHA512    = crypto.SHA512
	RIPEMD160 = crypto.RIPEMD160
)

// ByMd2 encrypts by md2.
// 通过 md2 加密
func (e Encrypter) ByMd2() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := md2.New()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// ByMd4 encrypts by md4.
// 通过 md4 加密
func (e Encrypter) ByMd4() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := md4.New()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// ByMd5 encrypts by md5.
// 通过 md5 加密
func (e Encrypter) ByMd5() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := md5.New()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// BySha1 encrypts by sha1.
// 通过 sha1 加密
func (e Encrypter) BySha1() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := sha1.New()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// BySha3 encrypts by sha3.
// 通过 BySha3 加密
func (e Encrypter) BySha3(size int) Encrypter {
	var hasher hash.Hash
	if len(e.src) == 0 {
		return e
	}
	switch size {
	case 224:
		hasher = sha3.New224()
	case 256:
		hasher = sha3.New256()
	case 384:
		hasher = sha3.New384()
	case 512:
		hasher = sha3.New512()
	default:
		e.Error = invalidHashSizeError()
		return e
	}
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// BySha224 encrypts by sha224.
// 通过 sha224 加密
func (e Encrypter) BySha224() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := sha256.New224()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// BySha256 encrypts by sha256.
// 通过 sha256 加密
func (e Encrypter) BySha256() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := sha256.New()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// BySha384 encrypts by sha384.
// 通过 sha384 加密
func (e Encrypter) BySha384() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := sha512.New384()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// BySha512 encrypts by sha512.
// 通过 sha512 加密
func (e Encrypter) BySha512(size ...int) Encrypter {
	var hasher hash.Hash
	if len(e.src) == 0 {
		return e
	}
	if len(size) == 0 {
		hasher = sha512.New()
		hasher.Write(e.src)
		e.dst = hasher.Sum(nil)
		return e
	}
	switch size[0] {
	case 224:
		hasher = sha512.New512_224()
	case 256:
		hasher = sha512.New512_256()
	default:
		e.Error = invalidHashSizeError()
		return e
	}
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}

// ByRipemd160 encrypts by ripemd160.
// 通过 ripemd160 加密
func (e Encrypter) ByRipemd160() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := ripemd160.New()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}
