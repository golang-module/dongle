package dongle

import (
	"crypto/cipher"
)

// encrypter defines a encrypter struct.
// 定义 encrypter 结构体
type encrypter struct {
	dongle
}

// returns a new encrypter instance.
// 初始化 encrypter 结构体
func newEncrypter() encrypter {
	return encrypter{}
}

// FromString encrypts from string.
// 对字符串加密
func (e encrypter) FromString(s string) encrypter {
	e.src = string2bytes(s)
	return e
}

// FromBytes encrypts from byte slice.
// 对字节切片加密
func (e encrypter) FromBytes(b []byte) encrypter {
	e.src = b
	return e
}

// String implements the interface Stringer for encrypter struct.
// 实现 Stringer 接口
func (e encrypter) String() string {
	return e.ToRawString()
}

// ToRawString outputs as raw string without encoding.
// 输出未经编码的原始字符串
func (e encrypter) ToRawString() string {
	return bytes2string(e.dst)
}

// ToHexString outputs as string with hex encoding.
// 输出经过 hex 编码的字符串
func (e encrypter) ToHexString() string {
	return Encode.FromBytes(e.dst).ByHex().ToString()
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (e encrypter) ToBase64String() string {
	return Encode.FromBytes(e.dst).ByBase64().ToString()
}

// ToRawBytes outputs as raw byte slice without encoding.
// 输出未经编码的原始字节切片
func (e encrypter) ToRawBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}

// ToHexBytes outputs as byte with hex encoding.
// 输出经过 hex 编码的字节切片
func (e encrypter) ToHexBytes() []byte {
	return Encode.FromBytes(e.dst).ByHex().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
// 输出经过 base64 编码的字节切片
func (e encrypter) ToBase64Bytes() []byte {
	return Encode.FromBytes(e.dst).ByBase64().ToBytes()
}

// encrypts with given mode and padding
// 根据指定的分组模式和填充模式进行加密
func (e encrypter) encrypt(c *Cipher, b cipher.Block) (dst []byte, err error) {
	src, mode, padding, size := e.src, c.mode, c.padding, b.BlockSize()

	if len(src) == 0 {
		return nil, nil
	}

	switch padding {
	case No:
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
		return c.NewECBEncrypter(src, b), nil
	case CBC:
		return c.NewCBCEncrypter(src, b), nil
	case CTR:
		return c.NewCTREncrypter(src, b), nil
	case CFB:
		return c.NewCFBEncrypter(src, b), nil
	case OFB:
		return c.NewOFBEncrypter(src, b), nil
	default:
		return nil, invalidModeError(mode)
	}
}
