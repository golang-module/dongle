package dongle

import (
	"crypto/aes"
)

// ByAes encrypts by aes.
// 通过 aes 加密
func (e encrypt) ByAes(c *Cipher) encrypt {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		e.Error = invalidAesKeyError(len(c.key))
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalidAesIVError(len(c.iv))
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// ByAes decrypts by aes.
// 通过 aes 解密
func (d decrypt) ByAes(c *Cipher) decrypt {
	if d.Error != nil {
		return d
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		d.Error = invalidAesKeyError(len(c.key))
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = invalidAesIVError(len(c.iv))
		return d
	}
	d.dst, d.Error = d.decrypt(c, block)
	return d
}
