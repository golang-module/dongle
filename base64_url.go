package dongle

import (
	"encoding/base64"
)

// ByBase64URL encodes by base64 for url.
// 通过 base64 对 url 编码
func (e encode) ByBase64URL() encode {
	if len(e.src) == 0 {
		return e
	}
	buf := make([]byte, base64.URLEncoding.EncodedLen(len(e.src)))
	base64.URLEncoding.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByBase64URL decodes by base64 for url.
// 通过 base64 对 url 解码
func (d decode) ByBase64URL() decode {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base64.URLEncoding.DecodedLen(len(d.src)))
	n, err := base64.URLEncoding.Decode(buf, d.src)
	d.dst, d.Error = buf[:n], err
	return d
}
