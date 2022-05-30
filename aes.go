package dongle

import (
	"crypto/aes"
	"fmt"
)

// ByAes encrypts by aes algorithm.
// 通过 aes 加密
func (e encrypt) ByAes(c *Cipher) encrypt {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		e.Error = invalidAesKeyError(len(c.key))
		return e
	}
	e.dst, e.Error = e.encrypt(c, block)
	return e
}

// ByAes decrypts by aes algorithm.
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
	d.dst, d.Error = d.decrypt(c, block)
	return d
}

// returns an invalid aes key error
// 返回无效的 aes 密钥大小错误
func invalidAesKeyError(size int) error {
	return fmt.Errorf("invalid aes key size %d, the key must be 16, 24 or 32 bytes", size)
}
