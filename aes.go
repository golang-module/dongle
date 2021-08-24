package dongle

import (
	"crypto/aes"
)

// ByAes encrypts by aes algorithm.
func (e encrypt) ByAes(c *Cipher) encrypt {
	if e.Error != nil {
		return e
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		e.Error = err
		return e
	}
	e.output, e.Error = e.encrypt(c, block)
	return e
}

// ByAes decrypts by aes algorithm.
func (d decrypt) ByAes(c *Cipher) decrypt {
	if d.Error != nil {
		return d
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		d.Error = err
		return d
	}
	d.output, d.Error = d.decrypt(c, block)
	return d
}
