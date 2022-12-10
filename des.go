package dongle

import (
	"crypto/des"
)

// ByDes encrypts by des.
// 通过 des 加密
func (e Encrypter) ByDes(c *Cipher) Encrypter {
	block, err := des.NewCipher(c.key)
	if err != nil {
		e.Error = invalidDesKeyError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalidDesIVError()
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = invalidDesSrcError()
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// ByDes decrypts by des.
// 通过 des 解密
func (d Decrypter) ByDes(c *Cipher) Decrypter {
	if d.Error != nil {
		return d
	}
	block, err := des.NewCipher(c.key)
	if err != nil {
		d.Error = invalidDesKeyError()
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = invalidDesIVError()
		return d
	}
	if (c.mode == CBC || c.padding == No) && len(d.src)%block.BlockSize() != 0 {
		d.Error = invalidDesSrcError()
		return d
	}
	d.dst, d.Error = d.decrypt(c, block)
	return d
}
