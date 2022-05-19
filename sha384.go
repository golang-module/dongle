package dongle

import (
	"crypto/sha512"
)

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
