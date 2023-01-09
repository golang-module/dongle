package dongle

import (
	"bytes"
	"golang.org/x/crypto/tea"
)

// ByTea encrypts by tea.
// 通过 tea 加密
func (e Encrypter) ByTea(key interface{}, rounds ...int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
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
	src, size := e.src, tea.BlockSize
	if len(src) > size {
		src = NewCipher().EmptyPadding(e.src, size)
	}
	dst, buffer := make([]byte, size), bytes.NewBufferString("")
	for _, chunk := range bytesSplit(src, size) {
		block, err := tea.NewCipherWithRounds(interface2bytes(key), rounds[0])
		if err != nil {
			e.Error = invalidTeaKeyError()
			return e
		}
		block.Encrypt(dst, chunk)
		buffer.Write(dst)
	}
	e.dst = buffer.Bytes()
	return e
}

// ByTea decrypts by tea.
// 通过 tea 解密
func (d Decrypter) ByTea(key interface{}, rounds ...int) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
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
	src, size := d.src, tea.BlockSize
	dst, buffer := make([]byte, size), bytes.NewBufferString("")
	for _, chunk := range bytesSplit(src, size) {
		block, err := tea.NewCipherWithRounds(interface2bytes(key), rounds[0])
		if err != nil {
			d.Error = invalidTeaKeyError()
			return d
		}
		block.Decrypt(dst, chunk)
		buffer.Write(dst)
	}
	d.dst = NewCipher().EmptyUnPadding(buffer.Bytes())
	return d
}
