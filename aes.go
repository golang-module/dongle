package dongle

import (
	"crypto/aes"
)

// ByAes encrypts by aes.
// 通过 aes 加密
func (e Encrypter) ByAes(c *Cipher) Encrypter {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		e.Error = invalidAesKeyError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalidAesIVError()
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = invalidAesSrcError()
		return e
	}
	e.dst, e.Error = c.encrypt(e.src, block)
	return e
}

// ByAes decrypts by aes.
// 通过 aes 解密
func (d Decrypter) ByAes(c *Cipher) Decrypter {
	if d.Error != nil {
		return d
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		d.Error = invalidAesKeyError()
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = invalidAesIVError()
		return d
	}
	if (c.mode == CBC || c.padding == No) && len(d.src)%block.BlockSize() != 0 {
		d.Error = invalidAesSrcError()
		return d
	}
	d.dst, d.Error = c.decrypt(d.src, block)
	return d
}
