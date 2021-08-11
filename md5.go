package dongle

import (
	"crypto/md5"
)

// ByMd5 encrypts by md5.
func (e encrypt) ByMd5() encrypt {
	if e.Error != nil {
		return e
	}
	if len(e.input) == 0 {
		return e
	}
	hash := md5.New()
	hash.Write(e.input)
	e.output = hash.Sum(nil)
	return e
}
