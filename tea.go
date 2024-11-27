package dongle

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/tea"
)

type TeaError struct {
}

func NewTeaError() TeaError {
	return TeaError{}
}

func (e TeaError) KeyError(key []byte) error {
	return fmt.Errorf("tea: invalid key size %d, the key must be 16 bytes", len(key))
}

func (e TeaError) RoundsError() error {
	return fmt.Errorf("tea: invalid rounds, the rounds must be even")
}

var teaError = NewTeaError()

// ByTea encrypts by tea.
func (e Encrypter) ByTea(key []byte, rounds ...int) Encrypter {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	if len(rounds) == 0 {
		// 64 is the standard number of rounds in tea.
		rounds = []int{64}
	}
	if rounds[0]&1 != 0 {
		e.Error = teaError.RoundsError()
		return e
	}
	src, size := e.src, tea.BlockSize
	if len(src) > size {
		src = NewCipher().EmptyPadding(e.src, size)
	}
	dst, buffer := make([]byte, size), bytes.NewBufferString("")
	for _, chunk := range bytesSplit(src, size) {
		block, err := tea.NewCipherWithRounds(key, rounds[0])
		if err != nil {
			e.Error = teaError.KeyError(key)
			return e
		}
		block.Encrypt(dst, chunk)
		buffer.Write(dst)
	}
	e.dst = buffer.Bytes()
	return e
}

// ByTea decrypts by tea.
func (d Decrypter) ByTea(key []byte, rounds ...int) Decrypter {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	if len(rounds) == 0 {
		// 64 is the standard number of rounds in tea.
		rounds = []int{64}
	}
	if rounds[0]&1 != 0 {
		d.Error = teaError.RoundsError()
		return d
	}
	src, size := d.src, tea.BlockSize
	dst, buffer := make([]byte, size), bytes.NewBufferString("")
	for _, chunk := range bytesSplit(src, size) {
		block, err := tea.NewCipherWithRounds(key, rounds[0])
		if err != nil {
			d.Error = teaError.KeyError(key)
			return d
		}
		block.Decrypt(dst, chunk)
		buffer.Write(dst)
	}
	d.dst = NewCipher().EmptyUnPadding(buffer.Bytes())
	return d
}
