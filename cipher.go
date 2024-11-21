package dongle

import (
	"bytes"
	"crypto/cipher"
)

type cipherMode string

const (
	CBC cipherMode = "cbc"
	ECB cipherMode = "ecb"
	CFB cipherMode = "cfb"
	OFB cipherMode = "ofb"
	CTR cipherMode = "ctr"
)

type cipherPadding string

const (
	No       cipherPadding = "no"
	Empty    cipherPadding = "empty"
	Zero     cipherPadding = "zero"
	PKCS5    cipherPadding = "pkcs5"
	PKCS7    cipherPadding = "pkcs7"
	AnsiX923 cipherPadding = "ansi-x-923"
	ISO97971 cipherPadding = "iso9797-1"
)

type Cipher struct {
	mode    cipherMode
	padding cipherPadding
	key     []byte
	iv      []byte
}

// NewCipher returns a new Cipher instance.
func NewCipher() *Cipher {
	return &Cipher{
		mode:    CBC,
		padding: PKCS7,
	}
}

// SetMode sets mode.
func (c *Cipher) SetMode(mode cipherMode) {
	c.mode = mode
}

// SetPadding sets padding.
func (c *Cipher) SetPadding(padding cipherPadding) {
	c.padding = padding
}

// SetKey sets key string.
func (c *Cipher) SetKey(key string) {
	c.key = string2bytes(key)
}

// SetIV sets iv string.
func (c *Cipher) SetIV(iv string) {
	c.iv = string2bytes(iv)
}

// WithKey sets key slice.
func (c *Cipher) WithKey(key []byte) {
	c.key = key
}

// WithIV sets iv slice.
func (c *Cipher) WithIV(iv []byte) {
	c.iv = iv
}

// EmptyPadding padding with Empty mode.
func (c *Cipher) EmptyPadding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte(" "), paddingSize)
	return append(src, paddingText...)
}

// EmptyUnPadding removes padding with Empty mode.
func (c *Cipher) EmptyUnPadding(src []byte) []byte {
	return bytes.TrimRight(src, " ")
}

// ZeroPadding padding with Zero mode.
func (c *Cipher) ZeroPadding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(0)}, paddingSize)
	return append(src, paddingText...)
}

// ZeroUnPadding removes padding with Zero mode.
func (c *Cipher) ZeroUnPadding(src []byte) []byte {
	return bytes.TrimRight(src, string([]byte{0}))
}

// ISO97971Padding padding with ISO/IEC 9797-1 mode.
func (c *Cipher) ISO97971Padding(src []byte, blockSize int) []byte {
	return c.ZeroPadding(append(src, 0x80), blockSize)
}

// ISO97971UnPadding removes padding with ISO/IEC 9797-1 mode.
func (c *Cipher) ISO97971UnPadding(src []byte) []byte {
	dst := c.ZeroUnPadding(src)
	return dst[:len(dst)-1]
}

// AnsiX923Padding padding with ANSI X.923 mode.
func (c *Cipher) AnsiX923Padding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := append(bytes.Repeat([]byte{byte(0)}, paddingSize-1), byte(paddingSize))
	return append(src, paddingText...)
}

// AnsiX923UnPadding removes padding with ANSI X.923 mode.
func (c *Cipher) AnsiX923UnPadding(src []byte) []byte {
	if len(src) == 0 {
		return []byte("")
	}
	n := len(src) - int(src[len(src)-1])
	return src[0:n]
}

// PKCS5Padding padding with PKCS5 mode.
func (c *Cipher) PKCS5Padding(src []byte) []byte {
	return c.PKCS7Padding(src, 16)
}

// PKCS5UnPadding removes padding with PKCS5 mode.
func (c *Cipher) PKCS5UnPadding(src []byte) []byte {
	return c.PKCS7UnPadding(src)
}

// PKCS7Padding padding with PKCS7 mode.
func (c *Cipher) PKCS7Padding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(src, paddingText...)
}

// PKCS7UnPadding removes padding with PKCS7 mode.
func (c *Cipher) PKCS7UnPadding(src []byte) []byte {
	if len(src) == 0 {
		return []byte("")
	}
	n := len(src) - int(src[len(src)-1])
	return src[0:n]
}

// NewCBCEncrypter encrypts with CBC mode.
func (c *Cipher) NewCBCEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCEncrypter(block, c.iv).CryptBlocks(dst, src)
	return
}

// NewCBCDecrypter decrypts with CBC mode.
func (c *Cipher) NewCBCDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCBCDecrypter(block, c.iv).CryptBlocks(dst, src)
	return
}

// NewCFBEncrypter encrypts with CFB mode.
func (c *Cipher) NewCFBEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBEncrypter(block, c.iv).XORKeyStream(dst, src)
	return
}

// NewCFBDecrypter decrypts with CFB mode.
func (c *Cipher) NewCFBDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCFBDecrypter(block, c.iv).XORKeyStream(dst, src)
	return
}

// NewCTREncrypter encrypts with CTR mode.
func (c *Cipher) NewCTREncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv).XORKeyStream(dst, src)
	return
}

// NewCTRDecrypter decrypts with CTR mode.
func (c *Cipher) NewCTRDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewCTR(block, c.iv).XORKeyStream(dst, src)
	return
}

// NewECBEncrypter encrypts with ECB mode.
func (c *Cipher) NewECBEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	encrypted, blockSize := dst, block.BlockSize()
	for len(src) > 0 {
		block.Encrypt(encrypted, src[:blockSize])
		src = src[blockSize:]
		encrypted = encrypted[blockSize:]
	}
	return
}

// NewECBDecrypter decrypts with ECB mode.
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
func (c *Cipher) NewOFBEncrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewOFB(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// NewOFBDecrypter decrypts with OFB mode.
func (c *Cipher) NewOFBDecrypter(src []byte, block cipher.Block) (dst []byte) {
	dst = make([]byte, len(src))
	cipher.NewOFB(block, c.iv[:block.BlockSize()]).XORKeyStream(dst, src)
	return
}

// Encrypt encrypts with given mode and padding
func (c *Cipher) Encrypt(src []byte, block cipher.Block) (dst []byte, err error) {
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
	}

	switch mode {
	case ECB:
		dst = c.NewECBEncrypter(src, block)
	case CBC:
		dst = c.NewCBCEncrypter(src, block)
	case CTR:
		dst = c.NewCTREncrypter(src, block)
	case CFB:
		dst = c.NewCFBEncrypter(src, block)
	case OFB:
		dst = c.NewOFBEncrypter(src, block)
	}
	return
}

// Decrypt decrypts with given mode and padding.
func (c *Cipher) Decrypt(src []byte, block cipher.Block) (dst []byte, err error) {
	mode, padding := c.mode, c.padding
	dst = []byte("")
	if len(src) == 0 {
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
