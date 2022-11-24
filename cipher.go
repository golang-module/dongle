package dongle

import (
	"bytes"
	"crypto/cipher"
)

type cipherMode string

// mode constants
// 模式常量
const (
	CBC cipherMode = "cbc"
	ECB cipherMode = "ecb"
	CFB cipherMode = "cfb"
	OFB cipherMode = "ofb"
	CTR cipherMode = "ctr"
)

type cipherPadding string

// padding constants
// 填充常量
const (
	No    cipherPadding = "no"
	Zero  cipherPadding = "zero"
	PKCS5 cipherPadding = "pkcs5"
	PKCS7 cipherPadding = "pkcs7"
)

// Cipher defines a Cipher struct.
// 定义 Cipher 结构体
type Cipher struct {
	mode    cipherMode    // 分组模式
	padding cipherPadding // 填充模式
	key     []byte        // 密钥
	iv      []byte        // 初始向量
}

// NewCipher returns a new Cipher instance.
// 初始化 Cipher 结构体
func NewCipher() *Cipher {
	return &Cipher{
		mode:    CBC,
		padding: PKCS7,
	}
}

// SetMode sets mode.
// 设置分组模式
func (c *Cipher) SetMode(mode cipherMode) {
	c.mode = mode
}

// SetPadding sets padding.
// 设置填充模式
func (c *Cipher) SetPadding(padding cipherPadding) {
	c.padding = padding
}

// SetKey sets key.
// 设置密钥
func (c *Cipher) SetKey(key interface{}) {
	switch v := key.(type) {
	case string:
		c.key = string2bytes(v)
	case []byte:
		c.key = v
	}
}

// SetIV sets iv.
// 设置初始向量
func (c *Cipher) SetIV(iv interface{}) {
	switch v := iv.(type) {
	case string:
		c.iv = string2bytes(v)
	case []byte:
		c.iv = v
	}
}

// ZeroPadding padding with Zero mode.
// 进行零填充
func (c *Cipher) ZeroPadding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(0)}, paddingSize)
	return append(src, paddingText...)
}

// ZeroUnPadding removes padding with Zero mode.
// 移除零填充
func (c *Cipher) ZeroUnPadding(src []byte) []byte {
	return bytes.TrimRight(src, string([]byte{0}))
}

// PKCS5Padding padding with PKCS5 mode.
// 进行 PKCS5 填充
func (c *Cipher) PKCS5Padding(src []byte) []byte {
	return c.PKCS7Padding(src, 8)
}

// PKCS5UnPadding removes padding with PKCS5 mode.
// 移除 PKCS5 填充
func (c *Cipher) PKCS5UnPadding(src []byte) []byte {
	return c.PKCS7UnPadding(src)
}

// PKCS7Padding padding with PKCS7 mode.
// 进行 PKCS7 填充
func (c *Cipher) PKCS7Padding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(src, paddingText...)
}

// PKCS7UnPadding removes padding with PKCS7 mode.
// 移除 PKCS7 填充
func (c *Cipher) PKCS7UnPadding(src []byte) []byte {
	trim := len(src) - int(src[len(src)-1])
	return src[:trim]
}

// NewCBCEncrypter encrypts with CBC mode.
// CBC 加密
func (c *Cipher) NewCBCEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCEncrypter(block, c.iv[:block.BlockSize()]).CryptBlocks(dst, src)
	return
}

// NewCBCDecrypter decrypts with CBC mode.
// CBC 解密
func (c *Cipher) NewCBCDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCDecrypter(block, c.iv[:block.BlockSize()]).CryptBlocks(dst, src)
	return
}

// NewCFBEncrypter encrypts with CFB mode.
// CFB 加密
func (c *Cipher) NewCFBEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBEncrypter(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// NewCFBDecrypter decrypts with CFB mode.
// CFB 解密
func (c *Cipher) NewCFBDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBDecrypter(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// NewCTREncrypter encrypts with CTR mode.
// CTR 加密
func (c *Cipher) NewCTREncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// NewCTRDecrypter decrypts with CTR mode.
// CTR 解密
func (c *Cipher) NewCTRDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// NewECBEncrypter encrypts with ECB mode.
// ECB 加密
func (c *Cipher) NewECBEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	encrypted, size := dst, block.BlockSize()
	for len(src) > 0 {
		block.Encrypt(encrypted, src[:size])
		src = src[size:]
		encrypted = encrypted[size:]
	}
	return
}

// NewECBDecrypter decrypts with ECB mode.
// ECB 解密
func (c *Cipher) NewECBDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	decrypted, size := dst, block.BlockSize()
	for len(src) > 0 {
		block.Decrypt(decrypted, src[:size])
		src = src[size:]
		decrypted = decrypted[size:]
	}
	return
}

// NewOFBEncrypter encrypts with OFB mode.
// OFB 加密
func (c *Cipher) NewOFBEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewOFB(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// NewOFBDecrypter decrypts with OFB mode.
// OFB 解密
func (c *Cipher) NewOFBDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewOFB(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}
