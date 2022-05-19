package dongle

import (
	"crypto/sha1"
)

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
