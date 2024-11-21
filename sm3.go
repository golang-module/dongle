package dongle

import (
	"github.com/emmansun/gmsm/sm3"
)

// BySm3 encrypts by sm3.
func (e Encrypter) BySm3() Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	h := sm3.New()
	h.Write(e.src)
	e.dst = h.Sum(nil)
	return e
}
