package dongle

import (
	"crypto/cipher"
)

// Decrypter defines Decrypter struct.
// 定义 Decrypter 结构体
type Decrypter struct {
	dongle
}

// newDecrypter returns a new Decrypter instance.
// 初始化 Decrypter 结构体
func newDecrypter() Decrypter {
	return Decrypter{}
}

// FromRawString decrypts from raw string without encoding.
// 对未经编码的原始字符串进行解密
func (d Decrypter) FromRawString(s string) Decrypter {
	d.src = string2bytes(s)
	return d
}

// FromHexString decrypts from string with hex encoding.
// 对经过 hex 编码的字符串进行解密
func (d Decrypter) FromHexString(s string) Decrypter {
	decode := Decode.FromString(s).ByHex()
	if decode.Error != nil {
		d.Error = invalidDecodingError("hex")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromBase64String decrypts from string with base64 encoding.
// 对经过 base64 编码的字符串进行解密
func (d Decrypter) FromBase64String(s string) Decrypter {
	decode := Decode.FromString(s).ByBase64()
	if decode.Error != nil {
		d.Error = invalidDecodingError("base64")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromRawBytes decrypts from raw byte slice without encoding.
// 对未经编码的原始字节切片进行解密
func (d Decrypter) FromRawBytes(b []byte) Decrypter {
	d.src = b
	return d
}

// FromHexBytes decrypts from byte slice with hex encoding.
// 对经过 hex 编码的字节切片进行解密
func (d Decrypter) FromHexBytes(b []byte) Decrypter {
	decode := Decode.FromBytes(b).ByHex()
	if decode.Error != nil {
		d.Error = invalidDecodingError("hex")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromBase64Bytes decrypts from byte slice with base64 encoding.
// 对经过 base64 编码的字节切片进行解密
func (d Decrypter) FromBase64Bytes(b []byte) Decrypter {
	decode := Decode.FromBytes(b).ByBase64()
	if decode.Error != nil {
		d.Error = invalidDecodingError("base64")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// String implements Stringer interface for Decrypter struct.
// 实现 Stringer 接口
func (d Decrypter) String() string {
	return d.ToString()
}

// ToString outputs as string.
// 输出字符串
func (d Decrypter) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (d Decrypter) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}

// decrypts with given mode and padding.
// 根据指定的分组模式和填充模式进行解密
func (d Decrypter) decrypt(c *Cipher, b cipher.Block) (dst []byte, err error) {
	src, mode, padding := d.src, c.mode, c.padding
	if len(src) == 0 {
		return nil, nil
	}
	if !padding.isSupported() {
		return nil, invalidPaddingError(padding)
	}

	switch mode {
	case ECB:
		src = c.NewECBDecrypter(src, b)
	case CBC:
		src = c.NewCBCDecrypter(src, b)
	case CTR:
		src = c.NewCTRDecrypter(src, b)
	case CFB:
		src = c.NewCFBDecrypter(src, b)
	case OFB:
		src = c.NewOFBDecrypter(src, b)
	default:
		return nil, invalidModeError(mode)
	}

	switch padding {
	case Zero:
		return c.ZeroUnPadding(src), nil
	case PKCS5:
		return c.PKCS5UnPadding(src), nil
	case PKCS7:
		return c.PKCS7UnPadding(src), nil
	case AnsiX923:
		return c.AnsiX923UnPadding(src), nil
	case ISO97971:
		return c.ISO97971UnPadding(src), nil
	}
	return src, nil
}
