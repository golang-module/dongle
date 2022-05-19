package dongle

import (
	"crypto/md5"
)

// ByMd5 encrypts by md5.
// 通过 md5 加密
func (e encrypt) ByMd5() encrypt {
	if len(e.src) == 0 {
		return e
	}
	hash := md5.New()
	hash.Write(e.src)
	e.dst = hash.Sum(nil)
	return e
}
