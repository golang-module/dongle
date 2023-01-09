package dongle

import (
	"gitee.com/golang-package/dongle/morse"
)

// ByMorse encodes by morse.
// 通过 morse 编码
func (e Encoder) ByMorse(separator ...string) Encoder {
	if len(e.src) == 0 || e.Error != nil {
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
func (d Decoder) ByMorse(separator ...string) Decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	if len(separator) == 0 {
		separator = []string{"/"}
	}
	dst, err := morse.Decode(d.src, separator[0])
	if err != nil {
		d.Error = morseDecodingError()
		return d
	}
	d.dst = string2bytes(dst)
	return d
}
