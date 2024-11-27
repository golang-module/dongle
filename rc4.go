package dongle

import (
	"crypto/rc4"
	"fmt"
)

type Rc4Error struct {
}

func NewRc4Error() Rc4Error {
	return Rc4Error{}
}

func (e Rc4Error) KeyError() error {
	return fmt.Errorf("rc4: invalid key, the key at least 1 byte and at most 256 bytes")
}

var rc4Error = NewRc4Error()

// ByRc4 encrypts by rc4.
func (e Encrypter) ByRc4(key []byte) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		e.Error = rc4Error.KeyError()
		return e
	}
	dst := make([]byte, len(e.src))
	cipher.XORKeyStream(dst, e.src)
	e.dst = dst
	return e
}

// ByRc4 decrypts by rc4.
func (d Decrypter) ByRc4(key []byte) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		d.Error = rc4Error.KeyError()
		return d
	}
	dst := make([]byte, len(d.src))
	cipher.XORKeyStream(dst, d.src)
	d.dst = dst
	return d
}
