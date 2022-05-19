package dongle

import (
	"encoding/hex"
)

// ByHex encodes by hex.
// 通过 hex 编码
func (e encode) ByHex() encode {
	buf := make([]byte, hex.EncodedLen(len(e.src)))
	hex.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByHex decodes by hex.
// 通过 hex 解码
func (d decode) ByHex() decode {
	buf := make([]byte, hex.DecodedLen(len(d.src)))
	n, err := hex.Decode(buf, d.src)
	if n > 0 {
		d.dst = buf
	}
	d.Error = err
	return d
}
