package dongle

// defines a decoder struct.
// 定义 decoder 结构体
type decoder struct {
	dongle
}

// returns a new decoder instance.
// 初始化 decoder 结构体
func newDecoder() decoder {
	return decoder{}
}

// FromString decodes from string.
// 对字符串进行解密
func (d decoder) FromString(s string) decoder {
	d.src = string2bytes(s)
	return d
}

// FromBytes decodes from byte slice.
// 对字节切片进行解密
func (d decoder) FromBytes(b []byte) decoder {
	d.src = b
	return d
}

// String implements the interface Stringer for decoder struct.
// 实现 Stringer 接口
func (d decoder) String() string {
	return d.ToString()
}

// ToString outputs as string.
// 输出字符串
func (d decoder) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (d decoder) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}
