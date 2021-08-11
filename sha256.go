package dongle

import (
	"crypto/sha256"
)

// BySha256 encrypts by sha256.
func (e encrypt) BySha256() encrypt {
	if e.Error != nil {
		return e
	}
	if len(e.input) == 0 {
		return e
	}
	bytes := sha256.Sum256(e.input)
	e.output = bytes[:]
	return e
}
