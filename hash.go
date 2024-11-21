package dongle

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/dromara/dongle/md2"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type HashError struct {
}

func NewHashError() HashError {
	return HashError{}
}

func (e HashError) SizeError() error {
	return fmt.Errorf("hash: invalid size, the size is unsupported")
}

// ByMd2 encrypts by md2.
func (e Encrypter) ByMd2() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := md2.New()
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByMd4 encrypts by md4.
func (e Encrypter) ByMd4() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := md4.New()
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByMd5 encrypts by md5.
func (e Encrypter) ByMd5() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	dst := md5.Sum(e.src)
	e.dst = dst[:]
	return e
}

// BySha1 encrypts by sha1.
func (e Encrypter) BySha1() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	dst := sha1.Sum(e.src)
	e.dst = dst[:]
	return e
}

// BySha3 encrypts by sha3.
func (e Encrypter) BySha3(size int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	switch size {
	case 224:
		dst := sha3.Sum224(e.src)
		e.dst = dst[:]
	case 256:
		dst := sha3.Sum256(e.src)
		e.dst = dst[:]
	case 384:
		dst := sha3.Sum384(e.src)
		e.dst = dst[:]
	case 512:
		dst := sha3.Sum512(e.src)
		e.dst = dst[:]
	default:
		hashError := HashError{}
		e.Error = hashError.SizeError()
	}
	return e
}

// BySha224 encrypts by sha224.
func (e Encrypter) BySha224() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	dst := sha256.Sum224(e.src)
	e.dst = dst[:]
	return e
}

// BySha256 encrypts by sha256.
func (e Encrypter) BySha256() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	dst := sha256.Sum256(e.src)
	e.dst = dst[:]
	return e
}

// BySha384 encrypts by sha384.
func (e Encrypter) BySha384() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	dst := sha512.Sum384(e.src)
	e.dst = dst[:]
	return e
}

// BySha512 encrypts by sha512.
func (e Encrypter) BySha512(size ...int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	if len(size) == 0 {
		dst := sha512.Sum512(e.src)
		e.dst = dst[:]
		return e
	}
	switch size[0] {
	case 224:
		dst := sha512.Sum512_224(e.src)
		e.dst = dst[:]
	case 256:
		dst := sha512.Sum512_256(e.src)
		e.dst = dst[:]
	default:
		hashError := HashError{}
		e.Error = hashError.SizeError()
	}
	return e
}

// ByShake128 encrypts by shake128.
func (e Encrypter) ByShake128(size int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := make([]byte, size/8)
	sha3.ShakeSum128(h, e.src)
	e.dst = h
	return e
}

// ByShake256 encrypts by shake256.
func (e Encrypter) ByShake256(size int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := make([]byte, size/8)
	sha3.ShakeSum256(h, e.src)
	e.dst = h
	return e
}

// ByRipemd160 encrypts by ripemd160.
func (e Encrypter) ByRipemd160() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := ripemd160.New()
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}

// ByBlake2b encrypts by blake2b.
func (e Encrypter) ByBlake2b(size int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	switch size {
	case 256:
		dst := blake2b.Sum256(e.src)
		e.dst = dst[:]
	case 384:
		dst := blake2b.Sum384(e.src)
		e.dst = dst[:]
	case 512:
		dst := blake2b.Sum512(e.src)
		e.dst = dst[:]
	default:
		hashError := HashError{}
		e.Error = hashError.SizeError()
	}
	return e
}

// ByBlake2s encrypts by blake2s.
func (e Encrypter) ByBlake2s(size int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	switch size {
	case 256:
		dst := blake2s.Sum256(e.src)
		e.dst = dst[:]
	default:
		hashError := HashError{}
		e.Error = hashError.SizeError()
	}
	return e
}
