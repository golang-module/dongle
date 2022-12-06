package dongle

import (
	"fmt"
)

var (
	// returns an invalid aes src error
	// 返回无效的 aes 明文错误
	invalidAesSrcError = func() error {
		return fmt.Errorf("aes: invalid src, the src with no padding must be multiple of 16 bytes")
	}
	// returns an invalid aes key error
	// 返回无效的 aes 密钥错误
	invalidAesKeyError = func() error {
		return fmt.Errorf("aes: invalid key, the key must be 16, 24 or 32 bytes")
	}
	// returns an invalid aes iv error、
	// 返回无效的 aes 偏移量错误
	invalidAesIVError = func() error {
		return fmt.Errorf("aes: invalid iv, the iv size must be 16 bytes")
	}
)

var (
	// returns an invalid des src error
	// 返回无效的 des 明文错误
	invalidDesSrcError = func() error {
		return fmt.Errorf("des: invalid src, the src with no padding must be multiple of 16 bytes")
	}
	// returns an invalid des key error
	// 返回无效的 des 密钥错误
	invalidDesKeyError = func() error {
		return fmt.Errorf("des: invalid key, the key must be 8 bytes")
	}
	// returns an invalid des iv error、
	// 返回无效的 des 偏移量错误
	invalidDesIVError = func() error {
		return fmt.Errorf("des: invalid iv, the iv size must be 8 bytes")
	}
)

var (
	// returns an invalid 3des src error
	// 返回无效的 3des 明文错误
	invalid3DesSrcError = func() error {
		return fmt.Errorf("3des: invalid src, the src with no padding must be multiple of 16 bytes")
	}
	// returns an invalid 3des key error
	// 返回无效的 3des 密钥错误
	invalid3DesKeyError = func() error {
		return fmt.Errorf("3des: invalids key, the key must be 24 bytes")
	}
	// returns an invalid 3des iv error、
	// 返回无效的 3des 偏移量错误
	invalid3DesIVError = func() error {
		return fmt.Errorf("3des: invalid iv, the iv size must be 8 bytes")
	}
)

var (
	// returns an invalid rsa public key error
	// 返回无效的 rsa 公钥错误
	invalidRsaPublicKeyError = func() error {
		return fmt.Errorf("rsa: invalid public key, please make sure the public key is valid")
	}
	// returns an invalid rsa private key error
	// 返回无效的 rsa 私钥错误
	invalidRsaPrivateKeyError = func() error {
		return fmt.Errorf("rsa: invalid private key, please make sure the private key is valid")
	}
	// returns an invalid rsa hash function error
	// 返回无效的 rsa 哈希函数错误
	invalidRsaHashError = func() error {
		return fmt.Errorf("rsa: invalid hash function, the hash function is unsupported")
	}
)

var (
	// returns an invalid tea src error
	// 返回无效的 tea 明文错误
	invalidTeaSrcError = func() error {
		return fmt.Errorf("tea: invalid src, the src must be 8 bytes")
	}
	// returns an invalid tea round error
	// 返回无效的 tea 迭代轮数错误
	invalidTeaRoundsError = func() error {
		return fmt.Errorf("tea: invalid rounds, the rounds must be even")
	}
	// returns an invalid tea key error
	// 返回无效的 tea 密钥错误
	invalidTeaKeyError = func() error {
		return fmt.Errorf("tea: invalid key, the key must be 16 bytes")
	}
)

var (
	// returns an invalid morse src error
	// 返回无效的 morse 明文错误
	invalidMorseSrcError = func() error {
		return fmt.Errorf("morse: invalid src, the src can't contain spaces")
	}
	// returns a morse decoding error
	// 返回 morse 解码错误
	morseDecodeError = func() error {
		return fmt.Errorf("morse: invalid decoding, the src can't be decoded")
	}
)

var (
	// returns an invalid bcrypt round error
	// 返回无效的 bcrypt 迭代轮数错误
	invalidBcryptRoundsError = func() error {
		return fmt.Errorf("bcrypt: invalid round, the round is outside allowed range (4,31)")
	}
)
var (
	// returns an invalid rc4 key error
	// 返回无效的 rc4 密钥错误
	invalidRc4KeyError = func() error {
		return fmt.Errorf("rc4: invalid key, the key at least 1 byte and at most 256 bytes")
	}
)
var (
	// returns an invalid hash digests error
	// 返回无效的哈希大小错误
	invalidHashSizeError = func() error {
		return fmt.Errorf("hash: invalid size, the size is unsupported")
	}
)

var (
	// returns an invalid encrypt or decrypt mode error
	// 返回无效的分组模式错误
	invalidModeError = func(mode cipherMode) error {
		return fmt.Errorf("invalid encrypt or decrypt mode %q", mode)
	}
	// returns an invalid encrypt or decrypt padding error
	// 返回无效的填充方式错误
	invalidPaddingError = func(padding cipherPadding) error {
		return fmt.Errorf("invalid encrypt or decrypt padding %q", padding)
	}
	// returns an invalid decoding error
	// 返回无效的解码方式错误
	invalidDecodingError = func(mode string) error {
		return fmt.Errorf("invalid decoding, the src can't be decoded by %s", mode)
	}
)
