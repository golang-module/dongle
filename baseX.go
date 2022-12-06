package dongle

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"github.com/golang-module/dongle/base100"
	"github.com/golang-module/dongle/base58"
	"github.com/golang-module/dongle/base62"
	"github.com/golang-module/dongle/base91"
)

// ByHex encodes by hex.
// 通过 hex 编码
func (e encoder) ByHex() encoder {
	buf := make([]byte, hex.EncodedLen(len(e.src)))
	hex.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByHex decodes by hex.
// 通过 hex 解码
func (d decoder) ByHex() decoder {
	buf := make([]byte, hex.DecodedLen(len(d.src)))
	n, err := hex.Decode(buf, d.src)
	if n > 0 {
		d.dst = buf
	}
	if err != nil {
		d.Error = invalidDecodingError("hex")
		return d
	}
	return d
}

// ByBase16 encodes by base16.
// 通过 base16 编码
func (e encoder) ByBase16() encoder {
	return e.ByHex()
}

// ByBase16 decodes by base16.
// 通过 base16 解码
func (d decoder) ByBase16() decoder {
	if d.ByHex().Error != nil {
		d.Error = invalidDecodingError("base16")
		return d
	}
	return d.ByHex()
}

// ByBase32 encodes by base32.
// 通过 base32 编码
func (e encoder) ByBase32() encoder {
	if len(e.src) == 0 {
		return e
	}
	buf := make([]byte, base32.StdEncoding.EncodedLen(len(e.src)))
	base32.StdEncoding.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByBase32 decodes by base32.
// 通过 base32 解码
func (d decoder) ByBase32() decoder {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base32.StdEncoding.DecodedLen(len(d.src)))
	n, err := base32.StdEncoding.Decode(buf, d.src)
	if err != nil {
		d.Error = invalidDecodingError("base32")
		return d
	}
	d.dst, d.Error = buf[:n], err
	return d
}

// ByBase58 encodes by base58.
// 通过 base58 编码
func (e encoder) ByBase58() encoder {
	if len(e.src) == 0 {
		return e
	}
	e.dst = base58.Encode(e.src)
	return e
}

// ByBase58 decodes by base58.
// 通过 base58 解码
func (d decoder) ByBase58() decoder {
	if len(d.src) == 0 {
		return d
	}
	d.dst = base58.Decode(d.src)
	return d
}

// ByBase62 encodes by base91.
// 通过 base62 编码
func (e encoder) ByBase62() encoder {
	if len(e.src) == 0 {
		return e
	}
	e.dst = base62.StdEncoding.Encode(e.src)
	return e
}

// ByBase62 decodes by base91.
// 通过 base62 解码
func (d decoder) ByBase62() decoder {
	if len(d.src) == 0 {
		return d
	}
	d.dst, d.Error = base62.StdEncoding.Decode(d.src)
	if d.Error != nil {
		d.Error = invalidDecodingError("base62")
		return d
	}
	return d
}

// ByBase64 encodes by base64.
// 通过 base64 编码
func (e encoder) ByBase64() encoder {
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
func (d decoder) ByBase64() decoder {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(d.src)))
	n, err := base64.StdEncoding.Decode(buf, d.src)
	if err != nil {
		d.Error = invalidDecodingError("base64")
		return d
	}
	d.dst = buf[:n]
	return d
}

// ByBase64URL encodes by base64 for url.
// 通过 base64 对 url 编码
func (e encoder) ByBase64URL() encoder {
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
func (d decoder) ByBase64URL() decoder {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base64.URLEncoding.DecodedLen(len(d.src)))
	n, err := base64.URLEncoding.Decode(buf, d.src)
	if err != nil {
		d.Error = invalidDecodingError("base64URL")
		return d
	}
	d.dst, d.Error = buf[:n], err
	return d
}

// ByBase85 encodes by base85.
// 通过 base85 编码
func (e encoder) ByBase85() encoder {
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
func (d decoder) ByBase85() decoder {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, len(d.src))
	n, _, err := ascii85.Decode(buf, d.src, true)
	if err != nil {
		d.Error = invalidDecodingError("base85")
		return d
	}
	d.dst = buf[:n]
	return d
}

// ByBase91 encodes by base91.
// 通过 base91 编码
func (e encoder) ByBase91() encoder {
	if len(e.src) == 0 {
		return e
	}
	buf := make([]byte, base91.StdEncoding.EncodedLen(len(e.src)))
	n := base91.StdEncoding.Encode(buf, e.src)
	e.dst = buf[:n]
	return e
}

// ByBase91 decodes by base91.
// 通过 base91 解码
func (d decoder) ByBase91() decoder {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base91.StdEncoding.DecodedLen(len(d.src)))
	n, err := base91.StdEncoding.Decode(buf, d.src)
	if err != nil {
		d.Error = invalidDecodingError("base91")
		return d
	}
	d.dst = buf[:n]
	return d
}

// ByBase100 encodes by base100.
// 通过 base100 编码
func (e encoder) ByBase100() encoder {
	if len(e.src) == 0 {
		return e
	}
	e.dst = base100.Encode(e.src)
	return e
}

// ByBase100 decodes by base100.
// 通过 base100 解码
func (d decoder) ByBase100() decoder {
	if len(d.src) == 0 {
		return d
	}
	d.dst, d.Error = base100.Decode(d.src)
	if d.Error != nil {
		d.Error = invalidDecodingError("base100")
	}
	return d
}
