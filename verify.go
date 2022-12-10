package dongle

// defines a verifier struct.
// 定义 verifier 结构体
type verifier struct {
	dongle
	sign []byte
}

// NewVerifier returns a new verifier instance.
// 初始化 verify 结构体
func NewVerifier() verifier {
	return verifier{}
}

// FromRawString verify signature from raw string.
// 对未经编码的原始字符串签名进行验签
func (v verifier) FromRawString(signature, message string) verifier {
	v.src = string2bytes(message)
	v.sign = string2bytes(signature)
	return v
}

// FromHexString decrypts from string with hex encoding.
// 对经过 hex 编码的字符串签名进行验签
func (v verifier) FromHexString(signature, message string) verifier {
	decode := Decode.FromString(signature).ByHex()
	if decode.Error != nil {
		v.Error = invalidDecodingError("hex")
		return v
	}
	v.sign = decode.ToBytes()
	v.src = string2bytes(message)
	return v
}

// FromBase64String decrypts from string with base64 encoding.
// 对经过 base64 编码的字符串签名进行验签
func (v verifier) FromBase64String(signature, message string) verifier {
	decode := Decode.FromString(signature).ByBase64()
	if decode.Error != nil {
		v.Error = invalidDecodingError("base64")
		return v
	}
	v.sign = decode.ToBytes()
	v.src = string2bytes(message)
	return v
}

// FromRawBytes verify signature from raw byte slice.
// 对未经编码的原始字节切片签名进行验签
func (v verifier) FromRawBytes(signature, message []byte) verifier {
	v.sign = signature
	v.src = message
	return v
}

// FromHexBytes verify from byte slice with hex encoding.
// 对经过 hex 编码的字节切片签名进行验签
func (v verifier) FromHexBytes(signature, message []byte) verifier {
	decode := Decode.FromBytes(signature).ByHex()
	if decode.Error != nil {
		v.Error = invalidDecodingError("hex")
		return v
	}
	v.sign = decode.ToBytes()
	v.src = message
	return v
}

// FromBase64Bytes verify from byte slice with base64 encoding.
// 对经过 base64 编码的字节切片签名进行验签
func (v verifier) FromBase64Bytes(signature, message []byte) verifier {
	decode := Decode.FromBytes(signature).ByBase64()
	if decode.Error != nil {
		v.Error = invalidDecodingError("base64")
		return v
	}
	v.sign = decode.ToBytes()
	v.src = message
	return v
}

// ToBool outputs as bool.
// 输出布尔值
func (v verifier) ToBool() bool {
	if len(v.src) == 0 || len(v.sign) == 0 {
		return false
	}
	return v.Error == nil
}
