package dongle

import (
	"crypto/cipher"
)

// decrypt defines decrypt struct
type decrypt struct {
	dongle
}

// newDecrypt returns a new decrypt instance.
func newDecrypt() decrypt {
	return decrypt{}
}

// FromString encrypts from string.
func (d decrypt) FromString(s string) decrypt {
	d.src = string2bytes(s)
	return d
}

func (d decrypt) FromHexString(s string) decrypt {
	d.src = Decode.FromString(s).ByHex().ToBytes()
	return d
}

func (d decrypt) FromBase32String(s string) decrypt {
	d.src = Decode.FromString(s).ByBase32().ToBytes()
	return d
}

func (d decrypt) FromBase64String(s string) decrypt {
	d.src = Decode.FromString(s).ByBase64().ToBytes()
	return d
}

// FromBytes encrypts from byte slice.
func (d decrypt) FromBytes(b []byte) decrypt {
	d.src = b
	return d
}

func (d decrypt) FromHexBytes(b []byte) decrypt {
	d.src = Decode.FromBytes(b).ByHex().ToBytes()
	return d
}

func (d decrypt) FromBase32Bytes(b []byte) decrypt {
	d.src = Decode.FromBytes(b).ByBase32().ToBytes()
	return d
}

func (d decrypt) FromBase64Bytes(b []byte) decrypt {
	d.src = Decode.FromBytes(b).ByBase64().ToBytes()
	return d
}

// String implements the interface Stringer for decrypt struct.
// 实现 Stringer 接口
func (d decrypt) String() string {
	return d.ToString()
}

// ToString outputs as string.
func (d decrypt) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
func (d decrypt) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}

func (d decrypt) decrypt(c *Cipher, b cipher.Block) (dst []byte, err error) {
	if len(d.src) == 0 {
		return nil, d.Error
	}
	if len(c.iv) != b.BlockSize() {
		return nil, invalidIVError(len(c.iv), b.BlockSize())
	}
	switch {
	case c.mode == CBC && c.padding == ZERO:
		dst = c.CBCDecrypt(d.src, b)
		dst = c.ZeroUnPadding(dst)
	case c.mode == CBC && c.padding == PKCS5:
		dst = c.CBCDecrypt(d.src, b)
		dst = c.PKCS5UnPadding(dst)
	case c.mode == CBC && c.padding == PKCS7:
		dst = c.CBCDecrypt(d.src, b)
		dst = c.PKCS7UnPadding(dst)
	case c.mode == CFB && c.padding == ZERO:
		dst = c.CFBDecrypt(d.src, b)
		dst = c.ZeroUnPadding(dst)
	case c.mode == CFB && c.padding == PKCS5:
		dst = c.CFBDecrypt(d.src, b)
		dst = c.PKCS5UnPadding(dst)
	case c.mode == CFB && c.padding == PKCS7:
		dst = c.CFBDecrypt(d.src, b)
		dst = c.PKCS7UnPadding(dst)
	default:
		return nil, invalidModeOrPaddingError(c.mode, c.padding)
	}
	return dst, nil
}
