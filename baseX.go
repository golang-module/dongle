package dongle

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"

	"github.com/dromara/dongle/base100"
	"github.com/dromara/dongle/base45"
	"github.com/dromara/dongle/base58"
	"github.com/dromara/dongle/base62"
	"github.com/dromara/dongle/base91"
)

// ByHex encodes by hex.
// 通过 hex 编码
func (e Encoder) ByHex() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	buf := make([]byte, hex.EncodedLen(len(e.src)))
	hex.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByHex decodes by hex.
// 通过 hex 解码
func (d Decoder) ByHex() Decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
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
func (e Encoder) ByBase16() Encoder {
	return e.ByHex()
}

// ByBase16 decodes by base16.
// 通过 base16 解码
func (d Decoder) ByBase16() Decoder {
	if d.ByHex().Error != nil {
		d.Error = invalidDecodingError("base16")
		return d
	}
	return d.ByHex()
}

// ByBase32 encodes by base32.
// 通过 base32 编码
func (e Encoder) ByBase32() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	buf := make([]byte, base32.StdEncoding.EncodedLen(len(e.src)))
	base32.StdEncoding.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByBase32 decodes by base32.
// 通过 base32 解码
func (d Decoder) ByBase32() Decoder {
	if len(d.src) == 0 || d.Error != nil {
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

// ByBase45 encodes by base45.
// 通过 base45 编码
func (e Encoder) ByBase45() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	dst := base45.Encode(bytes2string(e.src))
	e.dst = string2bytes(dst)
	return e
}

// ByBase45 decodes by base45.
// 通过 base45 解码
func (d Decoder) ByBase45() Decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	dst, err := base45.Decode(bytes2string(d.src))
	d.dst, d.Error = string2bytes(dst), err
	return d
}

// ByBase58 encodes by base58.
// 通过 base58 编码
func (e Encoder) ByBase58() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	e.dst = base58.Encode(e.src)
	return e
}

// ByBase58 decodes by base58.
// 通过 base58 解码
func (d Decoder) ByBase58() Decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	d.dst = base58.Decode(d.src)
	return d
}

// ByBase62 encodes by base91.
// 通过 base62 编码
func (e Encoder) ByBase62() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	e.dst = base62.StdEncoding.Encode(e.src)
	return e
}

// ByBase62 decodes by base91.
// 通过 base62 解码
func (d Decoder) ByBase62() Decoder {
	if len(d.src) == 0 || d.Error != nil {
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
func (e Encoder) ByBase64() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(e.src)))
	base64.StdEncoding.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByBase64 decodes by base64.
// 通过 base64 解码
func (d Decoder) ByBase64() Decoder {
	if len(d.src) == 0 || d.Error != nil {
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
func (e Encoder) ByBase64URL() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	buf := make([]byte, base64.URLEncoding.EncodedLen(len(e.src)))
	base64.URLEncoding.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByBase64URL decodes by base64 for url.
// 通过 base64 对 url 解码
func (d Decoder) ByBase64URL() Decoder {
	if len(d.src) == 0 || d.Error != nil {
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
func (e Encoder) ByBase85() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	buf := make([]byte, ascii85.MaxEncodedLen(len(e.src)))
	n := ascii85.Encode(buf, e.src)
	e.dst = buf[:n]
	return e
}

// ByBase85 decodes by base85.
// 通过 base85 解码
func (d Decoder) ByBase85() Decoder {
	if len(d.src) == 0 || d.Error != nil {
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
func (e Encoder) ByBase91() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	buf := make([]byte, base91.StdEncoding.EncodedLen(len(e.src)))
	n := base91.StdEncoding.Encode(buf, e.src)
	e.dst = buf[:n]
	return e
}

// ByBase91 decodes by base91.
// 通过 base91 解码
func (d Decoder) ByBase91() Decoder {
	if len(d.src) == 0 || d.Error != nil {
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
func (e Encoder) ByBase100() Encoder {
	if len(e.src) == 0 || e.Error != nil {
		return e
	}
	e.dst = base100.Encode(e.src)
	return e
}

// ByBase100 decodes by base100.
// 通过 base100 解码
func (d Decoder) ByBase100() Decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	d.dst, d.Error = base100.Decode(d.src)
	if d.Error != nil {
		d.Error = invalidDecodingError("base100")
	}
	return d
}
