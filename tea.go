package dongle

import (
	"golang.org/x/crypto/tea"
)

// ByTea encrypts by tea.
// 通过 tea 加密
func (e encrypter) ByTea(key interface{}, rounds ...int) encrypter {
	if len(e.src) == 0 {
		return e
	}
	if len(e.src) != 8 {
		e.Error = invalidTeaSrcError()
		return e
	}
	if len(rounds) == 0 {
		// 64 is the standard number of rounds in tea.
		rounds = []int{64}
	}
	if rounds[0]&1 != 0 {
		e.Error = invalidTeaRoundsError()
		return e
	}

	block, err := tea.NewCipherWithRounds(interface2bytes(key), rounds[0])
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
func (d decrypter) ByTea(key interface{}, rounds ...int) decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	if len(d.src) != 8 {
		d.Error = invalidTeaSrcError()
		return d
	}
	if len(rounds) == 0 {
		// 64 is the standard number of rounds in tea.
		rounds = []int{64}
	}
	if rounds[0]&1 != 0 {
		d.Error = invalidTeaRoundsError()
		return d
	}
	block, err := tea.NewCipherWithRounds(interface2bytes(key), rounds[0])
	if err != nil {
		d.Error = invalidTeaKeyError()
		return d
	}
	d.dst = make([]byte, 8)
	block.Decrypt(d.dst, d.src)
	return d
}
