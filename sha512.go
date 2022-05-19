package dongle

import (
	"crypto/sha512"
)

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
