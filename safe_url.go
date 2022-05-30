package dongle

import (
	"net/url"
)

// BySafeURL encodes as safe url.
// 对 URL 进行转义编码
func (e encode) BySafeURL() encode {
	if len(e.src) == 0 {
		return e
	}
	e.dst = string2bytes(url.QueryEscape(bytes2string(e.src)))
	return e
}

// BySafeURL decodes as safe url.
// 对 URL 进行转义解码
func (d decode) BySafeURL() decode {
	if len(d.src) == 0 {
		return d
	}
	dst, err := url.QueryUnescape(bytes2string(d.src))
	d.dst, d.Error = string2bytes(dst), err
	return d
}
