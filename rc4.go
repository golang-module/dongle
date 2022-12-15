package dongle

import (
	"crypto/rc4"
)

// ByRc4 encrypts by rc4.
// 通过 rc4 加密
func (e Encrypter) ByRc4(key interface{}) Encrypter {
	if len(e.src) == 0 {
		return e
	}
	cipher, err := rc4.NewCipher(interface2bytes(key))
	if err != nil {
		e.Error = invalidRc4KeyError()
		return e
	}
	dst := make([]byte, len(e.src))
	cipher.XORKeyStream(dst, e.src)
	e.dst = dst
	return e
}

// ByRc4 decrypts by rc4.
// 通过 rc4 解密
func (d Decrypter) ByRc4(key interface{}) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	cipher, err := rc4.NewCipher(interface2bytes(key))
	if err != nil {
		d.Error = invalidRc4KeyError()
		return d
	}
	dst := make([]byte, len(d.src))
	cipher.XORKeyStream(dst, d.src)
	d.dst = dst
	return d
}
