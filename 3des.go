package dongle

import (
	"crypto/des"
	"fmt"
)

type TripleDesError struct {
}

func NewTripleDesError() TripleDesError {
	return TripleDesError{}
}

func (e TripleDesError) SrcError() error {
	return fmt.Errorf("3des: invalid src, the src is not full blocks")
}

func (e TripleDesError) KeyError() error {
	return fmt.Errorf("3des: invalids key, the key must be 24 bytes")
}

func (e TripleDesError) IvError() error {
	return fmt.Errorf("3des: invalid iv, the iv size must be 8 bytes")
}

// By3Des encrypts by 3des.
func (e Encrypter) By3Des(c *Cipher) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	block, err := des.NewTripleDESCipher(c.key)
	tripleDesError := NewTripleDesError()
	if err != nil {
		e.Error = tripleDesError.KeyError()
		return e
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		e.Error = tripleDesError.IvError()
		return e
	}
	if c.padding == No && len(e.src)%block.BlockSize() != 0 {
		e.Error = tripleDesError.SrcError()
		return e
	}
	e.dst, e.Error = c.Encrypt(e.src, block)
	return e
}

// By3Des decrypts by 3des.
func (d Decrypter) By3Des(c *Cipher) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	block, err := des.NewTripleDESCipher(c.key)
	tripleDesError := NewTripleDesError()
	if err != nil {
		d.Error = tripleDesError.KeyError()
		return d
	}
	if c.mode != ECB && len(c.iv) != block.BlockSize() {
		d.Error = tripleDesError.IvError()
		return d
	}
	if (c.mode == CBC || c.padding == No) && len(d.src)%block.BlockSize() != 0 {
		d.Error = tripleDesError.SrcError()
		return d
	}
	d.dst, d.Error = c.Decrypt(d.src, block)
	return d
}
