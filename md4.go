package dongle

import (
	"golang.org/x/crypto/md4"
)

// ByMd4 encrypts by md4.
// 通过 md4 加密
func (e encrypt) ByMd4() encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := md4.New()
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}
