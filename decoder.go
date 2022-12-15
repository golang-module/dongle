package dongle

// Decoder defines Decoder struct.
// 定义 Decoder 结构体
type Decoder struct {
	dongle
}

// newDecoder returns a new Decoder instance.
// 初始化 Decoder 结构体
func newDecoder() Decoder {
	return Decoder{}
}

// FromString decodes from string.
// 对字符串进行解密
func (d Decoder) FromString(s string) Decoder {
	d.src = string2bytes(s)
	return d
}

// FromBytes decodes from byte slice.
// 对字节切片进行解密
func (d Decoder) FromBytes(b []byte) Decoder {
	d.src = b
	return d
}

// String implements Stringer interface for Decoder struct.
// 实现 Stringer 接口
func (d Decoder) String() string {
	return d.ToString()
}

// ToString outputs as string.
// 输出字符串
func (d Decoder) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (d Decoder) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}
