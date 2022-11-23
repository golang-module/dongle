package dongle

import (
	"fmt"
)

// returns an invalid plaintext error
// 返回无效的明文错误
var invalidPlaintextError = func() error {
	return fmt.Errorf("invalid plaintext, the plaintext must be multiple of 16 bytes")
}

// returns an invalid ciphertext error
// 返回无效的密文错误
var invalidCiphertextError = func(mode string) error {
	return fmt.Errorf("invalid ciphertext, the ciphertext can't be decoded by %s", mode)
}

// returns an invalid rc4 key error
// 返回无效的 rc4 密钥错误
var invalidRc4KeyError = func(size int) error {
	return fmt.Errorf("invalid rc4 key size %d, the key at least 1 byte and at most 256 bytes", size)
}

// returns an invalid aes key error
// 返回无效的 aes 密钥错误
var invalidAesKeyError = func(size int) error {
	return fmt.Errorf("invalid aes key size %d, the key must be 16, 24 or 32 bytes", size)
}

// returns an invalid des key error
// 返回无效的 des 密钥错误
var invalidDesKeyError = func(size int) error {
	return fmt.Errorf("invalid des key size %d, the key must be 8 bytes", size)
}

// returns an invalid 3des key error
// 返回无效的 3des 密钥错误
var invalid3DesKeyError = func(size int) error {
	return fmt.Errorf("invalid 3des key size %d, the key must be 24 bytes", size)
}

// returns an invalid aes iv error
// 返回无效的 aes 偏移量错误
var invalidAesIVError = func(size int) error {
	return fmt.Errorf("invalid aes iv size %d, the iv size must be 16 bytes", size)
}

// returns an invalid des iv error、
// 返回无效的 des 偏移量错误
var invalidDesIVError = func(size int) error {
	return fmt.Errorf("invalid des iv size %d, the iv size must be 8 bytes", size)
}

// returns an invalid 3des iv error
// 返回无效的 3des 偏移量错误
var invalid3DesIVError = func(size int) error {
	return fmt.Errorf("invalid 3des iv size %d, the iv size must be 8 bytes", size)
}

// returns an invalid encrypt or decrypt mode error
// 返回无效的分组模式错误
var invalidModeError = func(mode cipherMode) error {
	return fmt.Errorf("invalid encrypt or decrypt mode %q", mode)
}

// returns an invalid encrypt or decrypt padding error
// 返回无效的填充方式错误
var invalidPaddingError = func(padding cipherPadding) error {
	return fmt.Errorf("invalid encrypt or decrypt padding %q", padding)
}

// returns an invalid public key error
// 返回无效的公钥错误
var invalidPublicKeyError = func() error {
	return fmt.Errorf("invalid public key, please make sure the public key is valid")
}

// returns an invalid private key error
// 返回无效的私钥错误
var invalidPrivateKeyError = func() error {
	return fmt.Errorf("invalid private key, please make sure the private key is valid")
}
