package dongle

import (
	"crypto/rc4"
	"fmt"
)

var (
	invalidRc4KeyError = func(size int) error {
		return fmt.Errorf("invalid rc4 key size %d, the key at least 1 byte and at most 256 bytes", size)
	}
)

// ByRc4 encrypts by rc4.
// 通过 rc4 加密
func (e encrypt) ByRc4(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var cipher *rc4.Cipher
	size := 0
	switch v := key.(type) {
	case string:
		size = len(v)
		cipher, e.Error = rc4.NewCipher(string2bytes(v))
	case []byte:
		size = len(v)
		cipher, e.Error = rc4.NewCipher(v)
	}
	if e.Error != nil {
		e.Error = invalidRc4KeyError(size)
		return e
	}
	dst := make([]byte, len(e.src))
	cipher.XORKeyStream(dst, e.src)
	e.dst = dst
	return e
}
