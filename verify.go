package dongle

// defines a verifier struct.
// 定义 verifier 结构体
type verifier struct {
	dongle
	sign []byte
}

// returns a new verifier instance.
// 初始化 verify 结构体
func newVerifier() verifier {
	return verifier{}
}

// FromRawString verify signature from raw string.
// 对未经编码的原始字符串进行验签
func (v verifier) FromRawString(sign, message string) verifier {
	v.src = string2bytes(message)
	v.sign = string2bytes(sign)
	return v
}

// FromHexString decrypts from string with hex encoding.
// 对经过 hex 编码的字符串进行验签
func (v verifier) FromHexString(sign, message string) verifier {
	mac := Decode.FromString(sign).ByHex()
	if mac.Error != nil {
		v.Error = invalidDecodingError("hex")
		return v
	}
	v.sign = mac.ToBytes()
	v.src = string2bytes(message)
	return v
}

// FromBase64String decrypts from string with base64 encoding.
// 对经过 base64 编码的字符串进行验签
func (v verifier) FromBase64String(sign, message string) verifier {
	hash := Decode.FromString(sign).ByBase64()
	if hash.Error != nil {
		v.Error = invalidDecodingError("base64")
		return v
	}
	v.sign = hash.ToBytes()
	v.src = string2bytes(message)
	return v
}

// FromRawBytes verify signature from raw byte slice.
// 对未经编码的原始字节切片进行验签
func (v verifier) FromRawBytes(sign, message []byte) verifier {
	v.sign = sign
	v.src = message
	return v
}

// FromHexBytes verify from byte slice with hex encoding.
// 对经过 hex 编码的字节切片进行验签
func (v verifier) FromHexBytes(sign, message []byte) verifier {
	mac := Decode.FromBytes(sign).ByHex()
	if mac.Error != nil {
		v.Error = invalidDecodingError("hex")
		return v
	}
	v.sign = mac.ToBytes()
	v.src = message
	return v
}

// FromBase64Bytes verify from byte slice with base64 encoding.
// 对经过 base64 编码的字节切片进行验签
func (v verifier) FromBase64Bytes(sign, message []byte) verifier {
	mac := Decode.FromBytes(sign).ByBase64()
	if mac.Error != nil {
		v.Error = invalidDecodingError("base64")
		return v
	}
	v.sign = mac.ToBytes()
	v.src = message
	return v
}

// ToBool outputs as bool.
// 输出布尔值
func (v verifier) ToBool() bool {
	return v.Error == nil
}
