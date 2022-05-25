package dongle

import (
	"crypto/cipher"
)

// decrypt defines decrypt struct
// 定义 decrypt 结构体
type decrypt struct {
	dongle
}

// newDecrypt returns a new decrypt instance.
// 初始化 decrypt 结构体
func newDecrypt() decrypt {
	return decrypt{}
}

// FromString decrypts from string without encoding.
// 对未经编码的原始字符串进行解密
func (d decrypt) FromString(s string) decrypt {
	d.src = string2bytes(s)
	return d
}

// FromHexString decrypts from string with hex encoding.
// 对经过 hex 编码的字符串进行解密
func (d decrypt) FromHexString(s string) decrypt {
	mac := Decode.FromString(s).ByHex()
	if mac.Error != nil {
		d.Error = decodeSrcError("hex")
	}
	d.src = mac.ToBytes()
	return d
}

// FromBase32String decrypts from string with base32 encoding.
// 对经过 base32 编码的字符串进行解密
func (d decrypt) FromBase32String(s string) decrypt {
	mac := Decode.FromString(s).ByBase32()
	if mac.Error != nil {
		d.Error = decodeSrcError("base32")
	}
	d.src = mac.ToBytes()
	return d
}

// FromBase64String decrypts from string with base64 encoding.
// 对经过 base64 编码的字符串进行解密
func (d decrypt) FromBase64String(s string) decrypt {
	mac := Decode.FromString(s).ByBase64()
	if mac.Error != nil {
		d.Error = decodeSrcError("base64")
	}
	d.src = mac.ToBytes()
	return d
}

// FromBytes decrypts from byte slice without encoding.
// 对未经编码的原始字节切片进行解密
func (d decrypt) FromBytes(b []byte) decrypt {
	d.src = b
	return d
}

// FromHexBytes decrypts from byte slice with hex encoding.
// 对经过 hex 编码的字节切片进行解密
func (d decrypt) FromHexBytes(b []byte) decrypt {
	mac := Decode.FromBytes(b).ByHex()
	if mac.Error != nil {
		d.Error = decodeSrcError("hex")
	}
	d.src = mac.ToBytes()
	return d
}

// FromBase32Bytes decrypts from byte slice with base32 encoding.
// 对经过 base32 编码的字节切片进行解密
func (d decrypt) FromBase32Bytes(b []byte) decrypt {
	mac := Decode.FromBytes(b).ByBase32()
	if mac.Error != nil {
		d.Error = decodeSrcError("base32")
	}
	d.src = mac.ToBytes()
	return d
}

// FromBase64Bytes decrypts from byte slice with base64 encoding.
// 对经过 base64 编码的字节切片进行解密
func (d decrypt) FromBase64Bytes(b []byte) decrypt {
	mac := Decode.FromBytes(b).ByBase64()
	if mac.Error != nil {
		d.Error = decodeSrcError("base64")
	}
	d.src = mac.ToBytes()
	return d
}

// String implements the interface Stringer for decrypt struct.
// 实现 Stringer 接口
func (d decrypt) String() string {
	return d.ToString()
}

// ToString outputs as string.
// 输出字符串
func (d decrypt) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (d decrypt) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}

// decrypts with given mode and padding
// 根据指定的分组模式和填充模式进行解密
func (d decrypt) decrypt(c *Cipher, b cipher.Block) (dst []byte, err error) {
	if c.padding == No && len(d.src)%b.BlockSize() != 0 {
		return nil, invalidSrcError(len(d.src))
	}
	if len(c.iv) != b.BlockSize() {
		return nil, invalidIVError(len(c.iv), b.BlockSize())
	}
	switch {
	case c.mode == CBC && c.padding == No:
		return c.CBCDecrypt(d.src, b), nil
	case c.mode == CBC && c.padding == Zero:
		return c.ZeroUnPadding(c.CBCDecrypt(d.src, b)), nil
	case c.mode == CBC && c.padding == PKCS5:
		return c.PKCS5UnPadding(c.CBCDecrypt(d.src, b)), nil
	case c.mode == CBC && c.padding == PKCS7:
		return c.PKCS7UnPadding(c.CBCDecrypt(d.src, b)), nil
	case c.mode == CTR && c.padding == No:
		return c.CTRDecrypt(d.src, b), nil
	case c.mode == CTR && c.padding == Zero:
		return c.ZeroUnPadding(c.CTRDecrypt(d.src, b)), nil
	case c.mode == CTR && c.padding == PKCS5:
		return c.PKCS5UnPadding(c.CTRDecrypt(d.src, b)), nil
	case c.mode == CTR && c.padding == PKCS7:
		return c.PKCS7UnPadding(c.CTRDecrypt(d.src, b)), nil
	case c.mode == CFB && c.padding == No:
		return c.CFBDecrypt(d.src, b), nil
	case c.mode == CFB && c.padding == Zero:
		return c.ZeroUnPadding(c.CFBDecrypt(d.src, b)), nil
	case c.mode == CFB && c.padding == PKCS5:
		return c.PKCS5UnPadding(c.CFBDecrypt(d.src, b)), nil
	case c.mode == CFB && c.padding == PKCS7:
		return c.PKCS7UnPadding(c.CFBDecrypt(d.src, b)), nil
	case c.mode == OFB && c.padding == No:
		return c.OFBDecrypt(d.src, b), nil
	case c.mode == OFB && c.padding == Zero:
		return c.ZeroUnPadding(c.OFBDecrypt(d.src, b)), nil
	case c.mode == OFB && c.padding == PKCS5:
		return c.PKCS5UnPadding(c.OFBDecrypt(d.src, b)), nil
	case c.mode == OFB && c.padding == PKCS7:
		return c.PKCS7UnPadding(c.OFBDecrypt(d.src, b)), nil
	}
	return nil, invalidModeOrPaddingError(c.mode, c.padding)
}
