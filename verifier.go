package dongle

// Verifier defines a Verifier struct.
// 定义 Verifier 结构体
type Verifier struct {
	dongle
	sign []byte
}

// newVerifier returns a new Verifier instance.
// 初始化 Verifier 结构体
func newVerifier() Verifier {
	return Verifier{}
}

// FromRawString verify signature from raw string.
// 对未经编码的原始字符串签名进行验签
func (v Verifier) FromRawString(signature, message string) Verifier {
	v.src = string2bytes(message)
	v.sign = string2bytes(signature)
	return v
}

// FromHexString decrypts from string with hex encoding.
// 对经过 hex 编码的字符串签名进行验签
func (v Verifier) FromHexString(signature, message string) Verifier {
	d := Decode.FromString(signature).ByHex()
	if d.Error != nil {
		v.Error = invalidDecodingError("hex")
		return v
	}
	v.sign = d.ToBytes()
	v.src = string2bytes(message)
	return v
}

// FromBase64String decrypts from string with base64 encoding.
// 对经过 base64 编码的字符串签名进行验签
func (v Verifier) FromBase64String(signature, message string) Verifier {
	d := Decode.FromString(signature).ByBase64()
	if d.Error != nil {
		v.Error = invalidDecodingError("base64")
		return v
	}
	v.sign = d.ToBytes()
	v.src = string2bytes(message)
	return v
}

// FromRawBytes verify signature from raw byte slice.
// 对未经编码的原始字节切片签名进行验签
func (v Verifier) FromRawBytes(signature, message []byte) Verifier {
	v.sign = signature
	v.src = message
	return v
}

// FromHexBytes verify from byte slice with hex encoding.
// 对经过 hex 编码的字节切片签名进行验签
func (v Verifier) FromHexBytes(signature, message []byte) Verifier {
	d := Decode.FromBytes(signature).ByHex()
	if d.Error != nil {
		v.Error = invalidDecodingError("hex")
		return v
	}
	v.sign = d.ToBytes()
	v.src = message
	return v
}

// FromBase64Bytes verify from byte slice with base64 encoding.
// 对经过 base64 编码的字节切片签名进行验签
func (v Verifier) FromBase64Bytes(signature, message []byte) Verifier {
	d := Decode.FromBytes(signature).ByBase64()
	if d.Error != nil {
		v.Error = invalidDecodingError("base64")
		return v
	}
	v.sign = d.ToBytes()
	v.src = message
	return v
}

// ToBool outputs as bool.
// 输出布尔值
func (v Verifier) ToBool() bool {
	if len(v.src) == 0 || len(v.sign) == 0 {
		return false
	}
	return v.Error == nil
}
