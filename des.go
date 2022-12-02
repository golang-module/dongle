package dongle

import (
	"crypto/des"
)

// ByDes encrypts by des.
// 通过 des 加密
func (e encrypt) ByDes(c *Cipher) encrypt {
	block, err := des.NewCipher(c.key)
	if err != nil {
		e.Error = invalidDesKeyError()
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = invalidDesSrcError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalidDesIVError()
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// ByDes decrypts by des.
// 通过 des 解密
func (d decrypt) ByDes(c *Cipher) decrypt {
	if d.Error != nil {
		return d
	}
	block, err := des.NewCipher(c.key)
	if err != nil {
		d.Error = invalidDesKeyError()
		return d
	}
	if c.padding == No && len(d.src)%block.BlockSize() != 0 {
		d.Error = invalidDesSrcError()
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = invalidDesIVError()
		return d
	}
	d.dst, d.Error = d.decrypt(c, block)
	return d
}
