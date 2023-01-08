package dongle

import (
	"crypto/des"
)

// By3Des encrypts by 3des.
// 通过 3des 加密
func (e Encrypter) By3Des(c *Cipher) Encrypter {
	block, err := des.NewTripleDESCipher(c.key)
	if err != nil {
		e.Error = invalid3DesKeyError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalid3DesIVError()
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = invalid3DesSrcError()
		return e
	}
	e.dst, e.Error = c.encrypt(e.src, block)
	return e
}

// By3Des decrypts by 3des.
// 通过 3des 解密
func (d Decrypter) By3Des(c *Cipher) Decrypter {
	if d.Error != nil {
		return d
	}
	block, err := des.NewTripleDESCipher(c.key)
	if err != nil {
		d.Error = invalid3DesKeyError()
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = invalid3DesIVError()
		return d
	}
	if (c.mode == CBC || c.padding == No) && len(d.src)%block.BlockSize() != 0 {
		d.Error = invalid3DesSrcError()
		return d
	}
	d.dst, d.Error = c.decrypt(d.src, block)
	return d
}
