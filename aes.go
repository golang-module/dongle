package dongle

import (
	"crypto/aes"
	"fmt"
)

type AesError struct {
}

func NewAesError() AesError {
	return AesError{}
}

func (e AesError) SrcError() error {
	return fmt.Errorf("aes: invalid src, the src is not full blocks")
}

func (e AesError) KeyError() error {
	return fmt.Errorf("aes: invalid key, the key must be 16, 24 or 32 bytes")
}

func (e AesError) IvError() error {
	return fmt.Errorf("aes: invalid iv, the iv size must be 16 bytes")
}

var aesError = NewAesError()

// ByAes encrypts by aes.
func (e Encrypter) ByAes(c *Cipher) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		e.Error = aesError.KeyError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = aesError.IvError()
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = aesError.SrcError()
		return e
	}
	e.dst, e.Error = c.Encrypt(e.src, block)
	return e
}

// ByAes decrypts by aes.
func (d Decrypter) ByAes(c *Cipher) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		d.Error = aesError.KeyError()
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = aesError.IvError()
		return d
	}
	if (c.mode == CBC || c.padding == No) && len(d.src)%block.BlockSize() != 0 {
		d.Error = aesError.SrcError()
		return d
	}
	d.dst, d.Error = c.Decrypt(d.src, block)
	return d
}
