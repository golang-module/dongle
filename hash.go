package dongle

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"hash"
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

	SHA3_224   = crypto.SHA3_224
	SHA3_256   = crypto.SHA3_256
	SHA3_384   = crypto.SHA3_384
	SHA3_512   = crypto.SHA3_512
	SHA512_224 = crypto.SHA512_224
	SHA512_256 = crypto.SHA512_256
)

// ByMd4 encrypts by md4.
// 通过 md4 加密
func (e encrypter) ByMd4() encrypter {
	if len(e.src) == 0 {
		return e
	}
	e.dst = hashEncrypt(e.src, MD4)
	return e
}

// ByMd5 encrypts by md5.
// 通过 md5 加密
func (e encrypter) ByMd5() encrypter {
	if len(e.src) == 0 {
		return e
	}
	e.dst = hashEncrypt(e.src, MD5)
	return e
}

// BySha1 encrypts by sha1.
// 通过 sha1 加密
func (e encrypter) BySha1() encrypter {
	if len(e.src) == 0 {
		return e
	}
	e.dst = hashEncrypt(e.src, SHA1)
	return e
}

// BySha3 encrypts by sha3.
// 通过 BySha3 加密
func (e encrypter) BySha3(digests int) encrypter {
	if len(e.src) == 0 {
		return e
	}
	switch digests {
	case 224:
		e.dst = hashEncrypt(e.src, SHA3_224)
	case 256:
		e.dst = hashEncrypt(e.src, SHA3_256)
	case 384:
		e.dst = hashEncrypt(e.src, SHA3_384)
	case 512:
		e.dst = hashEncrypt(e.src, SHA3_512)
	default:
		e.Error = invalidHashDigestsError()
	}
	return e
}

// BySha224 encrypts by sha224.
// 通过 sha224 加密
func (e encrypter) BySha224() encrypter {
	if len(e.src) == 0 {
		return e
	}
	e.dst = hashEncrypt(e.src, SHA224)
	return e
}

// BySha256 encrypts by sha256.
// 通过 sha256 加密
func (e encrypter) BySha256() encrypter {
	if len(e.src) == 0 {
		return e
	}
	e.dst = hashEncrypt(e.src, SHA256)
	return e
}

// BySha384 encrypts by sha384.
// 通过 sha384 加密
func (e encrypter) BySha384() encrypter {
	if len(e.src) == 0 {
		return e
	}
	e.dst = hashEncrypt(e.src, SHA384)
	return e
}

// BySha512 encrypts by sha512.
// 通过 sha512 加密
func (e encrypter) BySha512(digests ...int) encrypter {
	if len(e.src) == 0 {
		return e
	}
	if len(digests) == 0 {
		e.dst = hashEncrypt(e.src, SHA512)
		return e
	}
	switch digests[0] {
	case 224:
		e.dst = hashEncrypt(e.src, SHA512_224)
	case 256:
		e.dst = hashEncrypt(e.src, SHA512_256)
	default:
		e.Error = invalidHashDigestsError()
	}
	return e
}

// ByRipemd160 encrypts by ripemd160.
// 通过 ripemd160 加密
func (e encrypter) ByRipemd160() encrypter {
	if len(e.src) == 0 {
		return e
	}
	e.dst = hashEncrypt(e.src, RIPEMD160)
	return e
}

// hashEncrypt hash encrypt src.
// 对 src 进行哈希加密
func hashEncrypt(src []byte, algo crypto.Hash) (dst []byte) {
	var mac hash.Hash
	switch algo {
	case MD4:
		mac = md4.New()
		mac.Write(src)
		dst = mac.Sum(nil)
	case MD5:
		mac = md5.New()
		mac.Write(src)
		dst = mac.Sum(nil)
	case RIPEMD160:
		mac = ripemd160.New()
		mac.Write(src)
		dst = mac.Sum(nil)
	case SHA1:
		bytes := sha1.Sum(src)
		dst = bytes[:]
	case SHA3_224:
		bytes := sha3.Sum224(src)
		dst = bytes[:]
	case SHA3_256:
		bytes := sha3.Sum256(src)
		dst = bytes[:]
	case SHA3_384:
		bytes := sha3.Sum384(src)
		dst = bytes[:]
	case SHA3_512:
		bytes := sha3.Sum512(src)
		dst = bytes[:]
	case SHA224:
		bytes := sha256.Sum224(src)
		dst = bytes[:]
	case SHA256:
		bytes := sha256.Sum256(src)
		dst = bytes[:]
	case SHA384:
		bytes := sha512.Sum384(src)
		dst = bytes[:]
	case SHA512:
		bytes := sha512.Sum512(src)
		dst = bytes[:]
	case SHA512_224:
		mac = sha512.New512_224()
		mac.Write(src)
		dst = mac.Sum(nil)
	case SHA512_256:
		mac = sha512.New512_256()
		mac.Write(src)
		dst = mac.Sum(nil)
	}
	return dst
}
