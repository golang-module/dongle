package dongle

// Signer defines a Signer struct.
// 定义 Signer 结构体
type Signer struct {
	dongle
}

// newSigner returns a new Signer instance.
// 初始化 Signer 结构体
func newSigner() Signer {
	return Signer{}
}

// FromString signs from string.
// 对字符串进行签名
func (s Signer) FromString(message string) Signer {
	s.src = string2bytes(message)
	return s
}

// FromBytes signs from byte slice.
// 对字节切片进行签名
func (s Signer) FromBytes(message []byte) Signer {
	s.src = message
	return s
}

// String implements Stringer interface for Signer struct.
// 实现 Stringer 接口
func (s Signer) String() string {
	return s.ToRawString()
}

// ToRawString outputs as raw string without encoding.
// 输出未经编码的原始字符串
func (s Signer) ToRawString() string {
	return bytes2string(s.dst)
}

// ToHexString outputs as string with hex encoding.
// 输出经过 hex 编码的字符串
func (s Signer) ToHexString() string {
	return Encode.FromBytes(s.dst).ByHex().ToString()
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (s Signer) ToBase64String() string {
	return Encode.FromBytes(s.dst).ByBase64().ToString()
}

// ToRawBytes outputs as raw byte slice without encoding.
// 输出未经编码的原始字节切片
func (s Signer) ToRawBytes() []byte {
	if len(s.dst) == 0 {
		return []byte("")
	}
	return s.dst
}

// ToHexBytes outputs as byte with hex encoding.
// 输出经过 hex 编码的字节切片
func (s Signer) ToHexBytes() []byte {
	return Encode.FromBytes(s.dst).ByHex().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
// 输出经过 base64 编码的字节切片
func (s Signer) ToBase64Bytes() []byte {
	return Encode.FromBytes(s.dst).ByBase64().ToBytes()
}
