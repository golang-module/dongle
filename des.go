package dongle

import (
	"crypto/des"
	"fmt"
)

type DesError struct {
}

func NewDesError() DesError {
	return DesError{}
}

func (e DesError) SrcError() error {
	return fmt.Errorf("des: invalid src, the src is not full blocks")
}

func (e DesError) KeyError() error {
	return fmt.Errorf("des: invalid key, the key must be 8 bytes")
}

func (e DesError) IvError() error {
	return fmt.Errorf("des: invalid iv, the iv size must be 8 bytes")
}

var desError = NewDesError()

// ByDes encrypts by des.
func (e Encrypter) ByDes(c *Cipher) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	block, err := des.NewCipher(c.key)
	if err != nil {
		e.Error = desError.KeyError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = desError.IvError()
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = desError.SrcError()
		return e
	}
	e.dst, e.Error = c.Encrypt(e.src, block)
	return e
}

// ByDes decrypts by des.
func (d Decrypter) ByDes(c *Cipher) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	block, err := des.NewCipher(c.key)
	if err != nil {
		d.Error = desError.KeyError()
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = desError.IvError()
		return d
	}
	if (c.mode == CBC || c.padding == No) && len(d.src)%block.BlockSize() != 0 {
		d.Error = desError.SrcError()
		return d
	}
	d.dst, d.Error = c.Decrypt(d.src, block)
	return d
}
