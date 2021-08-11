package dongle

import (
	"crypto/sha512"
)

// BySha384 encrypts by sha384.
func (e encrypt) BySha384() encrypt {
	if e.Error != nil {
		return e
	}
	if len(e.input) == 0 {
		return e
	}

	bytes := sha512.Sum384(e.input)
	e.output = bytes[:]
	return e
}
