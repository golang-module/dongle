package dongle

import (
	"bytes"
	"crypto/cipher"
)

// defines cipher mode enum type.
// 定义分组模式枚举类型
type cipherMode string

// cipher mode constants.
// 分组模式常量
const (
	CBC cipherMode = "cbc"
	ECB cipherMode = "ecb"
	CFB cipherMode = "cfb"
	OFB cipherMode = "ofb"
	CTR cipherMode = "ctr"
)

// defines cipher padding enum type.
// 定义填充模式枚举类型
type cipherPadding string

// cipher padding constants.
// 填充模式常量
const (
	No       cipherPadding = "no"
	Empty    cipherPadding = "empty"
	Zero     cipherPadding = "zero"
	PKCS5    cipherPadding = "pkcs5"
	PKCS7    cipherPadding = "pkcs7"
	AnsiX923 cipherPadding = "ansi-x-923"
	ISO97971 cipherPadding = "iso9797-1"
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
	c.key = interface2bytes(key)
}

// SetIV sets iv.
// 设置初始向量
func (c *Cipher) SetIV(iv interface{}) {
	c.iv = interface2bytes(iv)
}

// EmptyPadding padding with Empty mode.
// 进行空填充
func (c *Cipher) EmptyPadding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte(" "), paddingSize)
	return append(src, paddingText...)
}

// EmptyUnPadding removes padding with Empty mode.
// 移除空填充
func (c *Cipher) EmptyUnPadding(src []byte) []byte {
	return bytes.TrimRight(src, " ")
}

// ZeroPadding padding with Zero mode.
// 进行 0 填充
func (c *Cipher) ZeroPadding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(0)}, paddingSize)
	return append(src, paddingText...)
}

// ZeroUnPadding removes padding with Zero mode.
// 移除 0 填充
func (c *Cipher) ZeroUnPadding(src []byte) []byte {
	return bytes.TrimRight(src, string([]byte{0}))
}

// ISO97971Padding padding with ISO/IEC 9797-1 mode.
// 进行 ISO/IEC 9797-1 填充
func (c *Cipher) ISO97971Padding(src []byte, blockSize int) []byte {
	return c.ZeroPadding(append(src, 0x80), blockSize)
}

// ISO97971UnPadding removes padding with ISO/IEC 9797-1 mode.
// 移除 ISO/IEC 9797-1 填充
func (c *Cipher) ISO97971UnPadding(src []byte) []byte {
	dst := c.ZeroUnPadding(src)
	return dst[:len(dst)-1]
}

// AnsiX923Padding padding with ANSI X.923 mode.
// 进行 ANSI X.923 填充
func (c *Cipher) AnsiX923Padding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := append(bytes.Repeat([]byte{byte(0)}, paddingSize-1), byte(paddingSize))
	return append(src, paddingText...)
}

// AnsiX923UnPadding removes padding with ANSI X.923 mode.
// 移除 ANSI X.923 填充
func (c *Cipher) AnsiX923UnPadding(src []byte) []byte {
	n := len(src) - int(src[len(src)-1])
	return src[0:n]
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
	n := len(src) - int(src[len(src)-1])
	return src[0:n]
}

// NewCBCEncrypter encrypts with CBC mode.
// CBC 模式加密
func (c *Cipher) NewCBCEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCEncrypter(block, c.iv).CryptBlocks(dst, src)
	return
}

// NewCBCDecrypter decrypts with CBC mode.
// CBC 模式解密
func (c *Cipher) NewCBCDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCDecrypter(block, c.iv).CryptBlocks(dst, src)
	return
}

// NewCFBEncrypter encrypts with CFB mode.
// CFB 模式加密
func (c *Cipher) NewCFBEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBEncrypter(block, c.iv).XORKeyStream(dst, src)
	return
}

// NewCFBDecrypter decrypts with CFB mode.
// CFB 模式解密
func (c *Cipher) NewCFBDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBDecrypter(block, c.iv).XORKeyStream(dst, src)
	return
}

// NewCTREncrypter encrypts with CTR mode.
// CTR 模式加密
func (c *Cipher) NewCTREncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv).XORKeyStream(dst, src)
	return
}

// NewCTRDecrypter decrypts with CTR mode.
// CTR 模式解密
func (c *Cipher) NewCTRDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv).XORKeyStream(dst, src)
	return
}

// NewECBEncrypter encrypts with ECB mode.
// ECB 模式加密
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
// ECB 模式解密
func (c *Cipher) NewECBDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	decrypted, blockSize := dst, block.BlockSize()
	for len(src) > 0 {
		block.Decrypt(decrypted, src[:blockSize])
		src = src[blockSize:]
		decrypted = decrypted[blockSize:]
	}
	return
}

// NewOFBEncrypter encrypts with OFB mode.
// OFB 模式加密
func (c *Cipher) NewOFBEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewOFB(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// NewOFBDecrypter decrypts with OFB mode.
// OFB 模式解密
func (c *Cipher) NewOFBDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewOFB(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// whether is a supported padding
// 判断是否是支持的填充模式
func (c *Cipher) isSupportedPadding() bool {
	paddings := []cipherPadding{
		No, Empty, Zero, PKCS5, PKCS7, AnsiX923, ISO97971,
	}
	for _, padding := range paddings {
		if padding == c.padding {
			return true
		}
	}
	return false
}

// encrypts with given mode and padding
// 根据指定的分组模式和填充模式进行加密
func (c *Cipher) encrypt(src []byte, block cipher.Block) (dst []byte, err error) {
	mode, padding, size := c.mode, c.padding, block.BlockSize()
	dst = []byte("")
	if len(src) == 0 {
		return
	}

	switch padding {
	case No:
	case Empty:
		src = c.EmptyPadding(src, size)
	case Zero:
		src = c.ZeroPadding(src, size)
	case PKCS5:
		src = c.PKCS5Padding(src)
	case PKCS7:
		src = c.PKCS7Padding(src, size)
	case AnsiX923:
		src = c.AnsiX923Padding(src, size)
	case ISO97971:
		src = c.ISO97971Padding(src, size)
	default:
		err = invalidPaddingError(padding)
		return
	}

	switch mode {
	case ECB:
		return c.NewECBEncrypter(src, block), nil
	case CBC:
		return c.NewCBCEncrypter(src, block), nil
	case CTR:
		return c.NewCTREncrypter(src, block), nil
	case CFB:
		return c.NewCFBEncrypter(src, block), nil
	case OFB:
		return c.NewOFBEncrypter(src, block), nil
	default:
		err = invalidModeError(mode)
		return
	}
}

// decrypts with given mode and padding.
// 根据指定的分组模式和填充模式进行解密
func (c *Cipher) decrypt(src []byte, block cipher.Block) (dst []byte, err error) {
	mode, padding := c.mode, c.padding
	dst = []byte("")
	if len(src) == 0 {
		return
	}
	if !c.isSupportedPadding() {
		err = invalidPaddingError(padding)
		return
	}

	switch mode {
	case ECB:
		src = c.NewECBDecrypter(src, block)
	case CBC:
		src = c.NewCBCDecrypter(src, block)
	case CTR:
		src = c.NewCTRDecrypter(src, block)
	case CFB:
		src = c.NewCFBDecrypter(src, block)
	case OFB:
		src = c.NewOFBDecrypter(src, block)
	default:
		err = invalidModeError(mode)
		return
	}

	switch padding {
	case Zero:
		return c.ZeroUnPadding(src), nil
	case Empty:
		return c.EmptyUnPadding(src), nil
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
