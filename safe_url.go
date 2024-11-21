package dongle

import (
	"net/url"
)

// BySafeURL encodes as safe url.
func (e Encoder) BySafeURL() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	e.dst = string2bytes(url.QueryEscape(bytes2string(e.src)))
	return e
}

// BySafeURL decodes as safe url.
func (d Decoder) BySafeURL() Decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	dst, err := url.QueryUnescape(bytes2string(d.src))
	d.dst, d.Error = string2bytes(dst), err
	return d
}
