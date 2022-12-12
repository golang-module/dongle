package dongle

import (
	"crypto/ed25519"
)

// defines encoding mode enum type.
// 定义编码模式枚举类型
type encodingMode string

// encoding mode constants
// 编码模式常量
const (
	Raw    encodingMode = "raw"
	Hex    encodingMode = "hex"
	Base64 encodingMode = "base64"
)

// ByEd25519 signs by ed25519.
// 通过 ed25519 私钥签名
func (s Signer) ByEd25519(privateKey interface{}, mode encodingMode) Signer {
	if len(s.src) == 0 || s.Error != nil {
		return s
	}
	pri, err := getDecodedKey(interface2bytes(privateKey), mode)
	if err != nil {
		s.Error = err
		return s
	}
	if len(pri) != ed25519.PrivateKeySize {
		s.Error = invalidEd25519PrivateKeyError()
		return s
	}
	s.dst = ed25519.Sign(pri, s.src)
	return s
}

// ByEd25519 verify by ed25519.
// 通过 ed25519 公钥验签
func (v Verifier) ByEd25519(publicKey interface{}, mode encodingMode) Verifier {
	if len(v.src) == 0 || v.Error != nil {
		return v
	}
	pub, err := getDecodedKey(interface2bytes(publicKey), mode)
	if err != nil {
		v.Error = err
		return v
	}
	if len(pub) != ed25519.PublicKeySize {
		v.Error = invalidEd25519PublicKeyError()
		return v
	}
	if ed25519.Verify(pub, v.src, v.sign) == false {
		v.Error = invalidEd25519SignatureError()
		return v
	}
	return v
}

// gets the decoded key
// 获取解码的 key
func getDecodedKey(key []byte, mode encodingMode) (dst []byte, err error) {
	var decode Decoder
	switch mode {
	case Raw:
		dst = key
		return
	case Hex:
		decode = Decode.FromBytes(key).ByHex()
	case Base64:
		decode = Decode.FromBytes(key).ByBase64()
	}
	if decode.Error != nil {
		err = decode.Error
		return
	}
	dst = decode.ToBytes()
	return
}
