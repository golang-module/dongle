package dongle

import (
	"crypto/cipher"
)

// encrypt defines a encrypt struct.
// 定义 encrypt 结构体
type encrypt struct {
	dongle
}

// returns a new encrypt instance.
// 初始化 encrypt 结构体
func newEncrypt() encrypt {
	return encrypt{}
}

// FromString encrypts from string.
// 对字符串加密
func (e encrypt) FromString(s string) encrypt {
	e.src = string2bytes(s)
	return e
}

// FromBytes encrypts from byte slice.
// 对字节切片加密
func (e encrypt) FromBytes(b []byte) encrypt {
	e.src = b
	return e
}

// String implements the interface Stringer for encrypt struct.
// 实现 Stringer 接口
func (e encrypt) String() string {
	return e.ToHexString()
}

// ToString outputs as string without encoding.
// 输出未经编码的原始字符串
func (e encrypt) ToString() string {
	return bytes2string(e.dst)
}

// ToHexString outputs as string with hex encoding.
// 输出经过 hex 编码的字符串
func (e encrypt) ToHexString() string {
	return Encode.FromBytes(e.dst).ByHex().ToString()
}

// ToBase32String outputs as string with base32 encoding.
// 输出经过 base32 编码的字符串
func (e encrypt) ToBase32String() string {
	return Encode.FromBytes(e.dst).ByBase32().ToString()
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (e encrypt) ToBase64String() string {
	return Encode.FromBytes(e.dst).ByBase64().ToString()
}

// ToBytes outputs as byte slice without encoding.
// 输出未经编码的原始字节切片
func (e encrypt) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}

// ToHexBytes outputs as byte with hex encoding.
// 输出经过 hex 编码的字节切片
func (e encrypt) ToHexBytes() []byte {
	return Encode.FromBytes(e.dst).ByHex().ToBytes()
}

// ToBase32Bytes outputs as byte slice with base32 encoding.
// 输出经过 base32 编码的字节切片
func (e encrypt) ToBase32Bytes() []byte {
	return Encode.FromBytes(e.dst).ByBase32().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
// 输出经过 base64 编码的字节切片
func (e encrypt) ToBase64Bytes() []byte {
	return Encode.FromBytes(e.dst).ByBase64().ToBytes()
}

// encrypts with given mode and padding
// 根据指定的分组模式和填充模式进行加密
func (e encrypt) encrypt(c *Cipher, b cipher.Block) (dst []byte, err error) {
	src, mode, padding, size := e.src, c.mode, c.padding, b.BlockSize()

	if len(src) == 0 {
		return nil, nil
	}

	switch padding {
	case No:
		if len(src)%size != 0 {
			return nil, invalidPlaintextError()
		}
	case Zero:
		src = c.ZeroPadding(src, size)
	case PKCS5:
		src = c.PKCS5Padding(src)
	case PKCS7:
		src = c.PKCS7Padding(src, size)
	default:
		return nil, invalidPaddingError(padding)
	}

	switch mode {
	case ECB:
		return c.ECBEncrypt(src, b), nil
	case CBC:
		return c.CBCEncrypt(src, b), nil
	case CTR:
		return c.CTREncrypt(src, b), nil
	case CFB:
		return c.CFBEncrypt(src, b), nil
	case OFB:
		return c.OFBEncrypt(src, b), nil
	default:
		return nil, invalidModeError(mode)
	}
}
