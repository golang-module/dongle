package dongle

import (
	"encoding/ascii85"
)

// ByBase85 encodes by base85.
// 通过 base85 编码
func (e encode) ByBase85() encode {
	if len(e.src) == 0 {
		return e
	}
	buf := make([]byte, ascii85.MaxEncodedLen(len(e.src)))
	n := ascii85.Encode(buf, e.src)
	e.dst = buf[:n]
	return e
}

// ByBase85 decodes by base85.
// 通过 base85 解码
func (d decode) ByBase85() decode {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, len(d.src))
	n, _, err := ascii85.Decode(buf, d.src, true)
	if err != nil {
		d.Error = invalidCiphertextError("base85")
		return d
	}
	d.dst = buf[:n]
	return d
}
