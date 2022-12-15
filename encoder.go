package dongle

// Encoder defines Encoder struct.
// 定义 Encoder 结构体
type Encoder struct {
	dongle
}

// newEncoder returns a new Encoder instance.
// 初始化 encoder 结构体
func newEncoder() Encoder {
	return Encoder{}
}

// FromString encodes from string.
// 对字符串进行编码
func (e Encoder) FromString(s string) Encoder {
	e.src = string2bytes(s)
	return e
}

// FromBytes encodes from byte slice.
// 对字节切片进行编码
func (e Encoder) FromBytes(b []byte) Encoder {
	e.src = b
	return e
}

// String implements Stringer interface for Encoder struct.
// 实现 Stringer 接口
func (e Encoder) String() string {
	return e.ToString()
}

// ToString outputs as string.
// 输出字符串
func (e Encoder) ToString() string {
	return bytes2string(e.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (e Encoder) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}
