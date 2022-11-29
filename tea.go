package dongle

import (
	"fmt"
	"golang.org/x/crypto/tea"
)

var (
	invalidTeaSrcError = func() error {
		return fmt.Errorf("tea: invalid src, the src must be 8 bytes")
	}
	invalidTeaRoundsError = func() error {
		return fmt.Errorf("tea: invalid rounds, the rounds must be even")
	}
	invalidTeaKeyError = func() error {
		return fmt.Errorf("tea: invalid key, the key must be 16 bytes")
	}
)

// ByTea encrypts by tea.
// 通过 tea 加密
func (e encrypt) ByTea(key interface{}, rounds int) encrypt {
	if len(e.src) == 0 {
		return e
	}
	if len(e.src) != 8 {
		e.Error = invalidTeaSrcError()
		return e
	}
	if rounds&1 != 0 {
		e.Error = invalidTeaRoundsError()
		return e
	}
	var keys []byte
	switch v := key.(type) {
	case string:
		keys = string2bytes(v)
	case []byte:
		keys = v
	}
	block, err := tea.NewCipherWithRounds(keys, rounds)
	if err != nil {
		e.Error = invalidTeaKeyError()
		return e
	}
	e.dst = make([]byte, 8)
	block.Encrypt(e.dst, e.src)
	return e
}

// ByTea decrypts by tea.
// 通过 tea 解密
func (d decrypt) ByTea(key interface{}, rounds int) decrypt {
	if len(d.src) == 0 {
		return d
	}
	if len(d.src) != 8 {
		d.Error = invalidTeaSrcError()
		return d
	}
	if rounds&1 != 0 {
		d.Error = invalidTeaRoundsError()
		return d
	}
	var keys []byte
	switch v := key.(type) {
	case string:
		keys = string2bytes(v)
	case []byte:
		keys = v
	}
	block, err := tea.NewCipherWithRounds(keys, rounds)
	if err != nil {
		d.Error = invalidTeaKeyError()
		return d
	}
	d.dst = make([]byte, 8)
	block.Decrypt(d.dst, d.src)
	return d
}
