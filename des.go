package dongle

import (
	"crypto/des"
	"fmt"
)

// ByDes encrypts by des algorithm.
// 通过 des 加密
func (e encrypt) ByDes(c *Cipher) encrypt {
	block, err := des.NewCipher(c.key)
	if err != nil {
		e.Error = invalidDesKeyError(len(c.key))
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// ByDes decrypts by des algorithm.
// 通过 des 解密
func (d decrypt) ByDes(c *Cipher) decrypt {
	if d.Error != nil {
		return d
	}
	block, err := des.NewCipher(c.key)
	if err != nil {
		d.Error = invalidDesKeyError(len(c.key))
		return d
	}
	d.dst, d.Error = d.decrypt(c, block)
	return d
}

// returns an invalid des key error
// 返回无效的 des 密钥大小错误
func invalidDesKeyError(size int) error {
	return fmt.Errorf("invalid des key size %d, the key must be 8 bytes", size)
}
