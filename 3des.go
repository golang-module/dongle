package dongle

import "C"
import (
	"crypto/des"
)

// By3Des encrypts by 3des.
// 通过 3des 加密
func (e encrypt) By3Des(c *Cipher) encrypt {
	block, err := des.NewTripleDESCipher(c.key)
	if err != nil {
		e.Error = invalid3DesKeyError(len(c.key))
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalid3DesIVError(len(c.iv))
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// By3Des decrypts by 3des.
// 通过 3des 解密
func (d decrypt) By3Des(c *Cipher) decrypt {
	block, err := des.NewTripleDESCipher(c.key)
	if err != nil {
		d.Error = invalid3DesKeyError(len(c.key))
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = invalid3DesIVError(len(c.iv))
		return d
	}
	d.dst, d.Error = d.decrypt(c, block)
	return d
}
