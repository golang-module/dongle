package dongle

import (
	"crypto/cipher"
	"fmt"
)

// returns an invalid ciphertext error
// 返回无效的密文错误
var invalidCiphertextError = func(mode string) error {
	return fmt.Errorf("invalid ciphertext, the ciphertext can't be decoded by %s", mode)
}

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
		d.Error = invalidCiphertextError("hex")
		return d
	}
	d.src = mac.ToBytes()
	return d
}

// FromBase32String decrypts from string with base32 encoding.
// 对经过 base32 编码的字符串进行解密
func (d decrypt) FromBase32String(s string) decrypt {
	mac := Decode.FromString(s).ByBase32()
	if mac.Error != nil {
		d.Error = invalidCiphertextError("base32")
		return d
	}
	d.src = mac.ToBytes()
	return d
}

// FromBase64String decrypts from string with base64 encoding.
// 对经过 base64 编码的字符串进行解密
func (d decrypt) FromBase64String(s string) decrypt {
	mac := Decode.FromString(s).ByBase64()
	if mac.Error != nil {
		d.Error = invalidCiphertextError("base64")
		return d
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
		d.Error = invalidCiphertextError("hex")
		return d
	}
	d.src = mac.ToBytes()
	return d
}

// FromBase32Bytes decrypts from byte slice with base32 encoding.
// 对经过 base32 编码的字节切片进行解密
func (d decrypt) FromBase32Bytes(b []byte) decrypt {
	mac := Decode.FromBytes(b).ByBase32()
	if mac.Error != nil {
		d.Error = invalidCiphertextError("base32")
		return d
	}
	d.src = mac.ToBytes()
	return d
}

// FromBase64Bytes decrypts from byte slice with base64 encoding.
// 对经过 base64 编码的字节切片进行解密
func (d decrypt) FromBase64Bytes(b []byte) decrypt {
	mac := Decode.FromBytes(b).ByBase64()
	if mac.Error != nil {
		d.Error = invalidCiphertextError("base64")
		return d
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

// decrypts with given mode and padding.
// 根据指定的分组模式和填充模式进行解密
func (d decrypt) decrypt(c *Cipher, b cipher.Block) (dst []byte, err error) {
	src, mode, padding, size := d.src, c.mode, c.padding, b.BlockSize()
	if len(src) == 0 {
		return nil, nil
	}
	if padding != No && padding != Zero && padding != PKCS5 && padding != PKCS7 {
		return nil, invalidPaddingError(padding)
	}
	if padding == No && len(src)%size != 0 {
		return nil, invalidPlaintextError()
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
	}
	return src, nil
}
