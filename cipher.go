package dongle

import (
	"bytes"
	"crypto/cipher"
)

// group mode constant
const (
	CBC = "cbc"
	ECB = "ecb"
	CFB = "cfb"
	OFB = "ofb"
	CTR = "ctr"
	GCM = "gcm"
)

// padding mode constant
const (
	NoPadding       = "no_padding"
	ZeroPadding     = "zero_padding"
	PKCS1Padding    = "pkcs1_padding"
	PKCS5Padding    = "pkcs5_padding"
	PKCS7Padding    = "pkcs7_padding"
	TBCPadding      = "tbc_padding"
	AnsiX923Padding = "ansi_x923_padding"
	ISO10126Padding = "iso10226_padding"
)

// Cipher defines a Cipher struct.
type Cipher struct {
	mode       string // 分组模式
	padding    string // 填充模式
	key        []byte // 密钥
	iv         []byte // 初始向量
	publicKey  []byte // 公钥
	privateKey []byte // 私钥
}

// NewCipher returns a new Cipher instance.
func NewCipher() *Cipher {
	return &Cipher{
		mode:    CBC,
		padding: PKCS7Padding,
	}
}

// SetMode sets group mode.
func (c *Cipher) SetMode(mode string) {
	c.mode = mode
}

// SetPadding sets padding mode.
func (c *Cipher) SetPadding(padding string) {
	c.padding = padding
}

// SetKey sets key.
func (c *Cipher) SetKey(key interface{}) {
	switch v := key.(type) {
	case string:
		c.key = string2bytes(v)
	case []byte:
		c.key = v
	}
}

// SetIV sets initial vector.
func (c *Cipher) SetIV(iv interface{}) {
	switch v := iv.(type) {
	case string:
		c.iv = string2bytes(v)
	case []byte:
		c.iv = v
	}
}

// ZeroPadding padding with zero mode.
func (c *Cipher) ZeroPadding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	dst := bytes.Repeat([]byte{0}, padding)
	return append(src, dst...)
}

// ZeroUnPadding unpadding with zero mode.
func (c *Cipher) ZeroUnPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return emptyBytes
}

// PKCS5Padding padding with pkcs5 mode.
func (c *Cipher) PKCS5Padding(src []byte) []byte {
	return c.PKCS7Padding(src, 8)
}

// PKCS5UnPadding unpadding with pkcs5 mode.
func (c *Cipher) PKCS5UnPadding(src []byte) []byte {
	return c.PKCS7UnPadding(src)
}

// PKCS7Padding padding with pkcs7 mode.
func (c *Cipher) PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	dst := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, dst...)
}

// PKCS7UnPadding unpadding with pkcs7 mode.
func (c *Cipher) PKCS7UnPadding(src []byte) []byte {
	return src[:(len(src) - int(src[len(src)-1]))]
}

// CBCEncrypt returns a byte slice which encrypts with cbc mode.
func (c *Cipher) CBCEncrypt(block cipher.Block, src []byte) []byte {
	dst, size := make([]byte, len(src)), block.BlockSize()
	cipher.NewCBCEncrypter(block, c.iv[:size]).CryptBlocks(dst, src)
	return dst
}

// CBCDecrypt returns a byte slice which decrypts with cbc mode.
func (c *Cipher) CBCDecrypt(block cipher.Block, src []byte) []byte {
	dst, size := make([]byte, len(src)), block.BlockSize()
	cipher.NewCBCDecrypter(block, c.iv[:size]).CryptBlocks(dst, src)
	return dst
}
