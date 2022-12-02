package dongle

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
)

// ByBase16 encodes by base16.
// 通过 base16 编码
func (e encode) ByBase16() encode {
	return e.ByHex()
}

// ByBase16 decodes by base16.
// 通过 base16 解码
func (d decode) ByBase16() decode {
	if d.ByHex().Error != nil {
		d.Error = invalidCiphertextError("base16")
		return d
	}
	return d.ByHex()
}

// ByBase32 encodes by base32.
// 通过 base32 编码
func (e encode) ByBase32() encode {
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
func (d decode) ByBase32() decode {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base32.StdEncoding.DecodedLen(len(d.src)))
	n, err := base32.StdEncoding.Decode(buf, d.src)
	if err != nil {
		d.Error = invalidCiphertextError("base32")
		return d
	}
	d.dst, d.Error = buf[:n], err
	return d
}

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
	if err != nil {
		d.Error = invalidCiphertextError("base64")
		return d
	}
	d.dst = buf[:n]
	return d
}

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
	if err != nil {
		d.Error = invalidCiphertextError("base64URL")
		return d
	}
	d.dst, d.Error = buf[:n], err
	return d
}

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
