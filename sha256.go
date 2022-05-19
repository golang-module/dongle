package dongle

import (
	"crypto/sha256"
)

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
