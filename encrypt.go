package dongle

import (
	"crypto/cipher"
	"io/ioutil"
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

// FromFile encrypts from file.
// 对文件加密
func (e encrypt) FromFile(f interface{}) encrypt {
	filename := ""
	switch v := f.(type) {
	case string:
		filename = v
	case []byte:
		filename = bytes2string(v)
	}
	if len(filename) == 0 {
		return e
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		e.Error = invalidFileError(filename)
		return e
	}
	e.src = bytes
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
	if c.padding == No && len(e.src)%b.BlockSize() != 0 {
		return nil, invalidSrcError(len(e.src))
	}
	if len(c.iv) != b.BlockSize() {
		return nil, invalidIVError(len(c.iv), b.BlockSize())
	}
	switch {
	case c.mode == CBC && c.padding == No:
		return c.CBCEncrypt(e.src, b), nil
	case c.mode == CBC && c.padding == Zero:
		return c.CBCEncrypt(c.ZeroPadding(e.src, b.BlockSize()), b), nil
	case c.mode == CBC && c.padding == PKCS5:
		return c.CBCEncrypt(c.PKCS5Padding(e.src), b), nil
	case c.mode == CBC && c.padding == PKCS7:
		return c.CBCEncrypt(c.PKCS7Padding(e.src, b.BlockSize()), b), nil
	case c.mode == CTR && c.padding == No:
		return c.CTREncrypt(e.src, b), nil
	case c.mode == CTR && c.padding == Zero:
		return c.CTREncrypt(c.ZeroPadding(e.src, b.BlockSize()), b), nil
	case c.mode == CTR && c.padding == PKCS5:
		return c.CTREncrypt(c.PKCS5Padding(e.src), b), nil
	case c.mode == CTR && c.padding == PKCS7:
		return c.CTREncrypt(c.PKCS7Padding(e.src, b.BlockSize()), b), nil
	case c.mode == CFB && c.padding == No:
		return c.CFBEncrypt(e.src, b), nil
	case c.mode == CFB && c.padding == Zero:
		return c.CFBEncrypt(c.ZeroPadding(e.src, b.BlockSize()), b), nil
	case c.mode == CFB && c.padding == PKCS5:
		return c.CFBEncrypt(c.PKCS5Padding(e.src), b), nil
	case c.mode == CFB && c.padding == PKCS7:
		return c.CFBEncrypt(c.PKCS7Padding(e.src, b.BlockSize()), b), nil
	case c.mode == OFB && c.padding == No:
		return c.OFBEncrypt(e.src, b), nil
	case c.mode == OFB && c.padding == Zero:
		return c.OFBEncrypt(c.ZeroPadding(e.src, b.BlockSize()), b), nil
	case c.mode == OFB && c.padding == PKCS5:
		return c.OFBEncrypt(c.PKCS5Padding(e.src), b), nil
	case c.mode == OFB && c.padding == PKCS7:
		return c.OFBEncrypt(c.PKCS7Padding(e.src, b.BlockSize()), b), nil
	}
	return nil, invalidModeOrPaddingError(c.mode, c.padding)
}
