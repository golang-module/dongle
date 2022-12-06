package dongle

// encode defines a encoder struct.
// 定义 encoder 结构体
type encoder struct {
	dongle
}

// returns a new encoder instance.
// 初始化 encoder 结构体
func newEncoder() encoder {
	return encoder{}
}

// FromString encodes from string.
// 对字符串进行编码
func (e encoder) FromString(s string) encoder {
	e.src = string2bytes(s)
	return e
}

// FromBytes encodes from byte slice.
// 对字节切片进行编码
func (e encoder) FromBytes(b []byte) encoder {
	e.src = b
	return e
}

// String implements the interface Stringer for encoder struct.
// 实现 Stringer 接口
func (e encoder) String() string {
	return e.ToString()
}

// ToString outputs as string.
// 输出字符串
func (e encoder) ToString() string {
	return bytes2string(e.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (e encoder) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}
