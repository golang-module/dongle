package dongle

import (
	"golang.org/x/crypto/blowfish"
)

// ByBlowfish encrypts by blowfish.
// 通过 blowfish 加密
func (e Encrypter) ByBlowfish(c *Cipher) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	block, err := blowfish.NewCipher(c.key)
	if err != nil {
		e.Error = invalidBlowfishKeyError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = invalidBlowfishIVError()
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = invalidBlowfishSrcError()
		return e
	}
	e.dst, e.Error = c.Encrypt(e.src, block)
	return e
}

// ByBlowfish decrypts by blowfish.
// 通过 blowfish 解密
func (d Decrypter) ByBlowfish(c *Cipher) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	block, err := blowfish.NewCipher(c.key)
	if err != nil {
		d.Error = invalidBlowfishKeyError()
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = invalidBlowfishIVError()
		return d
	}
	if (c.mode == CBC || c.padding == No) && len(d.src)%block.BlockSize() != 0 {
		d.Error = invalidBlowfishSrcError()
		return d
	}
	d.dst, d.Error = c.Decrypt(d.src, block)
	return d
}
