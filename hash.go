package dongle

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

// ByMd4 encrypts by md4.
// 通过 md4 加密
func (e encrypt) ByMd4() encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := md4.New()
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// ByMd5 encrypts by md5.
// 通过 md5 加密
func (e encrypt) ByMd5() encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := md5.New()
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}

// BySha1 encrypts by sha1.
// 通过 sha1 加密
func (e encrypt) BySha1() encrypt {
	if len(e.src) == 0 {
		return e
	}
	bytes := sha1.Sum(e.src)
	e.dst = bytes[:]
	return e
}

// BySha224 encrypts by sha224.
// 通过 sha224 加密
func (e encrypt) BySha224() encrypt {
	if len(e.src) == 0 {
		return e
	}
	bytes := sha256.Sum224(e.src)
	e.dst = bytes[:]
	return e
}

// BySha256 encrypts by sha256.
// 通过 sha256 加密
func (e encrypt) BySha256() encrypt {
	if len(e.src) == 0 {
		return e
	}
	bytes := sha256.Sum256(e.src)
	e.dst = bytes[:]
	return e
}

// BySha384 encrypts by sha384.
// 通过 sha384 加密
func (e encrypt) BySha384() encrypt {
	if len(e.src) == 0 {
		return e
	}
	bytes := sha512.Sum384(e.src)
	e.dst = bytes[:]
	return e
}

// BySha512 encrypts by sha512.
// 通过 sha512 加密
func (e encrypt) BySha512() encrypt {
	if len(e.src) == 0 {
		return e
	}
	bytes := sha512.Sum512(e.src)
	e.dst = bytes[:]
	return e
}

// ByRipemd160 encrypts by ripemd160.
// 通过 ripemd160 加密
func (e encrypt) ByRipemd160() encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := ripemd160.New()
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}
