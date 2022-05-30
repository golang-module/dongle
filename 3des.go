package dongle

import (
	"crypto/des"
	"fmt"
)

// By3Des encrypts by 3des algorithm.
// 通过 3des 加密
func (e encrypt) By3Des(c *Cipher) encrypt {
	block, err := des.NewTripleDESCipher(c.key)
	if err != nil {
		e.Error = invalid3DesKeyError(len(c.key))
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// By3Des decrypts by 3des algorithm.
// 通过 3des 解密
func (d decrypt) By3Des(c *Cipher) decrypt {
	if d.Error != nil {
		return d
	}
	block, err := des.NewTripleDESCipher(c.key)
	if err != nil {
		d.Error = invalid3DesKeyError(len(c.key))
		return d
	}
	d.dst, d.Error = d.decrypt(c, block)
	return d
}

// returns an invalid 3des key error
// 返回无效的 3des 密钥大小错误
func invalid3DesKeyError(size int) error {
	return fmt.Errorf("invalid 3des key size %d, the key must be 24 bytes", size)
}
