package dongle

import (
	"crypto/sha256"
)

// BySha224 encrypts by sha224.
func (e encrypt) BySha224() encrypt {
	if e.Error != nil {
		return e
	}
	if len(e.input) == 0 {
		return e
	}
	bytes := sha256.Sum224(e.input)
	e.output = bytes[:]
	return e
}
