package dongle

import (
	"crypto/des"
)

// ByDes encrypts by des.
// 通过 des 加密
func (e encrypt) ByDes(c *Cipher) encrypt {
	block, err := des.NewCipher(c.key)
	if err != nil {
		e.Error = invalidDesKeyError(len(c.key))
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalidDesIVError(len(c.iv))
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// ByDes decrypts by des.
// 通过 des 解密
func (d decrypt) ByDes(c *Cipher) decrypt {
	block, err := des.NewCipher(c.key)
	if err != nil {
		d.Error = invalidDesKeyError(len(c.key))
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = invalidDesIVError(len(c.iv))
		return d
	}
	d.dst, d.Error = d.decrypt(c, block)
	return d
}
