package dongle

// Decrypter defines Decrypter struct.
// 定义 Decrypter 结构体
type Decrypter struct {
	dongle
}

// newDecrypter returns a new Decrypter instance.
// 初始化 Decrypter 结构体
func newDecrypter() Decrypter {
	return Decrypter{}
}

// FromRawString decrypts from raw string without encoding.
// 对未经编码的原始字符串进行解密
func (d Decrypter) FromRawString(s string) Decrypter {
	d.src = string2bytes(s)
	return d
}

// FromHexString decrypts from string with hex encoding.
// 对经过 hex 编码的字符串进行解密
func (d Decrypter) FromHexString(s string) Decrypter {
	decode := Decode.FromString(s).ByHex()
	if decode.Error != nil {
		d.Error = invalidDecodingError("hex")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromBase64String decrypts from string with base64 encoding.
// 对经过 base64 编码的字符串进行解密
func (d Decrypter) FromBase64String(s string) Decrypter {
	decode := Decode.FromString(s).ByBase64()
	if decode.Error != nil {
		d.Error = invalidDecodingError("base64")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromRawBytes decrypts from raw byte slice without encoding.
// 对未经编码的原始字节切片进行解密
func (d Decrypter) FromRawBytes(b []byte) Decrypter {
	d.src = b
	return d
}

// FromHexBytes decrypts from byte slice with hex encoding.
// 对经过 hex 编码的字节切片进行解密
func (d Decrypter) FromHexBytes(b []byte) Decrypter {
	decode := Decode.FromBytes(b).ByHex()
	if decode.Error != nil {
		d.Error = invalidDecodingError("hex")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromBase64Bytes decrypts from byte slice with base64 encoding.
// 对经过 base64 编码的字节切片进行解密
func (d Decrypter) FromBase64Bytes(b []byte) Decrypter {
	decode := Decode.FromBytes(b).ByBase64()
	if decode.Error != nil {
		d.Error = invalidDecodingError("base64")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// String implements Stringer interface for Decrypter struct.
// 实现 Stringer 接口
func (d Decrypter) String() string {
	return d.ToString()
}

// ToString outputs as string.
// 输出字符串
func (d Decrypter) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (d Decrypter) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}
