package dongle

import (
	"github.com/tjfoc/gmsm/sm3"
)

// BySm3 encrypts by sm3.
// 通过 sm3 加密
func (e Encrypter) BySm3() Encrypter {
	if len(e.src) == 0 {
		return e
	}
	h := sm3.New()
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}
