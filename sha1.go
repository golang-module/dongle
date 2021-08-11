package dongle

import (
	"crypto/sha1"
)

// BySha1 encrypts by sha1.
func (e encrypt) BySha1() encrypt {
	if e.Error != nil {
		return e
	}
	if len(e.input) == 0 {
		return e
	}
	bytes := sha1.Sum(e.input)
	e.output = bytes[:]
	return e
}
