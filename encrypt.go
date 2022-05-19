package dongle

import (
	"bufio"
	"crypto/cipher"
	"io"
	"os"
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
func (e encrypt) FromFile(s string) encrypt {
	if len(s) == 0 {
		return e
	}
	file, err := os.Open(s)
	if err != nil {
		e.Error = invalidFileError(s)
		return e
	}
	defer file.Close()

	for buf, reader := make([]byte, 1024), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			e.Error = err
			return e
		}
		if n == 0 {
			break
		}
		e.src = append(e.src, buf[:n]...)
	}
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

// encrypts with given encryption mode and padding
// 根据指定的加密模式和填充模式进行加密
func (e encrypt) encrypt(c *Cipher, b cipher.Block) (dst []byte, err error) {
	if len(e.src) == 0 {
		return nil, e.Error
	}
	if len(c.iv) != b.BlockSize() {
		return nil, invalidIVError(len(c.iv), b.BlockSize())
	}
	switch {
	case c.mode == CBC && c.padding == ZERO:
		dst = c.ZeroPadding(e.src, b.BlockSize())
		return c.CBCEncrypt(dst, b), nil
	case c.mode == CBC && c.padding == PKCS5:
		dst = c.PKCS5Padding(e.src)
		return c.CBCEncrypt(dst, b), nil
	case c.mode == CBC && c.padding == PKCS7:
		dst = c.PKCS7Padding(e.src, b.BlockSize())
		return c.CBCEncrypt(dst, b), nil
	case c.mode == CFB && c.padding == PKCS5:
		dst = c.PKCS5Padding(e.src)
		return c.CFBEncrypt(dst, b), nil
	case c.mode == CFB && c.padding == PKCS7:
		dst = c.PKCS7Padding(e.src, b.BlockSize())
		return c.CFBEncrypt(dst, b), nil
	}
	return nil, invalidModeOrPaddingError(c.mode, c.padding)
}
