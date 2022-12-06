package dongle

import (
	"github.com/golang-module/dongle/morse"
)

// ByMorse encodes by morse.
// 通过 morse 编码
func (e encoder) ByMorse(separator ...string) encoder {
	if len(e.src) == 0 {
		return e
	}
	if len(separator) == 0 {
		separator = []string{"/"}
	}
	dst, err := morse.Encode(e.src, separator[0])
	if err != nil {
		e.Error = invalidMorseSrcError()
		return e
	}
	e.dst = string2bytes(dst)
	return e
}

// ByMorse decodes by morse.
// 通过 morse 解码
func (d decoder) ByMorse(separator ...string) decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	if len(separator) == 0 {
		separator = []string{"/"}
	}
	dst, err := morse.Decode(d.src, separator[0])
	if err != nil {
		d.Error = morseDecodeError()
		return d
	}
	d.dst = string2bytes(dst)
	return d
}
