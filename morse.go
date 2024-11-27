package dongle

import (
	"fmt"

	"github.com/dromara/dongle/morse"
)

type MorseError struct {
}

func NewMorseError() MorseError {
	return MorseError{}
}

func (e MorseError) SrcError() error {
	return fmt.Errorf("morse: invalid src, the src can't contain spaces")
}

func (e MorseError) DecodeError() error {
	return fmt.Errorf("morse: failed to decode src")
}

var morseError = NewMorseError()

// ByMorse encodes by morse.
func (e Encoder) ByMorse(separator ...string) Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	if len(separator) == 0 {
		separator = []string{"/"}
	}
	dst, err := morse.Encode(e.src, separator[0])
	if err != nil {
		e.Error = morseError.SrcError()
		return e
	}
	e.dst = string2bytes(dst)
	return e
}

// ByMorse decodes by morse.
func (d Decoder) ByMorse(separator ...string) Decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	if len(separator) == 0 {
		separator = []string{"/"}
	}
	dst, err := morse.Decode(d.src, separator[0])
	if err != nil {
		d.Error = morseError.DecodeError()
		return d
	}
	d.dst = string2bytes(dst)
	return d
}
