package dongle

import (
	"fmt"

	"golang.org/x/crypto/blowfish"
)

type BlowfishError struct {
}

func NewBlowfishError() BlowfishError {
	return BlowfishError{}
}

func (e BlowfishError) SrcError() error {
	return fmt.Errorf("blowfish: invalid src, the src is not full blocks")
}

func (e BlowfishError) KeyError(key []byte) error {
	return fmt.Errorf("blowfish: invalid key size %d, the key must from 1 to 56 bytes", len(key))
}

func (e BlowfishError) IvError(iv []byte) error {
	return fmt.Errorf("blowfish: invalid iv size %d, the iv size must be 8 bytes", len(iv))
}

var blowfishError = NewBlowfishError()

// ByBlowfish encrypts by blowfish.
func (e Encrypter) ByBlowfish(c *Cipher) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	block, err := blowfish.NewCipher(c.key)
	if err != nil {
		e.Error = blowfishError.KeyError(c.key)
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = blowfishError.IvError(c.iv)
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = blowfishError.SrcError()
		return e
	}
	e.dst, e.Error = c.Encrypt(e.src, block)
	return e
}

// ByBlowfish decrypts by blowfish.
func (d Decrypter) ByBlowfish(c *Cipher) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	block, err := blowfish.NewCipher(c.key)
	if err != nil {
		d.Error = blowfishError.KeyError(c.key)
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = blowfishError.IvError(c.iv)
		return d
	}
	if (c.mode == CBC || c.padding == No) && len(d.src)%block.BlockSize() != 0 {
		d.Error = blowfishError.SrcError()
		return d
	}
	d.dst, d.Error = c.Decrypt(d.src, block)
	return d
}
