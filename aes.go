package dongle

import (
	"crypto/aes"
)

// ByAes encrypts by aes algorithm.
// 通过 aes 加密
func (e encrypt) ByAes(c *Cipher) encrypt {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		e.Error = invalidKeyError(len(c.key))
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// ByAes encrypts by aes algorithm.
// 通过 aes 解密
func (d decrypt) ByAes(c *Cipher) decrypt {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		d.Error = invalidKeyError(len(c.key))
		return d
	}
	d.dst, d.Error = d.decrypt(c, block)
	return d
}
