package dongle

import (
	"crypto/sha512"
)

// BySha512 encrypts by sha512.
func (e encrypt) BySha512() encrypt {
	if e.Error != nil {
		return e
	}
	if len(e.input) == 0 {
		return e
	}
	bytes := sha512.Sum512(e.input)
	e.output = bytes[:]
	return e
}
