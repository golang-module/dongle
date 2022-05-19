package dongle

import (
	"crypto/sha256"
)

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
