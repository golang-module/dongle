package dongle

import (
	"github.com/tjfoc/gmsm/sm3"
)

// BySm3 encrypts by sm3.
// 通过 sm3 加密
func (e encrypter) BySm3() encrypter {
	if len(e.src) == 0 {
		return e
	}
	hasher := sm3.New()
	hasher.Write(e.src)
	e.dst = hasher.Sum(nil)
	return e
}
