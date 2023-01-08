package dongle

// Encrypter defines Encrypter struct.
// 定义 Encrypter 结构体
type Encrypter struct {
	dongle
}

// newEncrypter returns a new Encrypter instance.
// 初始化 Encrypter 结构体
func newEncrypter() Encrypter {
	return Encrypter{}
}

// FromString encrypts from string.
// 对字符串加密
func (e Encrypter) FromString(s string) Encrypter {
	e.src = string2bytes(s)
	return e
}

// FromBytes encrypts from byte slice.
// 对字节切片加密
func (e Encrypter) FromBytes(b []byte) Encrypter {
	e.src = b
	return e
}

// String implements Stringer interface for Encrypter struct.
// 实现 Stringer 接口
func (e Encrypter) String() string {
	return e.ToRawString()
}

// ToRawString outputs as raw string without encoding.
// 输出未经编码的原始字符串
func (e Encrypter) ToRawString() string {
	return bytes2string(e.dst)
}

// ToHexString outputs as string with hex encoding.
// 输出经过 hex 编码的字符串
func (e Encrypter) ToHexString() string {
	return Encode.FromBytes(e.dst).ByHex().ToString()
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (e Encrypter) ToBase64String() string {
	return Encode.FromBytes(e.dst).ByBase64().ToString()
}

// ToRawBytes outputs as raw byte slice without encoding.
// 输出未经编码的原始字节切片
func (e Encrypter) ToRawBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}

// ToHexBytes outputs as byte with hex encoding.
// 输出经过 hex 编码的字节切片
func (e Encrypter) ToHexBytes() []byte {
	return Encode.FromBytes(e.dst).ByHex().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
// 输出经过 base64 编码的字节切片
func (e Encrypter) ToBase64Bytes() []byte {
	return Encode.FromBytes(e.dst).ByBase64().ToBytes()
}
