package dongle

import (
	"encoding/base64"
)

// ByBase64 encodes by base64.
// 通过 base64 编码
func (e encode) ByBase64() encode {
	if len(e.src) == 0 {
		return e
	}
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(e.src)))
	base64.StdEncoding.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByBase64 decodes by base64.
// 通过 base64 解码
func (d decode) ByBase64() decode {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(d.src)))
	n, err := base64.StdEncoding.Decode(buf, d.src)
	d.dst, d.Error = buf[:n], err
	return d
}
