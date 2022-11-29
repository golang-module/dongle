package dongle

import (
	"crypto/des"
	"fmt"
)

var (
	invalid3DesKeyError = func() error {
		return fmt.Errorf("3des: invalids key, the key must be 24 bytes")
	}
	invalid3DesIVError = func() error {
		return fmt.Errorf("3des: invalid iv, the iv size must be 8 bytes")
	}
)

// By3Des encrypts by 3des.
// 通过 3des 加密
func (e encrypt) By3Des(c *Cipher) encrypt {
	block, err := des.NewTripleDESCipher(c.key)
	if err != nil {
		e.Error = invalid3DesKeyError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalid3DesIVError()
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// By3Des decrypts by 3des.
// 通过 3des 解密
func (d decrypt) By3Des(c *Cipher) decrypt {
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
	d.dst, d.Error = d.decrypt(c, block)
	return d
}
